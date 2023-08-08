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

package testexecution

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	testexecutionv1 "github.com/kubeshop/testkube-operator/apis/testexecution/v1"
)

// TestExecutionReconciler reconciles a TestExecution object
type TestExecutionReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	ServiceName string
	ServicePort int
}

//+kubebuilder:rbac:groups=tests.testkube.io,resources=testexecutions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tests.testkube.io,resources=testexecutions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tests.testkube.io,resources=testexecutions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TestExecution object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *TestExecutionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var testExecution testexecutionv1.TestExecution
	err := r.Get(ctx, req.NamespacedName, &testExecution)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	if testExecution.Spec.Test == nil {
		return ctrl.Result{}, nil
	}

	if testExecution.Spec.ExecutionRequest != nil {
		testExecution.Spec.ExecutionRequest.RunningContext = &testexecutionv1.RunningContext{
			Type_:   testexecutionv1.RunningContextTypeTestSuiteExecution,
			Context: testExecution.Name,
		}
	}

	jsonData, err := json.Marshal(testExecution.Spec.ExecutionRequest)
	if err != nil {
		return ctrl.Result{}, err
	}

	if _, err = r.executeTest(testExecution.Spec.Test.Name, testExecution.Name, testExecution.Namespace, jsonData); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestExecutionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	pred := predicate.GenerationChangedPredicate{}
	return ctrl.NewControllerManagedBy(mgr).
		For(&testexecutionv1.TestExecution{}).
		WithEventFilter(pred).
		Complete(r)
}

func (r *TestExecutionReconciler) executeTest(testName, testExecutionName, namespace string, jsonData []byte) (out string, err error) {
	request, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("http://%s.%s.svc.cluster.local:%d/v1/tests/%s/executions?testExecutionName=%s",
			r.ServiceName, namespace, r.ServicePort, testName, testExecutionName),
		bytes.NewBuffer(jsonData))
	if err != nil {
		return out, err
	}

	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 300 {
		return out, fmt.Errorf("could not POST, statusCode: %d", resp.StatusCode)
	}

	return fmt.Sprintf("status: %d - %s", resp.StatusCode, b), err
}
