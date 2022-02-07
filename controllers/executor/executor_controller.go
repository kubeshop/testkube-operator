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

package executor

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	executorv1 "github.com/kubeshop/testkube-operator/apis/executor/v1"
)

// ExecutorReconciler reconciles a Executor object
type ExecutorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=executor.testkube.io,resources=executors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=executor.testkube.io,resources=executors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=executor.testkube.io,resources=executors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Executor object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *ExecutorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExecutorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&executorv1.Executor{}).
		Complete(r)
}

// LoadDefaultExecutors loads default executors
func (r *ExecutorReconciler) LoadDefaultExecutors(namespace, data string) error {
	var executors []executorv1.ExecutorDetails

	dataDecoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(dataDecoded), &executors); err != nil {
		return err
	}

	for _, executor := range executors {
		if executor.Spec == nil {
			continue
		}

		var obj executorv1.Executor
		err := r.Client.Get(context.Background(), types.NamespacedName{Name: executor.Name, Namespace: namespace}, &obj)
		if err != nil && errors.IsNotFound(err) {
			return err
		}

		if err != nil {
			if err = r.Client.Create(context.Background(), &executorv1.Executor{
				ObjectMeta: metav1.ObjectMeta{
					Name:      executor.Name,
					Namespace: namespace,
				},
				Spec: *executor.Spec,
			}); err != nil {
				return err
			}
		}
	}

	return nil
}
