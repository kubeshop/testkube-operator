package v1

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testsuitev1 "github.com/kubeshop/testkube-operator/apis/testsuite/v1"
	"github.com/kubeshop/testkube-operator/utils"
)

// NewClient creates new TestSuite client
func NewClient(client client.Client) *TestSuitesClient {
	return &TestSuitesClient{
		Client: client,
	}
}

// TestSuitesClient implements methods to work with TestSuites
type TestSuitesClient struct {
	Client client.Client
}

// List lists TestSuites
func (s TestSuitesClient) List(namespace string, tags []string) (*testsuitev1.TestSuiteList, error) {
	list := &testsuitev1.TestSuiteList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	if len(tags) == 0 {
		return list, err
	}

	toReturn := &testsuitev1.TestSuiteList{}
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

// ListTags lists tags for TestSuites
func (s TestSuitesClient) ListTags(namespace string) ([]string, error) {
	tags := []string{}
	list := &testsuitev1.TestSuiteList{}
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

// Get returns TestSuite
func (s TestSuitesClient) Get(namespace, name string) (*testsuitev1.TestSuite, error) {
	test := &testsuitev1.TestSuite{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: name}, test)
	return test, err
}

// Create creates new TestSuite
func (s TestSuitesClient) Create(test *testsuitev1.TestSuite) (*testsuitev1.TestSuite, error) {
	err := s.Client.Create(context.Background(), test)
	return test, err
}

// Update updates existing TestSuite
func (s TestSuitesClient) Update(test *testsuitev1.TestSuite) (*testsuitev1.TestSuite, error) {
	err := s.Client.Update(context.Background(), test)
	return test, err
}

// Delete deletes existing TestSuite
func (s TestSuitesClient) Delete(namespace, name string) error {
	testSuite, err := s.Get(namespace, name)
	if err != nil {
		return err
	}

	err = s.Client.Delete(context.Background(), testSuite)
	return err
}

// DeleteAll delete all TestSuites
func (s TestSuitesClient) DeleteAll(namespace string) error {
	u := &unstructured.Unstructured{}
	u.SetKind("TestSuite")
	u.SetAPIVersion("tests.testkube.io/v1")
	err := s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(namespace))
	return err
}
