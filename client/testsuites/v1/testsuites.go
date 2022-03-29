package v1

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testsuitev1 "github.com/kubeshop/testkube-operator/apis/testsuite/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// NewClient creates new TestSuite client
func NewClient(client client.Client, namespace string) *TestSuitesClient {
	return &TestSuitesClient{
		Client:    client,
		Namespace: namespace,
	}
}

// TestSuitesClient implements methods to work with TestSuites
type TestSuitesClient struct {
	Client    client.Client
	Namespace string
}

// List lists TestSuites
func (s TestSuitesClient) List(selector string) (*testsuitev1.TestSuiteList, error) {
	list := &testsuitev1.TestSuiteList{}
	reqs, err := labels.ParseToRequirements(selector)
	if err != nil {
		return list, err
	}

	options := &client.ListOptions{
		Namespace:     s.Namespace,
		LabelSelector: labels.NewSelector().Add(reqs...),
	}

	if err = s.Client.List(context.Background(), list, options); err != nil {
		return list, err
	}

	return list, nil
}

// ListLabelslists labels for TestSuites
func (s TestSuitesClient) ListLabels() (map[string][]string, error) {
	labels := map[string][]string{}
	list := &testsuitev1.TestSuiteList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: s.Namespace})
	if err != nil {
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

// Get returns TestSuite
func (s TestSuitesClient) Get(name string) (*testsuitev1.TestSuite, error) {
	test := &testsuitev1.TestSuite{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: name}, test)
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
func (s TestSuitesClient) Delete(name string) error {
	testSuite, err := s.Get(name)
	if err != nil {
		return err
	}

	err = s.Client.Delete(context.Background(), testSuite)
	return err
}

// DeleteAll delete all TestSuites
func (s TestSuitesClient) DeleteAll() error {
	u := &unstructured.Unstructured{}
	u.SetKind("TestSuite")
	u.SetAPIVersion("tests.testkube.io/v1")
	err := s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace))
	return err
}
