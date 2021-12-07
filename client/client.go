package client

import (
	executorv1 "github.com/kubeshop/testkube-operator/apis/executor/v1"
	scriptv1 "github.com/kubeshop/testkube-operator/apis/script/v1"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetClient returns kubernetes CRD client with registered schemes
// GetClient returns kubernetes CRD client with registered schemes
func GetClient() (client.Client, error) {
	scheme := runtime.NewScheme()

	scriptv1.AddToScheme(scheme)
	executorv1.AddToScheme(scheme)

	kubeconfig, err := ctrl.GetConfig()
	if err != nil {
		return nil, err
	}

	return client.New(kubeconfig, client.Options{Scheme: scheme})
}
