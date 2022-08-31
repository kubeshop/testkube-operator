package v1

import (
	corev1 "k8s.io/api/core/v1"
)

type Header struct {
	// header key name
	Name string `json:"name,omitempty"`
	// header string value
	Value string `json:"value,omitempty"`
	// or load it from var source
	ValueFrom corev1.EnvVarSource `json:"valueFrom,omitempty"`
}
