package scripts

import (
	"context"

	scriptsAPI "github.com/kubeshop/testkube-operator/apis/script/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (s ScriptsClient) Delete(namespace, name string) error {
	script, err := s.Get(namespace, name)
	if err != nil {
		return err
	}

	err = s.Client.Delete(context.Background(), script)
	return err
}

func (s ScriptsClient) DeleteAll(namespace string) error {
	err := s.Client.DeleteAllOf(context.Background(), &scriptsAPI.Script{ObjectMeta: metav1.ObjectMeta{Namespace: namespace}})
	return err
}
