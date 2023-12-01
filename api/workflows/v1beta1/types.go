package v1beta1

import (
	"encoding/json"
	v3 "github.com/kubeshop/testkube-operator/api/tests/v3"
	corev1 "k8s.io/api/core/v1"
)

// +kubebuilder:validation:Enum=string;integer;number;boolean
type ParameterType string

const (
	ParameterTypeString  ParameterType = "string"
	ParameterTypeInteger ParameterType = "integer"
	ParameterTypeNumber  ParameterType = "number"
	ParameterTypeBoolean ParameterType = "boolean"
)

// +kubebuilder:validation:Enum=any;all
type RetrySuccess string

const (
	RetrySuccessAny RetrySuccess = "any"
	RetrySuccessAll RetrySuccess = "all"
)

// +kubebuilder:validation:Type=string
type Any json.RawMessage

// TODO: Consider `any` somehow for `const`, `default` and `enum`

// Parameter is similar to JSON Schema - includes `required` property directly like OpenAPI parameters
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:XPreserveUnknownFields
type Parameter struct {
	Description string `json:"description,omitempty"`
	// +kubebuilder:default=string
	Type ParameterType `json:"type,omitempty"`
	// +kubebuilder:default=false
	Required bool  `json:"required,omitempty"`
	Const    Any   `json:"const,omitempty"`
	Default  Any   `json:"default,omitempty"`
	Enum     []Any `json:"enum,omitempty"`

	// String modifiers

	MinLength *int64 `json:"minLength,omitempty"`
	MaxLength *int64 `json:"maxLength,omitempty"`
	Format    string `json:"format,omitempty"`
	Pattern   string `json:"pattern,omitempty"`

	// Custom string modifiers

	Image string `json:"image,omitempty"`

	// Number modifiers

	Minimum          *int64 `json:"minimum,omitempty"`
	Maximum          *int64 `json:"maximum,omitempty"`
	ExclusiveMinimum *int64 `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum *int64 `json:"exclusiveMaximum,omitempty"`
	MultipleOf       *int64 `json:"multipleOf,omitempty"`

	// Object modifiers
	// Don't support objects for now.

	// Array modifiers
	// Don't support lists for now.
}

type ContentGit struct {
	Uri               string         `json:"uri,omitempty"`
	Revision          string         `json:"revision,omitempty"`
	UsernameSecret    *v3.SecretRef  `json:"usernameSecret,omitempty"`
	TokenSecret       *v3.SecretRef  `json:"tokenSecret,omitempty"`
	CertificateSecret string         `json:"certificateSecret,omitempty"`
	AuthType          v3.GitAuthType `json:"authType,omitempty"`
	Paths             []string       `json:"paths,omitempty"`
}

type ContentFile struct {
	Path        string              `json:"path,omitempty"`
	Content     string              `json:"content,omitempty"`
	ContentFrom corev1.EnvVarSource `json:"contentFrom,omitempty"`
}

type Content struct {
	Git   *ContentGit   `json:"git,omitempty"`
	Files []ContentFile `json:"files,omitempty"`
}

type Resources struct {
	Limits   corev1.ResourceList `json:"limits,omitempty"`
	Requests corev1.ResourceList `json:"requests,omitempty"`
}

type Image struct {
	Name string `json:"name,omitempty"`
}

type ObjectRef struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name"`
}

type Variable struct {
	Value     string              `json:"value,omitempty"`
	ValueFrom corev1.EnvVarSource `json:"valueFrom,omitempty"`
}

type SpawnPod struct {
	ContainerConfig `json:",inline"`
	Image           string                 `json:"image"`
	Count           *int                   `json:"count,omitempty"`
	LivenessProbe   corev1.Probe           `json:"livenessProbe,omitempty"`
	ReadinessProbe  corev1.Probe           `json:"readinessProbe,omitempty"`
	StartupProbe    corev1.Probe           `json:"startupProbe,omitempty"`
	WorkingDir      string                 `json:"workingDir,omitempty"`
	Ports           []corev1.ContainerPort `json:"ports,omitempty"`
}

type StepSpawn *map[string]SpawnPod

type StepCache struct {
	Name     string   `json:"name"`
	RootPath string   `json:"rootPath,omitempty"` // defaults to workspace root
	Paths    []string `json:"paths,omitempty"`
}

type StepArtifacts struct {
	Name             string   `json:"name"`
	RootPath         string   `json:"rootPath,omitempty"` // defaults to workspace root
	Paths            []string `json:"paths,omitempty"`
	StorageClassName string   `json:"storageClassName,omitempty"`
	StorageBucket    string   `json:"storageBucket,omitempty"`
}

type StepReadCache struct {
	Name     string `json:"name"`
	RootPath string `json:"rootPath,omitempty"` // defaults to workspace root
}

type StepRun struct {
	ContainerConfig `json:",inline"`
	Negative        bool `json:"negative,omitempty"`
}

type StepJunit struct {
	RootPath string   `json:"rootPath,omitempty"` // defaults to workspace root
	Paths    []string `json:"paths"`
}

type StepExecuteWorkflow struct {
	Name string `json:"name,omitempty"`
}

type StepExecuteDuration struct {
	// +kubebuilder:validation:Pattern=^((0|[1-9][0-9]+)h)?((0|[1-9][0-9]+)m)?((0|[1-9][0-9]+)s)?((0|[1-9][0-9]+)ms)?$
	Minimum string `json:"minimum,omitempty"`
}

type StepExecute struct {
	Duration  StepExecuteDuration   `json:"duration,omitempty"`
	Workflows []StepExecuteWorkflow `json:"workflows,omitempty"`
}

type RetryConfig struct {
	Count int32 `json:"count,omitempty"`
	// +kubebuilder:default=any
	Success RetrySuccess `json:"success,omitempty"`
}

type ContainerConfig struct {
	//Dockerfile string `json:"dockerfile,omitempty"` TODO in future
	Image            string                        `json:"image,omitempty"`
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	ImagePullPolicy  *corev1.PullPolicy            `json:"imagePullPolicy,omitempty"`
	Resources        *Resources                    `json:"resources,omitempty"`
	Env              []corev1.EnvVar               `json:"env,omitempty"`
	EnvFrom          []corev1.EnvFromSource        `json:"envFrom,omitempty"`
	Entrypoint       *[]string                     `json:"entrypoint,omitempty"`
	Command          *[]string                     `json:"command,omitempty"`
	Args             []string                      `json:"args,omitempty"`
	Labels           map[string]string             `json:"labels,omitempty"`
	Annotations      map[string]string             `json:"annotations,omitempty"`
	SecurityContext  *corev1.SecurityContext       `json:"securityContext,omitempty"`
}

type Step struct {
	Name          string       `json:"name,omitempty"`
	Skip          string       `json:"skip,omitempty"`
	WorkingDir    string       `json:"workingDir,omitempty"`
	Retry         *RetryConfig `json:"retry,omitempty"`
	IgnoreFailure bool         `json:"ignoreFailure,omitempty"`
	Timeout       string       `json:"timeout,omitempty"`
	// +kubebuilder:validation:Pattern=^((0|[1-9][0-9]+)h)?((0|[1-9][0-9]+)m)?((0|[1-9][0-9]+)s)?((0|[1-9][0-9]+)ms)?$
	Delay string `json:"delay,omitempty"`

	Spawn     *map[string]StepSpawn `json:"spawn,omitempty"`
	Run       *StepRun              `json:"run,omitempty"`
	Cache     *StepCache            `json:"cache,omitempty"`
	ReadCache *StepReadCache        `json:"readCache,omitempty"`
	Junit     *StepJunit            `json:"junit,omitempty"`
	Execute   *StepExecute          `json:"execute,omitempty"`
}
