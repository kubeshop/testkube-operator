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
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	testsv3 "github.com/kubeshop/testkube-operator/apis/tests/v3"
	"github.com/kubeshop/testkube-operator/pkg/config"
	"github.com/kubeshop/testkube-operator/pkg/cronjob"
	"github.com/kubeshop/testkube-operator/pkg/telemetry"
)

// TestReconciler reconciles a Test object
type TestReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	CronJobClient *cronjob.Client
	ConfigMap     config.Repository
	AppVersion    string
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
	logger := log.FromContext(ctx)

	// Delete CronJob if it was created for deleted Test
	var test testsv3.Test
	err := r.Get(ctx, req.NamespacedName, &test)
	if err != nil {
		if errors.IsNotFound(err) {
			if err = r.CronJobClient.Delete(ctx,
				cronjob.GetMetadataName(req.NamespacedName.Name, cronjob.TestResourceURI), req.NamespacedName.Namespace); err != nil {
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	if test.Generation == 1 {
		telemetryEnabled, err := r.ConfigMap.GetTelemetryEnabled(ctx)
		if err != nil {
			logger.Error(err, "getting telemetry enabled error", "error")
		}

		if telemetryEnabled {
			clusterID, err := r.ConfigMap.GetUniqueClusterId(ctx)
			if err != nil {
				logger.Error(err, "getting cluster id error", "error")
			}

			host, err := os.Hostname()
			if err != nil {
				logger.Error(err, "getting hostname error")
			}

			var dataSource string
			if test.Spec.Content != nil {
				dataSource = test.Spec.Content.Type_
			}

			out, err := telemetry.SendCreateEvent("testkube_api_create_test", telemetry.CreateParams{
				AppVersion: r.AppVersion,
				DataSource: dataSource,
				Host:       host,
				ClusterID:  clusterID,
				TestType:   test.Spec.Type_,
			})
			if err != nil {
				logger.Error(err, "sending create test telemetry event error")
			} else {
				logger.Info("sending create test telemetry event", "output", out)
			}
		}
	}

	// Delete CronJob if it was created for cleaned Test schedule
	if test.Spec.Schedule == "" {
		if err := r.CronJobClient.Delete(ctx,
			cronjob.GetMetadataName(req.NamespacedName.Name, cronjob.TestResourceURI), req.NamespacedName.Namespace); err != nil {
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	data, err := json.Marshal(testsv3.ExecutionRequest{})
	if err != nil {
		return ctrl.Result{}, err
	}

	options := cronjob.CronJobOptions{
		Schedule: test.Spec.Schedule,
		Resource: cronjob.TestResourceURI,
		Data:     string(data),
		Labels:   test.Labels,
	}

	// Create CronJob if it was not created before for provided Test schedule
	cronJob, err := r.CronJobClient.Get(ctx,
		cronjob.GetMetadataName(req.NamespacedName.Name, cronjob.TestResourceURI), req.NamespacedName.Namespace)
	if err != nil {
		if errors.IsNotFound(err) {
			if err = r.CronJobClient.Create(ctx, test.Name,
				cronjob.GetMetadataName(test.Name, cronjob.TestResourceURI), req.NamespacedName.Namespace, options); err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, err
	}

	// Update CronJob if it was created before provided Test schedule
	if err = r.CronJobClient.Update(ctx, cronJob, test.Name,
		cronjob.GetMetadataName(test.Name, cronjob.TestResourceURI), req.NamespacedName.Namespace, options); err != nil {
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
