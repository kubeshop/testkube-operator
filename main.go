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

package main

import (
	"encoding/base64"
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.

	zapUber "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/kelseyhightower/envconfig"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	executorv1 "github.com/kubeshop/testkube-operator/apis/executor/v1"
	testkubev1 "github.com/kubeshop/testkube-operator/apis/script/v1"
	testkubev2 "github.com/kubeshop/testkube-operator/apis/script/v2"
	testsv1 "github.com/kubeshop/testkube-operator/apis/tests/v1"
	testsv2 "github.com/kubeshop/testkube-operator/apis/tests/v2"
	testsv3 "github.com/kubeshop/testkube-operator/apis/tests/v3"
	testsuitev1 "github.com/kubeshop/testkube-operator/apis/testsuite/v1"
	executorcontrollers "github.com/kubeshop/testkube-operator/controllers/executor"
	scriptcontrollers "github.com/kubeshop/testkube-operator/controllers/script"
	testscontrollers "github.com/kubeshop/testkube-operator/controllers/tests"
	testsuitecontrollers "github.com/kubeshop/testkube-operator/controllers/testsuite"
	"github.com/kubeshop/testkube-operator/pkg/cronjob"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

// config for HTTP server
type config struct {
	Port            int
	Fullname        string
	TemplateCronjob string
}

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(testkubev1.AddToScheme(scheme))
	utilruntime.Must(executorv1.AddToScheme(scheme))
	utilruntime.Must(testsv1.AddToScheme(scheme))
	utilruntime.Must(testkubev2.AddToScheme(scheme))
	utilruntime.Must(testsuitev1.AddToScheme(scheme))
	utilruntime.Must(testsv2.AddToScheme(scheme))
	utilruntime.Must(testsv3.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	setLogger()

	var httpConfig config
	err := envconfig.Process("APISERVER", &httpConfig)
	// Do we want to panic here or just ignore the err
	if err != nil {
		panic(err)
	}

	var templateCronjob string
	if httpConfig.TemplateCronjob != "" {
		data, err := base64.StdEncoding.DecodeString(httpConfig.TemplateCronjob)
		if err != nil {
			panic(err)
		}

		templateCronjob = string(data)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "47f0dfc1.testkube.io",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&scriptcontrollers.ScriptReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Script")
		os.Exit(1)
	}
	if err = (&executorcontrollers.ExecutorReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Executor")
		os.Exit(1)
	}
	if err = (&testscontrollers.TestReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		CronJobClient: cronjob.NewClient(mgr.GetClient(), httpConfig.Fullname, httpConfig.Port, templateCronjob),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Test")
		os.Exit(1)
	}
	if err = (&testsuitecontrollers.TestSuiteReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TestSuite")
		os.Exit(1)
	}

	if err = (&executorcontrollers.WebhookReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Webhook")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	if os.Getenv("ENABLE_WEBHOOKS") != "false" {
		if err = (&testkubev1.Script{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Script")
			os.Exit(1)
		}
		if err = (&testkubev2.Script{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Script")
			os.Exit(1)
		}
		if err = (&testsv2.Test{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Test")
			os.Exit(1)
		}
		if err = (&testsv3.Test{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Test")
			os.Exit(1)
		}
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

// setLogger sets up the zap logger to print error, panic and fatal messages to stderr and lower level messages to stdout
func setLogger() {
	highPriority := zapUber.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zapUber.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	consoleEncoder := zapcore.NewConsoleEncoder(zapUber.NewDevelopmentEncoderConfig())

	updateZapcore := func(c zapcore.Core) zapcore.Core {
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		)
		return core
	}

	var opts zap.Options
	opts.ZapOpts = append(opts.ZapOpts, zapUber.WrapCore(updateZapcore))

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
}
