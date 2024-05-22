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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// TestWorkflowExecutionSpec defines the desired state of TestWorkflowExecution
type TestWorkflowExecutionSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	TestWorkflow     *corev1.LocalObjectReference  `json:"testWorkflow" expr:"include"`
	ExecutionRequest *TestWorkflowExecutionRequest `json:"executionRequest,omitempty" expr:"include"`
}

// TestWorkflowExecutionRequest contains TestWorkflow execution parameters
type TestWorkflowExecutionRequest struct {
	// custom execution name
	Name   string                        `json:"name,omitempty" expr:"template"`
	Config map[string]intstr.IntOrString `json:"config,omitempty" expr:"template"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TestWorkflowExecution is the Schema for the workflows API
type TestWorkflowExecution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// TestWorkflowExecution specification
	Spec TestWorkflowTemplateSpec `json:"spec" expr:"include"`
	// TestWorkflowExecutionStatus soecification
	Status *TestWorkflowExecutionStatus `json:"status,omitempty"`
}

// test workflow execution status
type TestWorkflowExecutionStatus struct {
	LatestExecution *TestWorkflowExecution `json:"latestExecution,omitempty"`
	// test workflow execution generation
	Generation int64 `json:"generation,omitempty"`
}

//+kubebuilder:object:root=true

// TestWorkflowExecutionList contains a list of TestWorkflowExecutiom
type TestWorkflowExecutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestWorkflowExecution `json:"items" expr:"include"`
}

func init() {
	SchemeBuilder.Register(&TestWorkflowExecution{}, &TestWorkflowExecutionList{})
}
