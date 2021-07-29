package main

import (
	"context"
	"fmt"
	"log"

	scriptsAPI "github.com/kubeshop/kubetest/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var kclient client.Client

func init() {
	kclient = GetClient()
}

func main() {
	r, err := List("default")
	fmt.Printf("%+v\n", r)
	fmt.Printf("%+v\n", err)

}

func GetClient() client.Client {
	scheme := runtime.NewScheme()
	scriptsAPI.AddToScheme(scheme)
	kubeconfig := ctrl.GetConfigOrDie()
	controllerClient, err := client.New(kubeconfig, client.Options{Scheme: scheme})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return controllerClient
}

func List(namespace string) (result *scriptsAPI.ScriptList, err error) {
	list := &scriptsAPI.ScriptList{}
	err = kclient.List(context.TODO(), list, &client.ListOptions{Namespace: namespace})
	return list, err
}

func Create(deployment *scriptsAPI.Script) (sdep *scriptsAPI.Script, err error) {
	err = kclient.Create(context.TODO(), deployment)
	return deployment, err
}
