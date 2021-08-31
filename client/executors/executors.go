package scripts

import (
	"context"

	executorsAPI "github.com/kubeshop/kubtest-operator/apis/executor/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClient(client client.Client) ExecutorsClient {
	return ExecutorsClient{
		Client: client,
	}
}

type ExecutorsClient struct {
	Client client.Client
}

func (s ExecutorsClient) List(namespace string) (*executorsAPI.ExecutorList, error) {
	list := &executorsAPI.ExecutorList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	return list, err
}

func (s ExecutorsClient) Get(namespace, name string) (*executorsAPI.Executor, error) {
	script := &executorsAPI.Executor{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: name}, script)
	return script, err
}

func (s ExecutorsClient) GetByType(scriptType string) (*executorsAPI.Executor, error) {
	list := &executorsAPI.ExecutorList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{})

	for _, script := range list.Items {
		if script.Spec.Type_ == scriptType {
			return script, nil
		}

	}

	return script, err
}

func (s ExecutorsClient) Create(scripts *executorsAPI.Executor) (*executorsAPI.Executor, error) {
	err := s.Client.Create(context.Background(), scripts)
	return scripts, err
}
