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

package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	commonv1 "github.com/kubeshop/testkube-operator/api/common/v1"
	templatesv1 "github.com/kubeshop/testkube-operator/api/template/v1"
	testsv3 "github.com/kubeshop/testkube-operator/api/tests/v3"
	cronjobclient "github.com/kubeshop/testkube-operator/pkg/cronjob/client"
	cronjobmanager "github.com/kubeshop/testkube-operator/pkg/cronjob/manager"
)

// TestReconciler reconciles a Test object
type TestReconciler struct {
	client.Client
	Scheme          *runtime.Scheme
	CronJobClient   cronjobclient.Interface
	CronJobManager  cronjobmanager.Interface
	ServiceName     string
	ServicePort     int
	PurgeExecutions bool
}

//+kubebuilder:rbac:groups=tests.testkube.io,resources=tests,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tests.testkube.io,resources=tests/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tests.testkube.io,resources=tests/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Test object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *TestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	isNewArchitecture := r.CronJobManager.IsNamespaceForNewArchitecture(req.NamespacedName.Namespace)
	// Delete CronJob if it was created for deleted Test
	var test testsv3.Test
	if err := r.Get(ctx, req.NamespacedName, &test); err != nil {
		if errors.IsNotFound(err) {
			if !isNewArchitecture {
				if err = r.CronJobClient.Delete(ctx,
					cronjobclient.GetMetadataName(req.NamespacedName.Name, cronjobclient.TestResourceURI), req.NamespacedName.Namespace); err != nil {
					return ctrl.Result{}, err
				}
			}

			if _, err = r.deleteTest(req.NamespacedName.Name, req.NamespacedName.Namespace); err != nil {
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	if isNewArchitecture {
		return ctrl.Result{}, nil
	}

	// Delete CronJob if it was created for cleaned Test schedule
	if test.Spec.Schedule == "" {
		if err := r.CronJobClient.Delete(ctx,
			cronjobclient.GetMetadataName(req.NamespacedName.Name, cronjobclient.TestResourceURI), req.NamespacedName.Namespace); err != nil {
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	data, err := json.Marshal(testsv3.ExecutionRequest{
		RunningContext: &testsv3.RunningContext{
			Type_:   commonv1.RunningContextTypeScheduler,
			Context: test.Spec.Schedule,
		},
	})
	if err != nil {
		return ctrl.Result{}, err
	}

	var jobTemplate, jobTemplateExt string
	if test.Spec.ExecutionRequest != nil {
		jobTemplateExt = test.Spec.ExecutionRequest.CronJobTemplate
		if test.Spec.ExecutionRequest.CronJobTemplateReference != "" {
			var template templatesv1.Template
			object := types.NamespacedName{
				Namespace: req.Namespace,
				Name:      test.Spec.ExecutionRequest.CronJobTemplateReference,
			}
			if err = r.Get(ctx, object, &template); err != nil {
				return ctrl.Result{}, err
			}

			if template.Spec.Type_ != nil && *template.Spec.Type_ == templatesv1.CRONJOB_TemplateType {
				jobTemplate = template.Spec.Body
			} else {
				ctrl.Log.Info("not matched template type", "template", test.Spec.ExecutionRequest.CronJobTemplateReference)
			}
		}
	}

	options := cronjobclient.Options{
		Schedule:                  test.Spec.Schedule,
		Group:                     testsv3.Group,
		Resource:                  testsv3.Resource,
		Version:                   testsv3.Version,
		ResourceURI:               cronjobclient.TestResourceURI,
		Data:                      string(data),
		Labels:                    test.Labels,
		CronJobTemplate:           jobTemplate,
		CronJobTemplateExtensions: jobTemplateExt,
	}

	// Create CronJob if it was not created before for provided Test schedule
	cronJob, err := r.CronJobClient.Get(ctx,
		cronjobclient.GetMetadataName(req.NamespacedName.Name, cronjobclient.TestResourceURI), req.NamespacedName.Namespace)
	if err != nil {
		if errors.IsNotFound(err) {
			if err = r.CronJobClient.Create(ctx, test.Name,
				cronjobclient.GetMetadataName(test.Name, cronjobclient.TestResourceURI), req.NamespacedName.Namespace,
				string(test.UID), options); err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, err
	}

	// Update CronJob if it was created before provided Test schedule
	if err = r.CronJobClient.Update(ctx, cronJob, test.Name,
		cronjobclient.GetMetadataName(test.Name, cronjobclient.TestResourceURI), req.NamespacedName.Namespace,
		string(test.UID), options); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	pred := predicate.GenerationChangedPredicate{}
	return ctrl.NewControllerManagedBy(mgr).
		For(&testsv3.Test{}).
		WithEventFilter(pred).
		Complete(r)
}

func (r *TestReconciler) deleteTest(testName, namespace string) (out string, err error) {
	if !r.PurgeExecutions {
		return out, nil
	}

	request, err := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("http://%s.%s.svc.cluster.local:%d/v1/tests/%s?skipDeleteCRD=true",
			r.ServiceName, namespace, r.ServicePort, testName), nil)
	if err != nil {
		return out, err
	}

	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		return out, fmt.Errorf("could not DELETE, statusCode: %d", resp.StatusCode)
	}

	return fmt.Sprintf("status: %d", resp.StatusCode), err
}
