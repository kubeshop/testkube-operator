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
	commonv1 "github.com/kubeshop/testkube-operator/apis/common/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	// config files to be copied from host of form source:destination
	ConfigFiles []string `json:"configFiles,omitempty"`
}

type Variable commonv1.Variable

// TestContent defines test content
type TestContent struct {
	// test type
	Type_ string `json:"type,omitempty"`
	// repository of test content
	Repository *Repository `json:"repository,omitempty"`
	// test content body
	Data string `json:"data,omitempty"`
	// uri of test content
	Uri string `json:"uri,omitempty"`
}

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
	Type_ string `json:"type"`
	// uri of content file or git directory
	Uri string `json:"uri"`
	// branch/tag name for checkout
	Branch string `json:"branch,omitempty"`
	// commit id (sha) for checkout
	Commit string `json:"commit,omitempty"`
	// if needed we can checkout particular path (dir or file) in case of BIG/mono repositories
	Path           string     `json:"path,omitempty"`
	UsernameSecret *SecretRef `json:"usernameSecret,omitempty"`
	TokenSecret    *SecretRef `json:"tokenSecret,omitempty"`
}

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
	VariablesFile string              `json:"variablesFile,omitempty"`
	Variables     map[string]Variable `json:"variables,omitempty"`
	// test secret uuid
	TestSecretUUID string `json:"testSecretUUID,omitempty"`
	// test suite secret uuid, if it's run as a part of test suite
	TestSuiteSecretUUID string `json:"testSuiteSecretUUID,omitempty"`
	// additional executor binary arguments
	Args []string `json:"args,omitempty"`
	// container executor binary command
	Command []string `json:"command,omitempty"`
	// container executor image
	Image string `json:"image,omitempty"`
	// container executor image pull secrets
	ImagePullSecrets []v1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	// environment variables passed to executor
	Envs map[string]string `json:"envs,omitempty"`
	// execution variables passed to executor from secrets
	SecretEnvs map[string]string `json:"secretEnvs,omitempty"`
	// whether to start execution sync or async
	Sync bool `json:"sync,omitempty"`
	// http proxy for executor containers
	HttpProxy string `json:"httpProxy,omitempty"`
	// https proxy for executor containers
	HttpsProxy string `json:"httpsProxy,omitempty"`
}

// TestStatus defines the observed state of Test
type TestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	LastExecution   metav1.Time `json:"last_execution,omitempty"`
	ExecutionsCount int         `json:"executions_count,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// Test is the Schema for the tests API
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
