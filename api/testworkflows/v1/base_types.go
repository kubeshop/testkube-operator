package v1

type TestWorkflowSpecBase struct {
	// Important: Run "make" to regenerate code after modifying this file

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

	// events triggering execution of the test workflow
	Events []Event `json:"events,omitempty" expr:"include"`
}
