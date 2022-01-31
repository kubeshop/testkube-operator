package scripts

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	scriptsv2 "github.com/kubeshop/testkube-operator/apis/script/v2"
	"github.com/kubeshop/testkube-operator/utils"
)

func NewClient(client client.Client) *ScriptsClient {
	return &ScriptsClient{
		Client: client,
	}
}

type ScriptsClient struct {
	Client client.Client
}

func (s ScriptsClient) List(namespace string, tags []string) (*scriptsv2.ScriptList, error) {
	list := &scriptsv2.ScriptList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	if len(tags) == 0 {
		return list, err
	}

	toReturn := &scriptsv2.ScriptList{}
	for _, script := range list.Items {
		hasTags := false
		for _, tag := range tags {
			if utils.ContainsTag(script.Spec.Tags, tag) {
				hasTags = true
			} else {
				hasTags = false
			}
		}
		if hasTags {
			toReturn.Items = append(toReturn.Items, script)

		}
	}

	return toReturn, nil
}

func (s ScriptsClient) ListTags(namespace string) ([]string, error) {
	tags := []string{}
	list := &scriptsv2.ScriptList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	if err != nil {
		return tags, err
	}

	for _, test := range list.Items {
		tags = append(tags, test.Spec.Tags...)
	}

	tags = utils.RemoveDuplicates(tags)

	return tags, nil
}

func (s ScriptsClient) Get(namespace, name string) (*scriptsv2.Script, error) {
	script := &scriptsv2.Script{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: name}, script)
	return script, err
}

func (s ScriptsClient) Create(scripts *scriptsv2.Script) (*scriptsv2.Script, error) {
	err := s.Client.Create(context.Background(), scripts)
	return scripts, err
}

func (s ScriptsClient) Update(scripts *scriptsv2.Script) (*scriptsv2.Script, error) {
	err := s.Client.Update(context.Background(), scripts)
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

	u := &unstructured.Unstructured{}
	u.SetKind("script")
	u.SetAPIVersion("tests.testkube.io/v2")
	err := s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(namespace))
	return err
}
