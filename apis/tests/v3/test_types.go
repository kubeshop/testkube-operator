/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v3

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	commonv1 "github.com/kubeshop/testkube-operator/apis/common/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TestSpec defines the desired state of Test
type TestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// test type
	Type_ string `json:"type,omitempty"`
	// test name
	Name string `json:"name,omitempty"`
	// test content object
	Content *TestContent `json:"content,omitempty"`
	// reference to test source resource
	Source string `json:"source,omitempty"`
	// schedule in cron job format for scheduled test execution
	Schedule         string            `json:"schedule,omitempty"`
	ExecutionRequest *ExecutionRequest `json:"executionRequest,omitempty"`
	// files to be used from minio uploads
	Uploads []string `json:"uploads,omitempty"`
}

type Variable commonv1.Variable

// TestContent defines test content
type TestContent struct {
	// test type
	Type_ TestContentType `json:"type,omitempty"`
	// repository of test content
	Repository *Repository `json:"repository,omitempty"`
	// test content body
	Data string `json:"data,omitempty"`
	// uri of test content
	Uri string `json:"uri,omitempty"`
}

// +kubebuilder:validation:Enum=string;file-uri;git-file;git-dir;git
type TestContentType string

const (
	TestContentTypeString  TestContentType = "string"
	TestContentTypeFileURI TestContentType = "file-uri"
	// Deprecated: use git instead
	TestContentTypeGitFile TestContentType = "git-file"
	// Deprecated: use git instead
	TestContentTypeGitDir TestContentType = "git-dir"
	TestContentTypeGit    TestContentType = "git"
)

// Testkube internal reference for secret storage in Kubernetes secrets
type SecretRef struct {
	// object kubernetes namespace
	Namespace string `json:"namespace,omitempty"`
	// object name
	Name string `json:"name"`
	// object key
	Key string `json:"key"`
}

// Repository represents VCS repo, currently we're handling Git only
type Repository struct {
	// VCS repository type
	Type_ string `json:"type,omitempty"`
	// uri of content file or git directory
	Uri string `json:"uri,omitempty"`
	// branch/tag name for checkout
	Branch string `json:"branch,omitempty"`
	// commit id (sha) for checkout
	Commit string `json:"commit,omitempty"`
	// if needed we can checkout particular path (dir or file) in case of BIG/mono repositories
	Path           string     `json:"path,omitempty"`
	UsernameSecret *SecretRef `json:"usernameSecret,omitempty"`
	TokenSecret    *SecretRef `json:"tokenSecret,omitempty"`
	// git auth certificate secret for private repositories
	CertificateSecret string `json:"certificateSecret,omitempty"`
	// if provided we checkout the whole repository and run test from this directory
	WorkingDir string `json:"workingDir,omitempty"`
	// auth type for git requests
	AuthType GitAuthType `json:"authType,omitempty"`
}

// GitAuthType defines git auth type
// +kubebuilder:validation:Enum=basic;header
type GitAuthType string

const (
	// GitAuthTypeBasic for git basic auth requests
	GitAuthTypeBasic GitAuthType = "basic"
	// GitAuthTypeHeader for git header auth requests
	GitAuthTypeHeader GitAuthType = "header"
)

// artifact request body with test artifacts
type ArtifactRequest struct {
	// artifact storage class name for container executor
	StorageClassName string `json:"storageClassName"`
	// artifact volume mount path for container executor
	VolumeMountPath string `json:"volumeMountPath"`
	// artifact directories for scraping
	Dirs []string `json:"dirs,omitempty"`
}

// running context for test or test suite execution
type RunningContext struct {
	// One of possible context types
	Type_ RunningContextType `json:"type"`
	// Context value depending from its type
	Context string `json:"context,omitempty"`
}

type RunningContextType string

const (
	RunningContextTypeUserCLI     RunningContextType = "user-cli"
	RunningContextTypeUserUI      RunningContextType = "user-ui"
	RunningContextTypeTestSuite   RunningContextType = "testsuite"
	RunningContextTypeTestTrigger RunningContextType = "testtrigger"
	RunningContextTypeScheduler   RunningContextType = "scheduler"
	RunningContextTypeEmpty       RunningContextType = ""
)

// test execution request body
type ExecutionRequest struct {
	// test execution custom name
	Name string `json:"name,omitempty"`
	// unique test suite name (CRD Test suite name), if it's run as a part of test suite
	TestSuiteName string `json:"testSuiteName,omitempty"`
	// test execution number
	Number int32 `json:"number,omitempty"`
	// test execution labels
	ExecutionLabels map[string]string `json:"executionLabels,omitempty"`
	// test kubernetes namespace (\"testkube\" when not set)
	Namespace string `json:"namespace,omitempty"`
	// variables file content - need to be in format for particular executor (e.g. postman envs file)
	VariablesFile           string              `json:"variablesFile,omitempty"`
	IsVariablesFileUploaded bool                `json:"isVariablesFileUploaded,omitempty"`
	Variables               map[string]Variable `json:"variables,omitempty"`
	// test secret uuid
	TestSecretUUID string `json:"testSecretUUID,omitempty"`
	// test suite secret uuid, if it's run as a part of test suite
	TestSuiteSecretUUID string `json:"testSuiteSecretUUID,omitempty"`
	// additional executor binary arguments
	Args []string `json:"args,omitempty"`
	// usage mode for arguments
	ArgsMode ArgsModeType `json:"argsMode,omitempty"`
	// executor binary command
	Command []string `json:"command,omitempty"`
	// container executor image
	Image string `json:"image,omitempty"`
	// container executor image pull secrets
	ImagePullSecrets []v1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	// Environment variables passed to executor.
	// Deprecated: use Basic Variables instead
	Envs map[string]string `json:"envs,omitempty"`
	// Execution variables passed to executor from secrets.
	// Deprecated: use Secret Variables instead
	SecretEnvs map[string]string `json:"secretEnvs,omitempty"`
	// whether to start execution sync or async
	Sync bool `json:"sync,omitempty"`
	// http proxy for executor containers
	HttpProxy string `json:"httpProxy,omitempty"`
	// https proxy for executor containers
	HttpsProxy string `json:"httpsProxy,omitempty"`
	// negative test will fail the execution if it is a success and it will succeed if it is a failure
	NegativeTest bool `json:"negativeTest,omitempty"`
	// Optional duration in seconds the pod may be active on the node relative to
	// StartTime before the system will actively try to mark it failed and kill associated containers.
	// Value must be a positive integer.
	ActiveDeadlineSeconds int64            `json:"activeDeadlineSeconds,omitempty"`
	ArtifactRequest       *ArtifactRequest `json:"artifactRequest,omitempty"`
	// job template extensions
	JobTemplate string `json:"jobTemplate,omitempty"`
	// cron job template extensions
	CronJobTemplate string `json:"cronJobTemplate,omitempty"`
	// script to run before test execution
	PreRunScript string `json:"preRunScript,omitempty"`
	// script to run after test execution
	PostRunScript string `json:"postRunScript,omitempty"`
	// scraper template extensions
	ScraperTemplate string `json:"scraperTemplate,omitempty"`
	// config map references
	EnvConfigMaps []EnvReference `json:"envConfigMaps,omitempty"`
	// secret references
	EnvSecrets     []EnvReference  `json:"envSecrets,omitempty"`
	RunningContext *RunningContext `json:"runningContext,omitempty"`
}

// ArgsModeType defines args mode type
// +kubebuilder:validation:Enum=append;override
type ArgsModeType string

const (
	// ArgsModeTypeAppend for append args mode
	ArgsModeTypeAppend ArgsModeType = "append"
	// ArgsModeTypeOverride for override args mode
	ArgsModeTypeOverride ArgsModeType = "override"
)

// Reference to env resource
type EnvReference struct {
	v1.LocalObjectReference `json:"reference"`
	// whether we shoud mount resource
	Mount bool `json:"mount,omitempty"`
	// where we shoud mount resource
	MountPath string `json:"mountPath,omitempty"`
	// whether we shoud map to variables from resource
	MapToVariables bool `json:"mapToVariables,omitempty"`
}

// +kubebuilder:validation:Enum=queued;running;passed;failed;aborted;timeout
type ExecutionStatus string

// List of ExecutionStatus
const (
	QUEUED_ExecutionStatus  ExecutionStatus = "queued"
	RUNNING_ExecutionStatus ExecutionStatus = "running"
	PASSED_ExecutionStatus  ExecutionStatus = "passed"
	FAILED_ExecutionStatus  ExecutionStatus = "failed"
	ABORTED_ExecutionStatus ExecutionStatus = "aborted"
	TIMEOUT_ExecutionStatus ExecutionStatus = "timeout"
)

// test execution core
type ExecutionCore struct {
	// execution id
	Id string `json:"id,omitempty"`
	// execution number
	Number int32 `json:"number,omitempty"`
	// test start time
	StartTime metav1.Time `json:"startTime,omitempty"`
	// test end time
	EndTime metav1.Time      `json:"endTime,omitempty"`
	Status  *ExecutionStatus `json:"status,omitempty"`
}

// TestStatus defines the observed state of Test
type TestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// latest execution result
	LatestExecution *ExecutionCore `json:"latestExecution,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// Test is the Schema for the tests API
// +kubebuilder:printcolumn:name="Type",type=string,JSONPath=`.spec.type`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type Test struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestSpec   `json:"spec,omitempty"`
	Status TestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TestList contains a list of Test
type TestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Test `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Test{}, &TestList{})
}
