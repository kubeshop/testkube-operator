package tests

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testsv2 "github.com/kubeshop/testkube-operator/apis/tests/v2"
	"k8s.io/apimachinery/pkg/labels"
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
func (s TestsClient) List(namespace string, selector string) (*testsv2.TestList, error) {
	list := &testsv2.TestList{}
	reqs, err := labels.ParseToRequirements(selector)
	if err != nil {
		return list, err
	}

	options := &client.ListOptions{
		Namespace:     namespace,
		LabelSelector: labels.NewSelector().Add(reqs...),
	}
	if err = s.Client.List(context.Background(), list, options); err != nil {
		return list, err
	}

	return list, nil
}

// ListLabels labels for Tests
func (s TestsClient) ListLabels(namespace string) (map[string][]string, error) {
	labels := map[string][]string{}
	list := &testsv2.TestList{}
	if err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace}); err != nil {
		return labels, err
	}

	for _, test := range list.Items {
		for key, value := range test.Labels {
			if values, ok := labels[key]; !ok {
				labels[key] = []string{value}
			} else {
				for _, v := range values {
					if v == value {
						continue
					}
				}
				labels[key] = append(labels[key], value)
			}
		}
	}

	return labels, nil
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
