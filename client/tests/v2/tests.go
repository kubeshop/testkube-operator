package tests

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testsv2 "github.com/kubeshop/testkube-operator/apis/tests/v2"
	"github.com/kubeshop/testkube-operator/utils"
)

// NewClent creates new Test client
func NewClient(client client.Client) *TestsClient {
	return &TestsClient{
		Client: client,
	}
}

// TestClient implements methods to work with Test
type TestsClient struct {
	Client client.Client
}

// List lists Tests
func (s TestsClient) List(namespace string, tags []string) (*testsv2.TestList, error) {
	list := &testsv2.TestList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	if len(tags) == 0 {
		return list, err
	}

	toReturn := &testsv2.TestList{}
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

// ListTags tags for Tests
func (s TestsClient) ListTags(namespace string) ([]string, error) {
	tags := []string{}
	list := &testsv2.TestList{}
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

// Get returns Test
func (s TestsClient) Get(namespace, name string) (*testsv2.Test, error) {
	test := &testsv2.Test{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: name}, test)
	return test, err
}

// Create creates new Test
func (s TestsClient) Create(test *testsv2.Test) (*testsv2.Test, error) {
	err := s.Client.Create(context.Background(), test)
	return test, err
}

// Update updates existing Test
func (s TestsClient) Update(test *testsv2.Test) (*testsv2.Test, error) {
	err := s.Client.Update(context.Background(), test)
	return test, err
}

// Delete deletes existing Test
func (s TestsClient) Delete(namespace, name string) error {
	test, err := s.Get(namespace, name)
	if err != nil {
		return err
	}

	err = s.Client.Delete(context.Background(), test)
	return err
}

// DeleteAll deletes all Tests
func (s TestsClient) DeleteAll(namespace string) error {

	u := &unstructured.Unstructured{}
	u.SetKind("Test")
	u.SetAPIVersion("tests.testkube.io/v2")
	err := s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(namespace))
	return err
}
