package configmap

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Client provide methods to manage configmaps
type Client struct {
	client.Client
	namespace string
}

// NewClient is a method to create new configmap client
func NewClient(cli client.Client, namespace string) (*Client, error) {
	return &Client{
		Client:    cli,
		namespace: namespace,
	}, nil
}

// Get is a method to retrieve an existing configmap
func (c *Client) Get(id string) (map[string]string, error) {
	configMap := &corev1.ConfigMap{}
	ctx := context.Background()

	if err := c.Client.Get(ctx, client.ObjectKey{
		Namespace: c.namespace, Name: id}, configMap); err != nil {
		return nil, err
	}

	stringData := map[string]string{}
	for key, value := range configMap.Data {
		stringData[key] = value
	}

	return stringData, nil
}
