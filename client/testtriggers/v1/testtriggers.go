package v1

import (
	"context"
	testtriggerv1 "github.com/kubeshop/testkube-operator/apis/testtriggers/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewClient creates new TestTrigger client
func NewClient(client client.Client, namespace string) *TestTriggersClient {
	return &TestTriggersClient{
		Client:    client,
		Namespace: namespace,
	}
}

// TestTriggersClient implements methods to work with TestTriggers
type TestTriggersClient struct {
	Client    client.Client
	Namespace string
}

// List lists TestTriggers
func (s *TestTriggersClient) List(ctx context.Context, selector string) (*testtriggerv1.TestTriggerList, error) {
	list := &testtriggerv1.TestTriggerList{}
	reqs, err := labels.ParseToRequirements(selector)
	if err != nil {
		return list, err
	}

	options := &client.ListOptions{
		Namespace:     s.Namespace,
		LabelSelector: labels.NewSelector().Add(reqs...),
	}

	if err = s.Client.List(ctx, list, options); err != nil {
		return list, err
	}

	return list, nil
}

// Get returns TestTrigger
func (s *TestTriggersClient) Get(ctx context.Context, name string) (*testtriggerv1.TestTrigger, error) {
	testtrigger := &testtriggerv1.TestTrigger{}
	if err := s.Client.Get(ctx, client.ObjectKey{Namespace: s.Namespace, Name: name}, testtrigger); err != nil {
		return nil, err
	}

	return testtrigger, nil
}

// Create creates new TestTrigger
func (s *TestTriggersClient) Create(ctx context.Context, t *testtriggerv1.TestTrigger) (*testtriggerv1.TestTrigger, error) {
	err := s.Client.Create(ctx, t)
	return t, err
}

// Update updates existing TestTrigger
func (s *TestTriggersClient) Update(ctx context.Context, t *testtriggerv1.TestTrigger) (*testtriggerv1.TestTrigger, error) {
	err := s.Client.Update(ctx, t)
	return t, err
}

// Delete deletes existing TestTrigger
func (s *TestTriggersClient) Delete(ctx context.Context, name string) error {
	testtrigger, err := s.Get(ctx, name)
	if err != nil {
		return err
	}

	if err := s.Client.Delete(ctx, testtrigger); err != nil {
		return err
	}

	return nil
}

// DeleteAll delete all TestTriggers
func (s *TestTriggersClient) DeleteAll(ctx context.Context) error {
	u := &unstructured.Unstructured{}
	u.SetKind("TestTrigger")
	u.SetAPIVersion("tests.testkube.io/v1")
	if err := s.Client.DeleteAllOf(ctx, u, client.InNamespace(s.Namespace)); err != nil {
		return err
	}

	return nil
}
