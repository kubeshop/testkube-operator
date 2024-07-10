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

package testworkflowexecution

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

	testworkflowsv1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"
)

// TestWorkflowExecutionReconciler reconciles a TestWorkflowExecution object
type TestWorkflowExecutionReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	ServiceName string
	ServicePort int
}

//+kubebuilder:rbac:groups=testworkflows.testkube.io,resources=testworkflowexecutions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=testworkflows.testkube.io,resources=testworkflowexecutions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=testworkflows.testkube.io,resources=testworkflowexecutions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TestWorkflowExecution object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *TestWorkflowExecutionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var testWorkflowExecution testworkflowsv1.TestWorkflowExecution
	err := r.Get(ctx, req.NamespacedName, &testWorkflowExecution)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	if testWorkflowExecution.Spec.TestWorkflow == nil {
		return ctrl.Result{}, nil
	}

	if testWorkflowExecution.Spec.ExecutionRequest == nil {
		testWorkflowExecution.Spec.ExecutionRequest = &testworkflowsv1.TestWorkflowExecutionRequest{}
	}

	interface_ := testworkflowsv1.API_TestWorkflowRunningContextInterface
	actor := testworkflowsv1.TESTWORKFLOWEXECUTION_TestWorkflowRunningContextActor
	callerResourceType := testworkflowsv1.TESTWORKFLOWEXECUTION_TestWorkflowRunningContextCallerResourceType
	testWorkflowExecution.Spec.ExecutionRequest.RunningContext = []testworkflowsv1.TestWorkflowRunningContext{
		{
			Interface_: &interface_,
			Actor:      &actor,
			Caller: &testworkflowsv1.TestWorkflowRunningContextCaller{
				CallerResourceType: &callerResourceType,
				CallerResourceName: testWorkflowExecution.Name,
			},
		},
	}

	jsonData, err := json.Marshal(testWorkflowExecution.Spec.ExecutionRequest)
	if err != nil {
		return ctrl.Result{}, err
	}

	if testWorkflowExecution.Generation == testWorkflowExecution.Status.Generation {
		return ctrl.Result{}, nil
	}

	if _, err = r.executeTestWorkflow(testWorkflowExecution.Spec.TestWorkflow.Name, testWorkflowExecution.Name, testWorkflowExecution.Namespace, jsonData); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestWorkflowExecutionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	pred := predicate.GenerationChangedPredicate{}
	return ctrl.NewControllerManagedBy(mgr).
		For(&testworkflowsv1.TestWorkflowExecution{}).
		WithEventFilter(pred).
		Complete(r)
}

func (r *TestWorkflowExecutionReconciler) executeTestWorkflow(testWorkflowName, testWorkflowExecutionName, namespace string, jsonData []byte) (out string, err error) {
	request, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("http://%s.%s.svc.cluster.local:%d/v1/test-workflows/%s/executions?testWorkflowExecutionName=%s",
			r.ServiceName, namespace, r.ServicePort, testWorkflowName, testWorkflowExecutionName),
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
