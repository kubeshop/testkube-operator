/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testsuite

import (
	"context"
	"encoding/json"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	templatesv1 "github.com/kubeshop/testkube-operator/apis/template/v1"
	testsuitev3 "github.com/kubeshop/testkube-operator/apis/testsuite/v3"
	"github.com/kubeshop/testkube-operator/pkg/cronjob"
)

// TestSuiteReconciler reconciles a TestSuite object
type TestSuiteReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	CronJobClient *cronjob.Client
}

//+kubebuilder:rbac:groups=tests.testkube.io,resources=testsuites,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tests.testkube.io,resources=testsuites/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tests.testkube.io,resources=testsuites/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TestSuite object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *TestSuiteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Delete CronJob if it was created for deleted TestSuite
	var testSuite testsuitev3.TestSuite
	err := r.Get(ctx, req.NamespacedName, &testSuite)
	if err != nil {
		if errors.IsNotFound(err) {
			if err = r.CronJobClient.Delete(ctx,
				cronjob.GetMetadataName(req.NamespacedName.Name, cronjob.TestSuiteResourceURI), req.NamespacedName.Namespace); err != nil {
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	// Delete CronJob if it was created for cleaned TestSuite schedule
	if testSuite.Spec.Schedule == "" {
		if err := r.CronJobClient.Delete(ctx,
			cronjob.GetMetadataName(req.NamespacedName.Name, cronjob.TestSuiteResourceURI), req.NamespacedName.Namespace); err != nil {
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	data, err := json.Marshal(testsuitev3.TestSuiteExecutionRequest{
		RunningContext: &testsuitev3.RunningContext{
			Type_:   testsuitev3.RunningContextTypeScheduler,
			Context: testSuite.Spec.Schedule,
		},
	})

	if err != nil {
		return ctrl.Result{}, err
	}

	var jobTemplate, jobTemplateExt string
	if testSuite.Spec.ExecutionRequest != nil {
		jobTemplateExt = testSuite.Spec.ExecutionRequest.CronJobTemplate
		if testSuite.Spec.ExecutionRequest.CronJobTemplateReference != "" {
			var template templatesv1.Template
			object := types.NamespacedName{
				Namespace: req.Namespace,
				Name:      testSuite.Spec.ExecutionRequest.CronJobTemplateReference,
			}
			if err = r.Get(ctx, object, &template); err != nil {
				return ctrl.Result{}, err
			}

			if template.Spec.Type_ != nil && *template.Spec.Type_ == templatesv1.CRONJOB_TemplateType {
				jobTemplate = template.Spec.Body
			} else {
				ctrl.Log.Info("not matching template type", "template", testSuite.Spec.ExecutionRequest.CronJobTemplateReference)
			}
		}
	}

	options := cronjob.CronJobOptions{
		Schedule:                  testSuite.Spec.Schedule,
		Resource:                  cronjob.TestSuiteResourceURI,
		Data:                      string(data),
		Labels:                    testSuite.Labels,
		CronJobTemplate:           jobTemplate,
		CronJobTemplateExtensions: jobTemplateExt,
	}

	// Create CronJob if it was not created before for provided TestSuite schedule
	cronJob, err := r.CronJobClient.Get(ctx,
		cronjob.GetMetadataName(req.NamespacedName.Name, cronjob.TestSuiteResourceURI), req.NamespacedName.Namespace)
	if err != nil {
		if errors.IsNotFound(err) {
			if err = r.CronJobClient.Create(ctx, testSuite.Name,
				cronjob.GetMetadataName(testSuite.Name, cronjob.TestSuiteResourceURI), req.NamespacedName.Namespace, options); err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, err
	}

	// Update CronJob if it was created before provided Test schedule
	if err = r.CronJobClient.Update(ctx, cronJob, testSuite.Name,
		cronjob.GetMetadataName(testSuite.Name, cronjob.TestSuiteResourceURI), req.NamespacedName.Namespace, options); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestSuiteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	pred := predicate.GenerationChangedPredicate{}
	return ctrl.NewControllerManagedBy(mgr).
		For(&testsuitev3.TestSuite{}).
		WithEventFilter(pred).
		Complete(r)
}
