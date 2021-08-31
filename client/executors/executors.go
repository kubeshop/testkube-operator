package scripts

import (
	"context"
	"fmt"

	executorsAPI "github.com/kubeshop/kubtest-operator/apis/executor/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClient(client client.Client) *ExecutorsClient {
	return &ExecutorsClient{
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

func (s ExecutorsClient) GetByType(executorType string) (*executorsAPI.Executor, error) {
	list := &executorsAPI.ExecutorList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{})
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, exec := range list.Items {
		names = append(names, fmt.Sprintf("%s/%s", exec.Namespace, exec.Name))
		for _, t := range exec.Spec.Types {
			if t == executorType {
				return &exec, nil
			}
		}
	}

	return nil, fmt.Errorf("executor type '%s' is not handled by any of executors (%s)", executorType, names)
}

func (s ExecutorsClient) Create(scripts *executorsAPI.Executor) (*executorsAPI.Executor, error) {
	err := s.Client.Create(context.Background(), scripts)
	return scripts, err
}
