package config

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kubeshop/testkube-operator/pkg/configmap"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewConfigMapConfig is a constructor for configmap config
func NewConfigMapConfig(cli client.Client, name, namespace string) (*ConfigMapConfig, error) {
	client, err := configmap.NewClient(cli, namespace)
	if err != nil {
		return nil, err
	}

	return &ConfigMapConfig{
		name:   name,
		client: client,
	}, nil
}

// ConfigMapConfig contains configmap config properties
type ConfigMapConfig struct {
	name   string
	client *configmap.Client
}

// GetUniqueClusterId gets unique cluster based ID
func (c *ConfigMapConfig) GetUniqueClusterId(ctx context.Context) (clusterId string, err error) {
	data, err := c.client.Get(c.name)
	if err != nil {
		return "", fmt.Errorf("reading config map error: %w", err)
	}

	return data["clusterId"], nil
}

// GetTelemetryEnabled get telemetry enabled
func (c *ConfigMapConfig) GetTelemetryEnabled(ctx context.Context) (ok bool, err error) {
	data, err := c.client.Get(c.name)
	if err != nil {
		return true, fmt.Errorf("reading config map error: %w", err)
	}

	var result bool
	if enableTelemetry, ok := data["enableTelemetry"]; ok {
		result, err = strconv.ParseBool(enableTelemetry)
		if err != nil {
			return true, fmt.Errorf("parsing enable telemetry error: %w", err)
		}
	}

	return result, nil
}
