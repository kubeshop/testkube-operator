package v1

import corev1 "k8s.io/api/core/v1"

type ServiceRestartPolicy string

const (
	ServiceRestartPolicyOnFailure ServiceRestartPolicy = "OnFailure"
	ServiceRestartPolicyNever     ServiceRestartPolicy = "Never"
)

type IndependentServiceSpec struct {
	StepExecuteStrategy `json:",inline" expr:"include"`

	// service description to display
	Description string `json:"description,omitempty" expr:"template"`

	// should save logs for the service (false if not specified)
	Logs *string `json:"logs,omitempty" expr:"expression"`

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

type ServiceSpec struct {
	// multiple templates to include in this step
	Use []TemplateRef `json:"use,omitempty" expr:"include"`

	IndependentServiceSpec `json:",inline" expr:"include"`
}
