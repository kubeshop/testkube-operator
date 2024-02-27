package v1

import "k8s.io/apimachinery/pkg/util/intstr"

type RetryPolicy struct {
	// how many times at most it should retry
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	Count int32 `json:"count,omitempty"`

	// until when it should retry (defaults to: "passed")
	Until string `json:"until,omitempty" expr:"expression"`
}

type StepBase struct {
	// readable name for the step
	Name string `json:"name,omitempty" expr:"template"`

	// expression to declare under which conditions the step should be run
	// defaults to: "passed", except artifacts where it defaults to "always"
	Condition string `json:"condition,omitempty" expr:"expression"`

	// is the step expected to fail
	Negative bool `json:"negative,omitempty"`

	// is the step optional, so its failure won't affect the TestWorkflow result
	Optional bool `json:"optional,omitempty"`

	// should not display it as a nested group
	VirtualGroup bool `json:"virtualGroup,omitempty"`

	// policy for retrying the step
	Retry *RetryPolicy `json:"retry,omitempty" expr:"include"`

	// maximum time this step may take
	// +kubebuilder:validation:Pattern=^((0|[1-9][0-9]*)h)?((0|[1-9][0-9]*)m)?((0|[1-9][0-9]*)s)?((0|[1-9][0-9]*)ms)?$
	Timeout string `json:"timeout,omitempty"`

	// delay before the step
	// +kubebuilder:validation:Pattern=^((0|[1-9][0-9]*)h)?((0|[1-9][0-9]*)m)?((0|[1-9][0-9]*)s)?((0|[1-9][0-9]*)ms)?$
	Delay string `json:"delay,omitempty"`

	// content that should be fetched for this step
	Content *Content `json:"content,omitempty" expr:"include"`

	// script to run in a default shell for the container
	Shell string `json:"shell,omitempty" expr:"template"`

	// run specific container in the current step
	Run *StepRun `json:"run,omitempty" expr:"include"`

	// working directory to use for this step
	WorkingDir *string `json:"workingDir,omitempty" expr:"template"`

	// defaults for the containers in this step
	Container *ContainerConfig `json:"container,omitempty" expr:"include"`

	// execute other Testkube resources
	Execute *StepExecute `json:"execute,omitempty" expr:"include"`

	// scrape artifacts from the volumes
	Artifacts *StepArtifacts `json:"artifacts,omitempty" expr:"include"`
}

type IndependentStep struct {
	StepBase `json:",inline" expr:"include"`

	// sub-steps to run
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Steps []IndependentStep `json:"steps,omitempty" expr:"include"`
}

type Step struct {
	StepBase `json:",inline" expr:"include"`

	// multiple templates to include in this step
	Use []TemplateRef `json:"use,omitempty" expr:"include"`

	// single template to run in this step
	Template *TemplateRef `json:"template,omitempty" expr:"include"`

	// sub-steps to run
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Steps []Step `json:"steps,omitempty" expr:"include"`
}

type StepRun struct {
	ContainerConfig `json:",inline"`
}

type StepExecute struct {
	// how many resources could be scheduled in parallel
	Parallelism int32 `json:"parallelism,omitempty"`

	// only schedule the resources, don't watch the results (unless it is needed for parallelism)
	Async bool `json:"async,omitempty"`

	// tests to run
	Tests []StepExecuteTest `json:"tests,omitempty" expr:"include"`

	// workflows to run
	Workflows []StepExecuteWorkflow `json:"workflows,omitempty" expr:"include"`
}

type StepExecuteTest struct {
	// test name to run
	Name string `json:"name,omitempty" expr:"template"`
}

type StepExecuteWorkflow struct {
	// workflow name to run
	Name string `json:"name,omitempty" expr:"template"`
	// configuration to pass for the workflow
	Config map[string]intstr.IntOrString `json:"config,omitempty" expr:"template"`
}

type StepArtifacts struct {
	// compression options for the artifacts
	Compress *ArtifactCompression `json:"compress,omitempty" expr:"include"`
	// paths to fetch from the container
	Paths []string `json:"paths,omitempty" expr:"template"`
}

type ArtifactCompression struct {
	// artifact name
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name" expr:"template"`
}
