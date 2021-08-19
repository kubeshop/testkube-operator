package client

import (
	"log"

	executorv1 "github.com/kubeshop/kubtest-operator/apis/executor/v1"
	scriptv1 "github.com/kubeshop/kubtest-operator/apis/script/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GetClient returns kubernetes CRD client with registered schemes
// TODO I don't like thoses deps - rethink structure and refactor
// client is included in ScriptClient (but client need to know about script :/)
func GetClient() client.Client {
	scheme := runtime.NewScheme()

	scriptv1.AddToScheme(scheme)
	executorv1.AddToScheme(scheme)

	kubeconfig := ctrl.GetConfigOrDie()
	controllerClient, err := client.New(kubeconfig, client.Options{Scheme: scheme})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return controllerClient
}
