package v1

type TestWorkflowSpecBase struct {
	// Important: Run "make" to regenerate code after modifying this file

	// events triggering execution of the test workflow
	Events []Event `json:"events,omitempty" expr:"include"`

	// system configuration to define the orchestration behavior
	System *TestWorkflowSystem `json:"system,omitempty" expr:"include"`

	// make the instance configurable with some input data for scheduling it
	Config map[string]ParameterSchema `json:"config,omitempty" expr:"include"`

	// global content that should be fetched into all containers
	Content *Content `json:"content,omitempty" expr:"include"`

	// defaults for the containers for all the TestWorkflow steps
	Container *ContainerConfig `json:"container,omitempty" expr:"include"`

	// configuration for the scheduled job
	Job *JobConfig `json:"job,omitempty" expr:"include"`

	// configuration for the scheduled pod
	Pod *PodConfig `json:"pod,omitempty" expr:"include"`

	// configuration for notifications
	// Deprecated: field is not used
	Notifications *NotificationsConfig `json:"notifications,omitempty" expr:"include"`

	// values to be used for test workflow execution
	Execution *TestWorkflowTagSchema `json:"execution,omitempty" expr:"include"`

	// Specifies how to treat concurrent executions of a Workflow.
	// Valid values are:
	//
	// - "Allow" (default): allows Workflows to execute concurrently;
	// - "Forbid": forbids concurrent executions, rejecting next execution if previous execution hasn't finished yet;
	// - "Replace": cancels currently running Workflow and replaces it with a new one
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`
}

// ConcurrencyPolicy describes how the workflow will be handled.
// Only one of the following concurrent policies may be specified.
// If none of the following policies is specified, the default one
// is AllowConcurrent.
// +enum
type ConcurrencyPolicy string

const (
	// AllowConcurrent allows Workflows to execute concurrently.
	AllowConcurrent ConcurrencyPolicy = "Allow"

	// ForbidConcurrent disables concurrent executions, rejecting next execution if previous
	// hasn't finished yet.
	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	// ReplaceConcurrent cancels currently running Workflow and replaces it with a new one.
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

type TestWorkflowSystem struct {
	// assume all the steps are pure by default
	PureByDefault *bool `json:"pureByDefault,omitempty"`

	// disable the behavior of merging multiple operations in a single container
	IsolatedContainers *bool `json:"isolatedContainers,omitempty"`
}
