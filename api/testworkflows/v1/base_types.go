package v1

type TestWorkflowSpecBase struct {
	// Important: Run "make" to regenerate code after modifying this file

	// make the instance configurable with some input data for scheduling it
	Config *map[string]ParameterSchema `json:"config,omitempty"`

	// global content that should be fetched into all containers
	Content *Content `json:"content,omitempty"`

	// defaults for the containers for all the TestWorkflow steps
	Container *ContainerConfig `json:"container,omitempty"`

	// configuration for the scheduled job
	Job *JobConfig `json:"job,omitempty"`

	// configuration for the scheduled pod
	Pod *PodConfig `json:"pod,omitempty"`
}
