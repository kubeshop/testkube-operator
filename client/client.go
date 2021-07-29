package client

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetClient returns kubernetes CRD client with registered schemes
func GetClient(crdClients ...CRDClient) client.Client {
	scheme := runtime.NewScheme()
	for _, client := range crdClients {
		client.Register(scheme)
	}
	kubeconfig := ctrl.GetConfigOrDie()
	controllerClient, err := client.New(kubeconfig, client.Options{Scheme: scheme})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return controllerClient
}

type CRDClient interface {
	Register(scheme *runtime.Scheme)
}
