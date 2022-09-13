package testsources

import (
	"context"

	testsourcev1 "github.com/kubeshop/testkube-operator/apis/testsource/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewClient returns new client instance, needs kubernetes client to be passed as dependecy
func NewClient(client client.Client, namespace string) *TestSourcesClient {
	return &TestSourcesClient{
		Client:    client,
		Namespace: namespace,
	}
}

// TestSourcesClient client for getting test sources CRs
type TestSourcesClient struct {
	Client    client.Client
	Namespace string
}

// List shows list of available test sources
func (s TestSourcesClient) List(selector string) (*testsourcev1.TestSourceList, error) {
	list := &testsourcev1.TestSourceList{}
	reqs, err := labels.ParseToRequirements(selector)
	if err != nil {
		return list, err
	}

	options := &client.ListOptions{
		Namespace:     s.Namespace,
		LabelSelector: labels.NewSelector().Add(reqs...),
	}

	err = s.Client.List(context.Background(), list, options)
	return list, err
}

// Get gets testsource by name in given namespace
func (s TestSourcesClient) Get(name string) (*testsourcev1.TestSource, error) {
	testsource := &testsourcev1.TestSource{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: name}, testsource)
	return testsource, err
}

// Create creates new TestSource CR
func (s TestSourcesClient) Create(testsource *testsourcev1.TestSource) (*testsourcev1.TestSource, error) {
	err := s.Client.Create(context.Background(), testsource)
	return testsource, err
}

// Delete deletes testsource by name
func (s TestSourcesClient) Delete(name string) error {
	testsource := &testsourcev1.TestSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: s.Namespace,
		},
	}
	err := s.Client.Delete(context.Background(), testsource)
	return err
}

// Update updates test source
func (s TestSourcesClient) Update(testsource *testsourcev1.TestSource) (*testsourcev1.TestSource, error) {
	err := s.Client.Update(context.Background(), testsource)
	return testsource, err
}

// DeleteByLabels deletes test sources by labels
func (s TestSourcesClient) DeleteByLabels(selector string) error {
	reqs, err := labels.ParseToRequirements(selector)
	if err != nil {
		return err
	}

	u := &unstructured.Unstructured{}
	u.SetKind("TestSource")
	u.SetAPIVersion("tests.testkube.io/v1")
	err = s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace),
		client.MatchingLabelsSelector{Selector: labels.NewSelector().Add(reqs...)})
	return err
}
