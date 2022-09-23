package client

import (
	executorv1 "github.com/kubeshop/testkube-operator/apis/executor/v1"
	scriptv1 "github.com/kubeshop/testkube-operator/apis/script/v1"
	scriptv2 "github.com/kubeshop/testkube-operator/apis/script/v2"
	testsv1 "github.com/kubeshop/testkube-operator/apis/tests/v1"
	testsv2 "github.com/kubeshop/testkube-operator/apis/tests/v2"
	testsv3 "github.com/kubeshop/testkube-operator/apis/tests/v3"
	testsourcev1 "github.com/kubeshop/testkube-operator/apis/testsource/v1"
	testsuitev1 "github.com/kubeshop/testkube-operator/apis/testsuite/v1"
	testsuitev2 "github.com/kubeshop/testkube-operator/apis/testsuite/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetClient returns kubernetes CRD client with registered schemes
func GetClient() (client.Client, error) {
	scheme := runtime.NewScheme()

	scriptv1.AddToScheme(scheme)
	scriptv2.AddToScheme(scheme)
	executorv1.AddToScheme(scheme)
	testsv1.AddToScheme(scheme)
	testsv2.AddToScheme(scheme)
	testsv3.AddToScheme(scheme)
	testsuitev1.AddToScheme(scheme)
	corev1.AddToScheme(scheme)
	testsuitev2.AddToScheme(scheme)
	testsourcev1.AddToScheme(scheme)

	kubeconfig, err := ctrl.GetConfig()
	if err != nil {
		return nil, err
	}

	return client.New(kubeconfig, client.Options{Scheme: scheme})
}
