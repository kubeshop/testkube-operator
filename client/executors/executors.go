package scripts

import (
	"context"
	"fmt"

	executorsAPI "github.com/kubeshop/kubtest-operator/apis/executor/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewClient returns new client instance, needs kubernetes client to be passed as dependecy
func NewClient(client client.Client) *ExecutorsClient {
	return &ExecutorsClient{
		Client: client,
	}
}

// ExecutorsClient client for getting executors CRs
type ExecutorsClient struct {
	Client client.Client
}

// List shows list of available executors
func (s ExecutorsClient) List(namespace string) (*executorsAPI.ExecutorList, error) {
	list := &executorsAPI.ExecutorList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	return list, err
}

// Get gets executor by name in given namespace
func (s ExecutorsClient) Get(namespace, name string) (*executorsAPI.Executor, error) {
	script := &executorsAPI.Executor{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: name}, script)
	return script, err
}

// GetByType gets first available executor for given type
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

// Create creates new Executor CR
func (s ExecutorsClient) Create(scripts *executorsAPI.Executor) (*executorsAPI.Executor, error) {
	err := s.Client.Create(context.Background(), scripts)
	return scripts, err
}
