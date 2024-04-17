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

package testworkflows

import (
	"context"
	"fmt"
	"time"

	testworkflowsv1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"
	"github.com/kubeshop/testkube-operator/pkg/cronjob"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	reconciliationDateAnnotationName = "testworkflows.testkube.io/reconciliation-date"
)

// TestWorkflowTemplateReconciler reconciles a TestWorkflowTemplate object
type TestWorkflowTemplateReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=testworkflows.testkube.io,resources=testworkflowtemplates,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=testworkflows.testkube.io,resources=testworkflowtemplates/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TestWorkflowTemplate object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *TestWorkflowTemplateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var testWorkflowList testworkflowsv1.TestWorkflowList
	reqs, err := labels.ParseToRequirements(cronjob.GetSelector("yes", cronjob.TestWorkflowTemplateResourceURI))
	if err != nil {
		return ctrl.Result{}, err
	}

	options := &client.ListOptions{
		Namespace:     req.NamespacedName.Namespace,
		LabelSelector: labels.NewSelector().Add(reqs...),
	}
	if err := r.List(ctx, &testWorkflowList, options); err != nil {
		return ctrl.Result{}, nil
	}

	for i := range testWorkflowList.Items {
		var isUsed bool
		for _, template := range testWorkflowList.Items[i].Spec.Use {
			if template.Name == req.NamespacedName.Name {
				isUsed = true
				break
			}
		}

		if isUsed {
			if testWorkflowList.Items[i].Annotations == nil {
				testWorkflowList.Items[i].Annotations = make(map[string]string)
			}

			testWorkflowList.Items[i].Annotations[reconciliationDateAnnotationName] = fmt.Sprint(time.Now().UnixNano())
			if err := r.Update(ctx, &testWorkflowList.Items[i]); err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestWorkflowTemplateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testworkflowsv1.TestWorkflowTemplate{}).
		Complete(r)
}
