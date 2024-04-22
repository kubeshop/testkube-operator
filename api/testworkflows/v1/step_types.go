package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	testsv3 "github.com/kubeshop/testkube-operator/api/tests/v3"
)

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

	// configuration for pausing the step initially
	Pause *PauseConfig `json:"pause,omitempty" expr:"include"`

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

	// steps to run before other operations in this step
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Setup []IndependentStep `json:"setup,omitempty" expr:"include"`

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

	// steps to run before other operations in this step
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Setup []Step `json:"setup,omitempty" expr:"include"`

	// sub-steps to run
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Steps []Step `json:"steps,omitempty" expr:"include"`
}

type StepRun struct {
	ContainerConfig `json:",inline"`

	// script to run in a default shell for the container
	Shell *string `json:"shell,omitempty" expr:"template"`
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

type TarballRequest struct {
	// path to load the files from
	From string `json:"from,omitempty" expr:"template"`

	// file patterns to pack
	Files *DynamicList `json:"files,omitempty" expr:"template"`
}

type StepExecuteStrategy struct {
	// matrix of parameters to spawn instances (static)
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type="object"
	Matrix map[string]DynamicList `json:"matrix,omitempty" expr:"force"`

	// static number of sharded instances to spawn
	Count *intstr.IntOrString `json:"count,omitempty" expr:"expression"`

	// dynamic number of sharded instances to spawn - it will be lowered if there is not enough sharded values
	MaxCount *intstr.IntOrString `json:"maxCount,omitempty" expr:"expression"`

	// parameters that should be distributed across sharded instances
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type="object"
	Shards map[string]DynamicList `json:"shards,omitempty" expr:"force"`
}

type StepExecuteTest struct {
	// test name to run
	Name string `json:"name,omitempty" expr:"template"`

	// test execution description to display
	Description string `json:"description,omitempty" expr:"template"`

	StepExecuteStrategy `json:",inline" expr:"include"`

	// pack some data from the original file system to serve them down
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type="object"
	Tarball map[string]TarballRequest `json:"tarball,omitempty" expr:"template,include"`

	// pass the execution request overrides
	ExecutionRequest *TestExecutionRequest `json:"executionRequest,omitempty" expr:"include"`
}

type StepExecuteWorkflow struct {
	// workflow name to run
	Name string `json:"name,omitempty" expr:"template"`

	// test workflow execution description to display
	Description string `json:"description,omitempty" expr:"template"`

	StepExecuteStrategy `json:",inline" expr:"include"`

	// unique execution name to use
	ExecutionName string `json:"executionName,omitempty" expr:"template"`

	// pack some data from the original file system to serve them down
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type="object"
	Tarball map[string]TarballRequest `json:"tarball,omitempty" expr:"template,include"`

	// configuration to pass for the workflow
	Config map[string]intstr.IntOrString `json:"config,omitempty" expr:"template"`
}

type StepArtifacts struct {
	// working directory to override, so it will be used as a base dir
	WorkingDir *string `json:"workingDir,omitempty" expr:"template"`
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

type TestExecutionRequest struct {
	// test execution custom name
	Name string `json:"name,omitempty" expr:"template"`
	// test execution labels
	ExecutionLabels map[string]string `json:"executionLabels,omitempty" expr:"template,template"`
	// variables file content - need to be in format for particular executor (e.g. postman envs file)
	VariablesFile           string                      `json:"variablesFile,omitempty" expr:"template"`
	IsVariablesFileUploaded bool                        `json:"isVariablesFileUploaded,omitempty" expr:"ignore"`
	Variables               map[string]testsv3.Variable `json:"variables,omitempty" expr:"template,force"`
	// test secret uuid
	TestSecretUUID string `json:"testSecretUUID,omitempty" expr:"template"`
	// additional executor binary arguments
	Args []string `json:"args,omitempty" expr:"template"`
	// usage mode for arguments
	ArgsMode testsv3.ArgsModeType `json:"argsMode,omitempty" expr:"template"`
	// executor binary command
	Command []string `json:"command,omitempty" expr:"template"`
	// container executor image
	Image string `json:"image,omitempty" expr:"template"`
	// container executor image pull secrets
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" expr:"template"`
	// whether to start execution sync or async
	Sync bool `json:"sync,omitempty" expr:"ignore"`
	// http proxy for executor containers
	HttpProxy string `json:"httpProxy,omitempty" expr:"template"`
	// https proxy for executor containers
	HttpsProxy string `json:"httpsProxy,omitempty" expr:"template"`
	// negative test will fail the execution if it is a success and it will succeed if it is a failure
	NegativeTest bool `json:"negativeTest,omitempty" expr:"ignore"`
	// Optional duration in seconds the pod may be active on the node relative to
	// StartTime before the system will actively try to mark it failed and kill associated containers.
	// Value must be a positive integer.
	ActiveDeadlineSeconds int64                    `json:"activeDeadlineSeconds,omitempty" expr:"ignore"`
	ArtifactRequest       *testsv3.ArtifactRequest `json:"artifactRequest,omitempty" expr:"force"`
	// job template extensions
	JobTemplate string `json:"jobTemplate,omitempty" expr:"ignore"`
	// cron job template extensions
	CronJobTemplate string `json:"cronJobTemplate,omitempty" expr:"ignore"`
	// script to run before test execution
	PreRunScript string `json:"preRunScript,omitempty" expr:"template"`
	// script to run after test execution
	PostRunScript string `json:"postRunScript,omitempty" expr:"template"`
	// execute post run script before scraping (prebuilt executor only)
	ExecutePostRunScriptBeforeScraping bool `json:"executePostRunScriptBeforeScraping,omitempty" expr:"ignore"`
	// run scripts using source command (container executor only)
	SourceScripts bool `json:"sourceScripts,omitempty" expr:"ignore"`
	// scraper template extensions
	ScraperTemplate string `json:"scraperTemplate,omitempty" expr:"ignore"`
	// config map references
	EnvConfigMaps []testsv3.EnvReference `json:"envConfigMaps,omitempty" expr:"force"`
	// secret references
	EnvSecrets []testsv3.EnvReference `json:"envSecrets,omitempty" expr:"force"`
	// namespace for test execution (Pro edition only)
	ExecutionNamespace string `json:"executionNamespace,omitempty" expr:"template"`
}
