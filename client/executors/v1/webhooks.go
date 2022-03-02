package executors

import (
	"context"

	executorsv1 "github.com/kubeshop/testkube-operator/apis/executor/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewWebhooksClient returns new client instance, needs kubernetes client to be passed as dependecy
func NewWebhooksClient(client client.Client) *WebhooksClient {
	return &WebhooksClient{
		Client: client,
	}
}

// WebhooksClient client for getting webhooks CRs
type WebhooksClient struct {
	Client client.Client
}

// List shows list of available webhooks
func (s WebhooksClient) List(namespace string) (*executorsv1.WebhookList, error) {
	list := &executorsv1.WebhookList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: namespace})
	return list, err
}

// Get gets webhook by name in given namespace
func (s WebhooksClient) Get(namespace, name string) (*executorsv1.Webhook, error) {
	item := &executorsv1.Webhook{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: name}, item)
	return item, err
}

// GetByEvent gets all webhooks with given event
func (s WebhooksClient) GetByEvent(event string) (*executorsv1.WebhookList, error) {
	list := &executorsv1.WebhookList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{})
	if err != nil {
		return nil, err
	}

	for i, exec := range list.Items {
		hasEvent := false
		for _, t := range exec.Spec.Events {
			if t == event {
				hasEvent = true
			}
		}

		if !hasEvent {
			list.Items = append(list.Items[:i], list.Items[i+1:]...)
		}
	}

	return list, nil
}

// Create creates new Webhook CR
func (s WebhooksClient) Create(webhook *executorsv1.Webhook) (*executorsv1.Webhook, error) {
	err := s.Client.Create(context.Background(), webhook)
	return webhook, err
}

// Delete deletes Webhook by name
func (s WebhooksClient) Delete(name, namespace string) error {
	webhook := &executorsv1.Webhook{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	err := s.Client.Delete(context.Background(), webhook)
	return err
}

// Update updates Webhook
func (s WebhooksClient) Update(webhook *executorsv1.Webhook) (*executorsv1.Webhook, error) {
	err := s.Client.Update(context.Background(), webhook)
	return webhook, err
}
