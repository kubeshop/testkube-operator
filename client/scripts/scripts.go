package scripts

import (
	"context"

	scriptsAPI "github.com/kubeshop/kubtest-operator/apis/script/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClient(client client.Client) *ScriptsClient {
	return &ScriptsClient{
		Client: client,
	}
}

type ScriptsClient struct {
	Client client.Client
}

func (s ScriptsClient) List(namespace string) (*scriptsAPI.ScriptList, error) {
	list := &scriptsAPI.ScriptList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	return list, err
}

func (s ScriptsClient) Get(namespace, name string) (*scriptsAPI.Script, error) {
	script := &scriptsAPI.Script{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: name}, script)
	return script, err
}

func (s ScriptsClient) Create(scripts *scriptsAPI.Script) (*scriptsAPI.Script, error) {
	err := s.Client.Create(context.Background(), scripts)
	return scripts, err
}
