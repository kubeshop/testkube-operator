package v3

import (
	"context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"

	workflowsv1beta1 "github.com/kubeshop/testkube-operator/api/workflows/v1beta1"
)

//go:generate mockgen -destination=./mock_workflows.go -package=v3 "github.com/kubeshop/testkube-operator/pkg/client/workflows/v1beta1" Interface
type Interface interface {
	List(selector string) (*workflowsv1beta1.WorkflowList, error)
	ListLabels() (map[string][]string, error)
	Get(name string) (*workflowsv1beta1.Workflow, error)
	Create(workflow *workflowsv1beta1.Workflow) (*workflowsv1beta1.Workflow, error)
	Update(workflow *workflowsv1beta1.Workflow) (*workflowsv1beta1.Workflow, error)
	Delete(name string) error
	DeleteAll() error
	GetCurrentSecretUUID(workflowName string) (string, error)
	DeleteByLabels(selector string) error
	UpdateStatus(workflow *workflowsv1beta1.Workflow) error
}

// NewClient creates new Workflow client
func NewClient(client client.Client, namespace string) *WorkflowsClient {
	return &WorkflowsClient{
		Client:    client,
		Namespace: namespace,
	}
}

// WorkflowsClient implements methods to work with Workflows
type WorkflowsClient struct {
	Client    client.Client
	Namespace string
}

// List lists Workflows
func (s WorkflowsClient) List(selector string) (*workflowsv1beta1.WorkflowList, error) {
	list := &workflowsv1beta1.WorkflowList{}
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

// ListLabelslists labels for Workflows
func (s WorkflowsClient) ListLabels() (map[string][]string, error) {
	labels := map[string][]string{}
	list := &workflowsv1beta1.WorkflowList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: s.Namespace})
	if err != nil {
		return labels, err
	}

	for _, workflow := range list.Items {
		for key, value := range workflow.Labels {
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

// Get returns Workflow
func (s WorkflowsClient) Get(name string) (*workflowsv1beta1.Workflow, error) {
	workflow := &workflowsv1beta1.Workflow{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: name}, workflow)
	if err != nil {
		return nil, err
	}
	return workflow, nil
}

// Create creates new Workflow
func (s WorkflowsClient) Create(workflow *workflowsv1beta1.Workflow) (*workflowsv1beta1.Workflow, error) {
	return workflow, s.Client.Create(context.Background(), workflow)
}

// Update updates existing Workflow
func (s WorkflowsClient) Update(workflow *workflowsv1beta1.Workflow) (*workflowsv1beta1.Workflow, error) {
	return workflow, s.Client.Update(context.Background(), workflow)
}

// Delete deletes existing Workflow
func (s WorkflowsClient) Delete(name string) error {
	workflow, err := s.Get(name)
	if err != nil {
		return err
	}
	return s.Client.Delete(context.Background(), workflow)
}

// DeleteAll delete all Workflows
func (s WorkflowsClient) DeleteAll() error {
	u := &unstructured.Unstructured{}
	u.SetKind("Workflow")
	u.SetAPIVersion("workflows.testkube.io/v1beta1")
	return s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace))
}

// UpdateStatus updates existing Workflow status
func (s WorkflowsClient) UpdateStatus(workflow *workflowsv1beta1.Workflow) error {
	return s.Client.Status().Update(context.Background(), workflow)
}

// DeleteByLabels deletes Workflows by labels
func (s WorkflowsClient) DeleteByLabels(selector string) error {
	reqs, err := labels.ParseToRequirements(selector)
	if err != nil {
		return err
	}

	u := &unstructured.Unstructured{}
	u.SetKind("Workflow")
	u.SetAPIVersion("workflows.testkube.io/v1beta1")
	err = s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace),
		client.MatchingLabelsSelector{Selector: labels.NewSelector().Add(reqs...)})
	return err
}
