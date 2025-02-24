package main

import (
	"crypto/tls"
	"encoding/base64"
	"flag"
	"os"
	"path/filepath"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	zapUber "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/certwatcher"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/metrics/filters"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"github.com/kelseyhightower/envconfig"
	testtriggersv1 "github.com/kubeshop/testkube-operator/api/testtriggers/v1"
	testworkflowsv1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"

	executorv1 "github.com/kubeshop/testkube-operator/api/executor/v1"
	testkubev1 "github.com/kubeshop/testkube-operator/api/script/v1"
	testkubev2 "github.com/kubeshop/testkube-operator/api/script/v2"
	templatev1 "github.com/kubeshop/testkube-operator/api/template/v1"
	testexecutionv1 "github.com/kubeshop/testkube-operator/api/testexecution/v1"
	testsv1 "github.com/kubeshop/testkube-operator/api/tests/v1"
	testsv2 "github.com/kubeshop/testkube-operator/api/tests/v2"
	testsv3 "github.com/kubeshop/testkube-operator/api/tests/v3"
	testsourcev1 "github.com/kubeshop/testkube-operator/api/testsource/v1"
	testsuitev1 "github.com/kubeshop/testkube-operator/api/testsuite/v1"
	testsuitev2 "github.com/kubeshop/testkube-operator/api/testsuite/v2"
	testsuitev3 "github.com/kubeshop/testkube-operator/api/testsuite/v3"
	testsuiteexecutionv1 "github.com/kubeshop/testkube-operator/api/testsuiteexecution/v1"
	executorcontrollers "github.com/kubeshop/testkube-operator/internal/controller/executor"
	scriptcontrollers "github.com/kubeshop/testkube-operator/internal/controller/script"
	templatecontrollers "github.com/kubeshop/testkube-operator/internal/controller/template"
	testexecutioncontrollers "github.com/kubeshop/testkube-operator/internal/controller/testexecution"
	testscontrollers "github.com/kubeshop/testkube-operator/internal/controller/tests"
	testsourcecontrollers "github.com/kubeshop/testkube-operator/internal/controller/testsource"
	testsuitecontrollers "github.com/kubeshop/testkube-operator/internal/controller/testsuite"
	testsuiteexecutioncontrollers "github.com/kubeshop/testkube-operator/internal/controller/testsuiteexecution"
	testtriggerscontrollers "github.com/kubeshop/testkube-operator/internal/controller/testtriggers"
	testworkflowexecutioncontrollers "github.com/kubeshop/testkube-operator/internal/controller/testworkflowexecution"
	testworkflowscontrollers "github.com/kubeshop/testkube-operator/internal/controller/testworkflows"
	cronjobclient "github.com/kubeshop/testkube-operator/pkg/cronjob/client"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(testkubev1.AddToScheme(scheme))
	utilruntime.Must(executorv1.AddToScheme(scheme))
	utilruntime.Must(testsv1.AddToScheme(scheme))
	utilruntime.Must(testkubev2.AddToScheme(scheme))
	utilruntime.Must(testsuitev1.AddToScheme(scheme))
	utilruntime.Must(testsv2.AddToScheme(scheme))
	utilruntime.Must(testsv3.AddToScheme(scheme))
	utilruntime.Must(testsuitev2.AddToScheme(scheme))
	utilruntime.Must(testtriggersv1.AddToScheme(scheme))
	utilruntime.Must(testsourcev1.AddToScheme(scheme))
	utilruntime.Must(testsuitev3.AddToScheme(scheme))
	utilruntime.Must(testexecutionv1.AddToScheme(scheme))
	utilruntime.Must(testsuiteexecutionv1.AddToScheme(scheme))
	utilruntime.Must(templatev1.AddToScheme(scheme))
	utilruntime.Must(testworkflowsv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

// HttpServerConfig for HTTP server
type HttpServerConfig struct {
	Port            int
	Fullname        string
	TemplateCronjob string `split_words:"true"`
	Registry        string
	UseArgocdSync   bool `split_words:"true"`
	PurgeExecutions bool `split_words:"true"`
	Config          string
}

// nolint:gocyclo
func main() {
	var metricsAddr string
	var metricsCertPath, metricsCertName, metricsCertKey string
	var webhookCertPath, webhookCertName, webhookCertKey string
	var enableLeaderElection bool
	var probeAddr string
	var secureMetrics bool
	var enableHTTP2 bool
	var tlsOpts []func(*tls.Config)
	flag.StringVar(&metricsAddr, "metrics-bind-address", "0", "The address the metrics endpoint binds to. "+
		"Use :8443 for HTTPS or :8080 for HTTP, or leave as 0 to disable the metrics service.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.BoolVar(&secureMetrics, "metrics-secure", true,
		"If set, the metrics endpoint is served securely via HTTPS. Use --metrics-secure=false to use HTTP instead.")
	flag.StringVar(&webhookCertPath, "webhook-cert-path", "", "The directory that contains the webhook certificate.")
	flag.StringVar(&webhookCertName, "webhook-cert-name", "tls.crt", "The name of the webhook certificate file.")
	flag.StringVar(&webhookCertKey, "webhook-cert-key", "tls.key", "The name of the webhook key file.")
	flag.StringVar(&metricsCertPath, "metrics-cert-path", "",
		"The directory that contains the metrics server certificate.")
	flag.StringVar(&metricsCertName, "metrics-cert-name", "tls.crt", "The name of the metrics server certificate file.")
	flag.StringVar(&metricsCertKey, "metrics-cert-key", "tls.key", "The name of the metrics server key file.")
	flag.BoolVar(&enableHTTP2, "enable-http2", false,
		"If set, HTTP/2 will be enabled for the metrics and webhook servers")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	setLogger()

	// if the enable-http2 flag is false (the default), http/2 should be disabled
	// due to its vulnerabilities. More specifically, disabling http/2 will
	// prevent from being vulnerable to the HTTP/2 Stream Cancellation and
	// Rapid Reset CVEs. For more information see:
	// - https://github.com/advisories/GHSA-qppj-fm5r-hxr3
	// - https://github.com/advisories/GHSA-4374-p667-p6c8
	disableHTTP2 := func(c *tls.Config) {
		setupLog.Info("disabling http/2")
		c.NextProtos = []string{"http/1.1"}
	}

	if !enableHTTP2 {
		tlsOpts = append(tlsOpts, disableHTTP2)
	}

	// Create watchers for metrics and webhooks certificates
	var metricsCertWatcher, webhookCertWatcher *certwatcher.CertWatcher

	// Initial webhook TLS options
	webhookTLSOpts := tlsOpts

	if len(webhookCertPath) > 0 {
		setupLog.Info("Initializing webhook certificate watcher using provided certificates",
			"webhook-cert-path", webhookCertPath, "webhook-cert-name", webhookCertName, "webhook-cert-key", webhookCertKey)

		var err error
		webhookCertWatcher, err = certwatcher.New(
			filepath.Join(webhookCertPath, webhookCertName),
			filepath.Join(webhookCertPath, webhookCertKey),
		)
		if err != nil {
			setupLog.Error(err, "Failed to initialize webhook certificate watcher")
			os.Exit(1)
		}

		webhookTLSOpts = append(webhookTLSOpts, func(config *tls.Config) {
			config.GetCertificate = webhookCertWatcher.GetCertificate
		})
	}

	webhookServer := webhook.NewServer(webhook.Options{
		TLSOpts: webhookTLSOpts,
	})

	// Metrics endpoint is enabled in 'config/default/kustomization.yaml'. The Metrics options configure the server.
	// More info:
	// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.1/pkg/metrics/server
	// - https://book.kubebuilder.io/reference/metrics.html
	metricsServerOptions := metricsserver.Options{
		BindAddress:   metricsAddr,
		SecureServing: secureMetrics,
		TLSOpts:       tlsOpts,
	}

	if secureMetrics {
		// FilterProvider is used to protect the metrics endpoint with authn/authz.
		// These configurations ensure that only authorized users and service accounts
		// can access the metrics endpoint. The RBAC are configured in 'config/rbac/kustomization.yaml'. More info:
		// https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.1/pkg/metrics/filters#WithAuthenticationAndAuthorization
		metricsServerOptions.FilterProvider = filters.WithAuthenticationAndAuthorization
	}

	// If the certificate is not specified, controller-runtime will automatically
	// generate self-signed certificates for the metrics server. While convenient for development and testing,
	// this setup is not recommended for production.
	//
	// TODO(user): If you enable certManager, uncomment the following lines:
	// - [METRICS-WITH-CERTS] at config/default/kustomization.yaml to generate and use certificates
	// managed by cert-manager for the metrics server.
	// - [PROMETHEUS-WITH-CERTS] at config/prometheus/kustomization.yaml for TLS certification.
	if len(metricsCertPath) > 0 {
		setupLog.Info("Initializing metrics certificate watcher using provided certificates",
			"metrics-cert-path", metricsCertPath, "metrics-cert-name", metricsCertName, "metrics-cert-key", metricsCertKey)

		var err error
		metricsCertWatcher, err = certwatcher.New(
			filepath.Join(metricsCertPath, metricsCertName),
			filepath.Join(metricsCertPath, metricsCertKey),
		)
		if err != nil {
			setupLog.Error(err, "to initialize metrics certificate watcher", "error", err)
			os.Exit(1)
		}

		metricsServerOptions.TLSOpts = append(metricsServerOptions.TLSOpts, func(config *tls.Config) {
			config.GetCertificate = metricsCertWatcher.GetCertificate
		})
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		Metrics:                metricsServerOptions,
		WebhookServer:          webhookServer,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "47f0dfc1.testkube.io",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	var httpConfig HttpServerConfig
	err = envconfig.Process("APISERVER", &httpConfig)
	// TODO: Do we want to panic here or just ignore the error?
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
	cronJobClient := cronjobclient.New(mgr.GetClient(), httpConfig.Fullname, httpConfig.Port,
		templateCronjob, httpConfig.Registry, httpConfig.UseArgocdSync)
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
		Client:          mgr.GetClient(),
		Scheme:          mgr.GetScheme(),
		CronJobClient:   cronJobClient,
		ServiceName:     httpConfig.Fullname,
		ServicePort:     httpConfig.Port,
		PurgeExecutions: httpConfig.PurgeExecutions,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Test")
		os.Exit(1)
	}
	if err = (&testsuitecontrollers.TestSuiteReconciler{
		Client:          mgr.GetClient(),
		Scheme:          mgr.GetScheme(),
		CronJobClient:   cronJobClient,
		ServiceName:     httpConfig.Fullname,
		ServicePort:     httpConfig.Port,
		PurgeExecutions: httpConfig.PurgeExecutions,
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
	if err = (&executorcontrollers.WebhookTemplateReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "WebhookTemplate")
		os.Exit(1)
	}
	if err = (&testsourcecontrollers.TestSourceReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TestSource")
		os.Exit(1)
	}
	if err = (&testtriggerscontrollers.TestTriggerReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TestTrigger")
		os.Exit(1)
	}
	if err = (&testexecutioncontrollers.TestExecutionReconciler{
		Client:      mgr.GetClient(),
		Scheme:      mgr.GetScheme(),
		ServiceName: httpConfig.Fullname,
		ServicePort: httpConfig.Port,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TestExecution")
		os.Exit(1)
	}
	if err = (&testsuiteexecutioncontrollers.TestSuiteExecutionReconciler{
		Client:      mgr.GetClient(),
		Scheme:      mgr.GetScheme(),
		ServiceName: httpConfig.Fullname,
		ServicePort: httpConfig.Port,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TestSuiteExecution")
		os.Exit(1)
	}
	if err = (&templatecontrollers.TemplateReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Template")
		os.Exit(1)
	}
	if err = (&testworkflowscontrollers.TestWorkflowReconciler{
		Client:          mgr.GetClient(),
		Scheme:          mgr.GetScheme(),
		CronJobClient:   cronJobClient,
		ServiceName:     httpConfig.Fullname,
		ServicePort:     httpConfig.Port,
		PurgeExecutions: httpConfig.PurgeExecutions,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TestWorkflow")
		os.Exit(1)
	}
	if err = (&testworkflowscontrollers.TestWorkflowTemplateReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TestWorkflowTemplate")
		os.Exit(1)
	}
	if err = (&testworkflowexecutioncontrollers.TestWorkflowExecutionReconciler{
		Client:      mgr.GetClient(),
		Scheme:      mgr.GetScheme(),
		ServiceName: httpConfig.Fullname,
		ServicePort: httpConfig.Port,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TestWorkflowExecution")
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
		if err = (&testsv1.Test{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Test")
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
		if err = (&testsuitev1.TestSuite{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "TestSuite")
			os.Exit(1)
		}
		if err = (&testsuitev2.TestSuite{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "TestSuite")
			os.Exit(1)
		}
		if err = (&testsuitev3.TestSuite{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "TestSuite")
			os.Exit(1)
		}
		testtriggerValidator := testtriggerscontrollers.NewValidator(mgr.GetClient())
		if err = (&testtriggersv1.TestTrigger{}).SetupWebhookWithManager(mgr, testtriggerValidator); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "TestTrigger")
			os.Exit(1)
		}
	} else {
		setupLog.Info("Webhooks are disabled")
	}
	// +kubebuilder:scaffold:builder

	if metricsCertWatcher != nil {
		setupLog.Info("Adding metrics certificate watcher to manager")
		if err := mgr.Add(metricsCertWatcher); err != nil {
			setupLog.Error(err, "unable to add metrics certificate watcher to manager")
			os.Exit(1)
		}
	}

	if webhookCertWatcher != nil {
		setupLog.Info("Adding webhook certificate watcher to manager")
		if err := mgr.Add(webhookCertWatcher); err != nil {
			setupLog.Error(err, "unable to add webhook certificate watcher to manager")
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

	jsonEncoder := zapcore.NewJSONEncoder(zapUber.NewDevelopmentEncoderConfig())

	updateZapcore := func(c zapcore.Core) zapcore.Core {
		core := zapcore.NewTee(
			zapcore.NewCore(jsonEncoder, consoleErrors, highPriority),
			zapcore.NewCore(jsonEncoder, consoleDebugging, lowPriority),
		)
		return core
	}

	var opts zap.Options
	opts.ZapOpts = append(opts.ZapOpts, zapUber.WrapCore(updateZapcore))

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
}
