package tests

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testsAPI "github.com/kubeshop/testkube-operator/apis/tests/v1"
	"github.com/kubeshop/testkube-operator/utils"
)

func NewClient(client client.Client) *TestsClient {
	return &TestsClient{
		Client: client,
	}
}

type TestsClient struct {
	Client client.Client
}

func (s TestsClient) List(namespace string, tags []string) (*testsAPI.TestList, error) {
	list := &testsAPI.TestList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	if len(tags) == 0 {
		return list, err
	}

	toReturn := &testsAPI.TestList{}
	for _, test := range list.Items {
		hasTags := false
		for _, tag := range tags {
			if utils.ContainsTag(test.Spec.Tags, tag) {
				hasTags = true
			} else {
				hasTags = false
			}

		}
		if hasTags {
			toReturn.Items = append(toReturn.Items, test)

		}
	}
	return toReturn, nil
}

func (s TestsClient) ListTags(namespace string) ([]string, error) {
	tags := []string{}
	list := &testsAPI.TestList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	if err != nil {
		return tags, err
	}

	for _, test := range list.Items {
		tags = append(tags, test.Spec.Tags...)
	}

	return tags, nil
}

func (s TestsClient) Get(namespace, name string) (*testsAPI.Test, error) {
	Test := &testsAPI.Test{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: name}, Test)
	return Test, err
}

func (s TestsClient) Create(Tests *testsAPI.Test) (*testsAPI.Test, error) {
	err := s.Client.Create(context.Background(), Tests)
	return Tests, err
}

func (s TestsClient) Update(Tests *testsAPI.Test) (*testsAPI.Test, error) {
	err := s.Client.Update(context.Background(), Tests)
	return Tests, err
}

func (s TestsClient) Delete(namespace, name string) error {
	Test, err := s.Get(namespace, name)
	if err != nil {
		return err
	}

	err = s.Client.Delete(context.Background(), Test)
	return err
}

func (s TestsClient) DeleteAll(namespace string) error {
	u := &unstructured.Unstructured{}
	u.SetKind("Test")
	u.SetAPIVersion("tests.testkube.io/v1")
	err := s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(namespace))
	return err
}
