package v1

import corev1 "k8s.io/api/core/v1"

type ServiceRestartPolicy string

const (
	ServiceRestartPolicyOnFailure ServiceRestartPolicy = "OnFailure"
	ServiceRestartPolicyNever     ServiceRestartPolicy = "Never"
)

type ServiceSpec struct {
	StepExecuteStrategy `json:",inline" expr:"include"`

	// maximum time until reaching readiness
	// +kubebuilder:validation:Pattern=^((0|[1-9][0-9]*)h)?((0|[1-9][0-9]*)m)?((0|[1-9][0-9]*)s)?((0|[1-9][0-9]*)ms)?$
	Timeout string `json:"timeout,omitempty"`

	// instructions for transferring files
	Transfer []StepParallelTransfer `json:"transfer,omitempty" expr:"include"`

	// global content that should be fetched into all containers
	Content *Content `json:"content,omitempty" expr:"include"`

	// configuration for the scheduled pod
	Pod *PodConfig `json:"pod,omitempty" expr:"include"`

	StepRun `json:",inline" expr:"include"`

	// Restart policy for the main container in the pod. One of OnFailure or Never.
	RestartPolicy ServiceRestartPolicy `json:"restartPolicy,omitempty" expr:"template"`

	// Probe to check if the service has started correctly
	ReadinessProbe *corev1.Probe `json:"readinessProbe,omitempty" expr:"force"`
}
