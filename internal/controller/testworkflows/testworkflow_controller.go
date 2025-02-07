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
	"encoding/json"
	"fmt"
	"maps"
	"net/http"

	testworkflowsv1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"
	"github.com/kubeshop/testkube-operator/pkg/cronjob"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// TestWorkflowReconciler reconciles a TestWorkflow object
type TestWorkflowReconciler struct {
	client.Client
	Scheme          *runtime.Scheme
	CronJobClient   *cronjob.Client
	ServiceName     string
	ServicePort     int
	PurgeExecutions bool
}

//+kubebuilder:rbac:groups=testworkflows.testkube.io,resources=testworkflows,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=testworkflows.testkube.io,resources=testworkflows/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=testworkflows.testkube.io,resources=testworkflows/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TestWorkflow object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *TestWorkflowReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Delete CronJobs if it were created for deleted Test Workflow
	var testWorkflow testworkflowsv1.TestWorkflow
	if err := r.Get(ctx, req.NamespacedName, &testWorkflow); err != nil {
		if errors.IsNotFound(err) {
			if err = r.CronJobClient.DeleteAll(ctx,
				cronjob.GetSelector(req.NamespacedName.Name, cronjob.TestWorkflowResourceURI), req.NamespacedName.Namespace); err != nil {
				return ctrl.Result{}, err
			}

			if _, err = r.deleteTestWorkflow(req.NamespacedName.Name, req.NamespacedName.Namespace); err != nil {
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	events := testWorkflow.Spec.Events
	for _, template := range testWorkflow.Spec.Use {
		var testWorkflowTemplate testworkflowsv1.TestWorkflowTemplate
		if err := r.Get(ctx, types.NamespacedName{Namespace: testWorkflow.Namespace, Name: template.Name}, &testWorkflowTemplate); err != nil {
			if errors.IsNotFound(err) {
				continue
			}

			return ctrl.Result{}, err
		}

		events = append(events, testWorkflowTemplate.Spec.Events...)
	}

	hasTemplates := len(testWorkflow.Spec.Use) != 0
	_, ok := testWorkflow.Labels[cronjob.TestWorkflowTemplateResourceURI]
	if ok && !hasTemplates {
		delete(testWorkflow.Labels, cronjob.TestWorkflowTemplateResourceURI)
		return ctrl.Result{}, r.Update(ctx, &testWorkflow)
	}

	if !ok && hasTemplates {
		if testWorkflow.Labels == nil {
			testWorkflow.Labels = make(map[string]string)
		}

		testWorkflow.Labels[cronjob.TestWorkflowTemplateResourceURI] = "yes"
		return ctrl.Result{}, r.Update(ctx, &testWorkflow)
	}

	newCronJobConfigs := make(map[string]*testworkflowsv1.CronJobConfig)
	oldCronJobs := make(map[string]*batchv1.CronJob)
	cronJobList, err := r.CronJobClient.ListAll(ctx,
		cronjob.GetSelector(testWorkflow.Name, cronjob.TestWorkflowResourceURI), testWorkflow.Namespace)
	if err != nil {
		return ctrl.Result{}, err
	}

	for i := range cronJobList.Items {
		oldCronJobs[cronJobList.Items[i].Name] = &cronJobList.Items[i]
	}

	for _, event := range events {
		if event.Cronjob != nil {
			name, err := cronjob.GetHashedMetadataName(testWorkflow.Name, event.Cronjob.Cron, string(testWorkflow.UID), event.Cronjob.Config)
			if err != nil {
				return ctrl.Result{}, err
			}

			if cronJob, ok := newCronJobConfigs[name]; !ok {
				newCronJobConfigs[name] = &testworkflowsv1.CronJobConfig{
					Cron:        event.Cronjob.Cron,
					Labels:      event.Cronjob.Labels,
					Annotations: event.Cronjob.Annotations,
					Config:      event.Cronjob.Config,
				}
			} else {
				newCronJobConfigs[name] = MergeCronJobJobConfig(cronJob, event.Cronjob)
			}
		}
	}

	interface_ := testworkflowsv1.API_TestWorkflowRunningContextInterfaceType
	actor := testworkflowsv1.CRON_TestWorkflowRunningContextActorType
	request := testworkflowsv1.TestWorkflowExecutionRequest{
		RunningContext: &testworkflowsv1.TestWorkflowRunningContext{
			Interface_: &testworkflowsv1.TestWorkflowRunningContextInterface{
				Type_: &interface_,
			},
			Actor: &testworkflowsv1.TestWorkflowRunningContextActor{
				Type_: &actor,
			},
		},
	}

	for name, oldCronJob := range oldCronJobs {
		if newCronJobConfig, ok := newCronJobConfigs[name]; !ok {
			// Delete removed Cron Jobs
			if err = r.CronJobClient.Delete(ctx, name, testWorkflow.Namespace); err != nil {
				return ctrl.Result{}, err
			}
		} else {
			// Update CronJob if it was created before
			if newCronJobConfig.Labels == nil {
				newCronJobConfig.Labels = make(map[string]string)
			}

			request.Config = newCronJobConfig.Config
			data, err := json.Marshal(request)
			if err != nil {
				return ctrl.Result{}, err
			}

			newCronJobConfig.Labels[cronjob.TestWorkflowResourceURI] = testWorkflow.Name
			options := cronjob.CronJobOptions{
				Schedule:    newCronJobConfig.Cron,
				Group:       testworkflowsv1.Group,
				Resource:    testworkflowsv1.Resource,
				Version:     testworkflowsv1.Version,
				ResourceURI: cronjob.TestWorkflowResourceURI,
				Labels:      newCronJobConfig.Labels,
				Annotations: newCronJobConfig.Annotations,
				Data:        string(data),
			}

			if err = r.CronJobClient.Update(ctx, oldCronJob, testWorkflow.Name, name, testWorkflow.Namespace,
				string(testWorkflow.UID), options); err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	for name, newCronJobConfig := range newCronJobConfigs {
		if _, ok = oldCronJobs[name]; !ok {
			// Create new Cron Jobs
			if newCronJobConfig.Labels == nil {
				newCronJobConfig.Labels = make(map[string]string)
			}

			request.Config = newCronJobConfig.Config
			data, err := json.Marshal(request)
			if err != nil {
				return ctrl.Result{}, err
			}

			newCronJobConfig.Labels[cronjob.TestWorkflowResourceURI] = testWorkflow.Name
			options := cronjob.CronJobOptions{
				Schedule:    newCronJobConfig.Cron,
				Group:       testworkflowsv1.Group,
				Resource:    testworkflowsv1.Resource,
				Version:     testworkflowsv1.Version,
				ResourceURI: cronjob.TestWorkflowResourceURI,
				Labels:      newCronJobConfig.Labels,
				Annotations: newCronJobConfig.Annotations,
				Data:        string(data),
			}

			if err = r.CronJobClient.Create(ctx, testWorkflow.Name, name, testWorkflow.Namespace,
				string(testWorkflow.UID), options); err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

func MergeCronJobJobConfig(dst, include *testworkflowsv1.CronJobConfig) *testworkflowsv1.CronJobConfig {
	if dst == nil {
		return include
	} else if include == nil {
		return dst
	}

	if len(include.Labels) > 0 && dst.Labels == nil {
		dst.Labels = map[string]string{}
	}
	maps.Copy(dst.Labels, include.Labels)

	if len(include.Annotations) > 0 && dst.Annotations == nil {
		dst.Annotations = map[string]string{}
	}
	maps.Copy(dst.Annotations, include.Annotations)

	return dst
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestWorkflowReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testworkflowsv1.TestWorkflow{}).
		Complete(r)
}

func (r *TestWorkflowReconciler) deleteTestWorkflow(testWorkflowName, namespace string) (out string, err error) {
	if !r.PurgeExecutions {
		return out, nil
	}

	request, err := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("http://%s.%s.svc.cluster.local:%d/v1/test-workflows/%s?skipDeleteCRD=true",
			r.ServiceName, namespace, r.ServicePort, testWorkflowName), nil)
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
