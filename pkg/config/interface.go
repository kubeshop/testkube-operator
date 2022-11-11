package config

import (
	"context"
)

type Repository interface {
	// GetUniqueClusterId gets unique cluster based ID
	GetUniqueClusterId(ctx context.Context) (string, error)

	// GetTelemetryEnabled get telemetry enabled
	GetTelemetryEnabled(ctx context.Context) (ok bool, err error)
}
