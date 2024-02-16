package v1

import "k8s.io/apimachinery/pkg/util/intstr"

type RetryPolicy struct {
	// how many times at most it should retry
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	Count int32 `json:"count,omitempty"`

	// until when it should retry (defaults to: "passed")
	Until Expression `json:"until,omitempty"`
}

type StepBase struct {
	// readable name for the step
	Name string `json:"name,omitempty"`

	// expression to declare under which conditions the step should be run
	// defaults to: "passed", except artifacts where it defaults to "always"
	Condition Expression `json:"condition,omitempty"`

	// is the step expected to fail
	Negative bool `json:"negative,omitempty"`

	// is the step optional, so its failure won't affect the TestWorkflow result
	Optional bool `json:"optional,omitempty"`

	// should not display it as a nested group
	VirtualGroup bool `json:"virtualGroup,omitempty"`

	// policy for retrying the step
	Retry *RetryPolicy `json:"retry,omitempty"`

	// maximum time this step may take
	// +kubebuilder:validation:Pattern=^((0|[1-9][0-9]*)h)?((0|[1-9][0-9]*)m)?((0|[1-9][0-9]*)s)?((0|[1-9][0-9]*)ms)?$
	Timeout string `json:"timeout,omitempty"`

	// delay before the step
	// +kubebuilder:validation:Pattern=^((0|[1-9][0-9]*)h)?((0|[1-9][0-9]*)m)?((0|[1-9][0-9]*)s)?((0|[1-9][0-9]*)ms)?$
	Delay string `json:"delay,omitempty"`

	// working directory to use for this step
	WorkingDir *string `json:"workingDir,omitempty"`

	// defaults for the containers in this step
	Container *ContainerConfig `json:"container,omitempty"`

	// content that should be fetched for this step
	Content *Content `json:"content,omitempty"`

	// script to run in a default shell for the container
	Shell string `json:"shell,omitempty"`

	// run specific container in the current step
	Run *StepRun `json:"run,omitempty"`

	// execute other Testkube resources
	Execute *StepExecute `json:"execute,omitempty"`

	// scrape artifacts from the volumes
	Artifacts *StepArtifacts `json:"artifacts,omitempty"`
}

type IndependentStep struct {
	StepBase `json:",inline"`

	// sub-steps to run
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Steps []IndependentStep `json:"steps,omitempty"`
}

type Step struct {
	StepBase `json:",inline"`

	// multiple templates to include in this step
	Use []TemplateRef `json:"use,omitempty"`

	// single template to run in this step
	Template TemplateRef `json:"template,omitempty"`

	// sub-steps to run
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Steps []Step `json:"steps,omitempty"`
}

type StepRun struct {
	Container ContainerConfig `json:",inline"`
}

type StepExecute struct {
	// how many resources could be scheduled in parallel
	Parallelism int32 `json:"parallelism,omitempty"`

	// only schedule the resources, don't watch the results (unless it is needed for parallelism)
	Async bool `json:"async,omitempty"`

	// tests to run
	Tests []StepExecuteTest `json:"tests,omitempty"`

	// workflows to run
	Workflows []StepExecuteWorkflow `json:"workflows,omitempty"`
}

type StepExecuteTest struct {
	// test name to run
	Name string `json:"name,omitempty"`
}

type StepExecuteWorkflow struct {
	// workflow name to run
	Name string `json:"name,omitempty"`
	// configuration to pass for the workflow
	Config *map[string]intstr.IntOrString `json:"config,omitempty"`
}

type StepArtifacts struct {
	// paths to fetch from the container
	Paths []string `json:"paths,omitempty"`
	// compression options for the artifacts
	Compress *ArtifactCompression `json:"compress,omitempty"`
}

type ArtifactCompression struct {
	// artifact name
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
}
