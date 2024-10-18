package v1

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testtriggersv1 "github.com/kubeshop/testkube-operator/api/testtriggers/v1"
)

//go:generate mockgen -destination=./mock_testtriggers.go -package=v1 "github.com/kubeshop/testkube-operator/pkg/client/testtriggers/v1" Interface
type Interface interface {
	List(selector string) (*testtriggersv1.TestTriggerList, error)
	Get(name string) (*testtriggersv1.TestTrigger, error)
	Create(trigger *testtriggersv1.TestTrigger) (*testtriggersv1.TestTrigger, error)
	Update(trigger *testtriggersv1.TestTrigger) (*testtriggersv1.TestTrigger, error)
	Delete(name string) error
	DeleteByLabels(selector string) error
}

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
func (s TestTriggersClient) List(selector string) (*testtriggersv1.TestTriggerList, error) {
	list := &testtriggersv1.TestTriggerList{}
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

// Get returns TestTrigger
func (s TestTriggersClient) Get(name string) (*testtriggersv1.TestTrigger, error) {
	trigger := &testtriggersv1.TestTrigger{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: name}, trigger)
	if err != nil {
		return nil, err
	}
	return trigger, nil
}

// Create creates new TestTrigger
func (s TestTriggersClient) Create(trigger *testtriggersv1.TestTrigger) (*testtriggersv1.TestTrigger, error) {
	return trigger, s.Client.Create(context.Background(), trigger)
}

// Update updates existing TestTrigger
func (s TestTriggersClient) Update(trigger *testtriggersv1.TestTrigger) (*testtriggersv1.TestTrigger, error) {
	return trigger, s.Client.Update(context.Background(), trigger)
}

// Delete deletes existing TestTrigger
func (s TestTriggersClient) Delete(name string) error {
	trigger, err := s.Get(name)
	if err != nil {
		return err
	}
	return s.Client.Delete(context.Background(), trigger)
}

// DeleteByLabels deletes TestTriggers by labels
func (s TestTriggersClient) DeleteByLabels(selector string) error {
	reqs, err := labels.ParseToRequirements(selector)
	if err != nil {
		return err
	}

	u := &unstructured.Unstructured{}
	u.SetKind("TestTrigger")
	u.SetAPIVersion(testtriggersv1.GroupVersion.String())
	err = s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace),
		client.MatchingLabelsSelector{Selector: labels.NewSelector().Add(reqs...)})
	return err
}
