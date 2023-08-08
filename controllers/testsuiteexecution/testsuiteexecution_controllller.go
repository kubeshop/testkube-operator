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

package testsuiteexecution

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

	testsuiteexecutionv1 "github.com/kubeshop/testkube-operator/apis/testsuiteexecution/v1"
)

// TestSuiteExecutionReconciler reconciles a TestSuiteExecution object
type TestSuiteExecutionReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	ServiceName string
	ServicePort int
}

//+kubebuilder:rbac:groups=tests.testkube.io,resources=testsuiteexecutions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tests.testkube.io,resources=testsuiteexecutions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tests.testkube.io,resources=testsuiteexecutions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TestSuiteExecution object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *TestSuiteExecutionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var testSuiteExecution testsuiteexecutionv1.TestSuiteExecution
	err := r.Get(ctx, req.NamespacedName, &testSuiteExecution)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	if testSuiteExecution.Spec.TestSuite == nil {
		return ctrl.Result{}, nil
	}

	if testSuiteExecution.Spec.ExecutionRequest != nil {
		testSuiteExecution.Spec.ExecutionRequest.RunningContext = &testsuiteexecutionv1.RunningContext{
			Type_:   testsuiteexecutionv1.RunningContextTypeTestSuiteExecution,
			Context: testSuiteExecution.Name,
		}
	}

	jsonData, err := json.Marshal(testSuiteExecution.Spec.ExecutionRequest)
	if err != nil {
		return ctrl.Result{}, err
	}

	if _, err = r.executeTestSuite(testSuiteExecution.Spec.TestSuite.Name, testSuiteExecution.Name, testSuiteExecution.Namespace, jsonData); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestSuiteExecutionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	pred := predicate.GenerationChangedPredicate{}
	return ctrl.NewControllerManagedBy(mgr).
		For(&testsuiteexecutionv1.TestSuiteExecution{}).
		WithEventFilter(pred).
		Complete(r)
}

func (r *TestSuiteExecutionReconciler) executeTestSuite(testSuiteName, testSuiteExecutionName, namespace string, jsonData []byte) (out string, err error) {
	request, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("http://%s.%s.svc.cluster.local:%d/v1/test-suites/%s/executions?testSuiteExecutionName=%s",
			r.ServiceName, namespace, r.ServicePort, testSuiteName, testSuiteExecutionName),
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
