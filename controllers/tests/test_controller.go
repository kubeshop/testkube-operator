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

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	testsv3 "github.com/kubeshop/testkube-operator/apis/tests/v3"
	"github.com/kubeshop/testkube-operator/pkg/cronjob"
)

// TestReconciler reconciles a Test object
type TestReconciler struct {
	client.Client
	Scheme *runtime.Scheme
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

	// TODO(user): your logic here
	var test testsv3.Test
	if err := r.Get(ctx, req.NamespacedName, &test); err != nil {
		if errors.IsNotFound(err) {
			var cronJob batchv1.CronJob
			if err = r.Get(ctx, types.NamespacedName{
				Name:      cronjob.GetMetadataName(req.NamespacedName.Name, "test"),
				Namespace: req.NamespacedName.Namespace}, &cronJob); err != nil {
				if errors.IsNotFound(err) {
					return ctrl.Result{}, nil
				}

				return ctrl.Result{}, err
			}

			if err = r.Delete(ctx, &cronJob, &client.DeleteOptions{}); err != nil {
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testsv3.Test{}).
		Complete(r)
}
