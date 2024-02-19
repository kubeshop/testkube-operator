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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestWorkflowTemplateSpec defines the desired state of TestWorkflow
type TestWorkflowTemplateSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	TestWorkflowSpecBase `json:",inline"`

	// steps for setting up the workflow
	Setup []IndependentStep `json:"setup,omitempty"`

	// steps to execute in the workflow
	// +kubebuilder:validation:MinItems=1
	Steps []IndependentStep `json:"steps"`

	// steps to run at the end of the workflow
	After []IndependentStep `json:"after,omitempty"`
}

// +kubebuilder:object:root=true

// TestWorkflowTemplate is the Schema for the workflows API
type TestWorkflowTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// TestWorkflowTemplate readable description
	Description string `json:"description,omitempty"`

	// TestWorkflowTemplate specification
	Spec TestWorkflowTemplateSpec `json:"spec"`
}

//+kubebuilder:object:root=true

// TestWorkflowTemplateList contains a list of TestWorkflowTemplate
type TestWorkflowTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestWorkflowTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestWorkflowTemplate{}, &TestWorkflowTemplateList{})
}
