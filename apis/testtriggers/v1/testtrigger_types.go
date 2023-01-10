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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TestTrigger is the Schema for the testtriggers API
type TestTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestTriggerSpec   `json:"spec,omitempty"`
	Status TestTriggerStatus `json:"status,omitempty"`
}

// TestTriggerSpec defines the desired state of TestTrigger
type TestTriggerSpec struct {
	// For which Resource do we monitor Event which triggers an Action on certain comditions
	Resource string `json:"resource"`
	// ResourceSelector identifies which Kubernetes Objects should be watched
	ResourceSelector TestTriggerSelector `json:"resourceSelector"`
	// On which Event for a Resource should an Action be triggered
	Event string `json:"event"`
	// What resource conditions should be matched
	Conditions []TestTriggerCondition `json:"conditions,omitempty"`
	// Action represents what needs to be executed for selected Execution
	Action string `json:"action"`
	// Execution identifies for which test execution should an Action be executed
	Execution string `json:"execution"`
	// TestSelector identifies on which Testkube Kubernetes Objects an Action should be taken
	TestSelector TestTriggerSelector `json:"testSelector"`
}

// TestTriggerSelector is used for selecting Kubernetes Objects
type TestTriggerSelector struct {
	// Name selector is used to identify a Kubernetes Object based on the metadata name
	Name string `json:"name,omitempty"`
	// Namespace of the Kubernetes object
	Namespace string `json:"namespace,omitempty"`
	// LabelSelector is used to identify a group of Kubernetes Objects based on their metadata labels
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty"`
}

// TestTriggerStatus defines the observed state of TestTrigger
type TestTriggerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// TestTriggerConditionStatuses defines condition statuses for test triggers
type TestTriggerConditionStatuses string

// List of TestTriggerConditionStatuses
const (
	TRUE_TestTriggerConditionStatuses    TestTriggerConditionStatuses = "True"
	FALSE_TestTriggerConditionStatuses   TestTriggerConditionStatuses = "False"
	UNKNOWN_TestTriggerConditionStatuses TestTriggerConditionStatuses = "Unknown"
)

// TestTriggerCondition is used for definition of condition for test triggers
type TestTriggerCondition struct {
	Status *TestTriggerConditionStatuses `json:"status,omitempty"`
	// test trigger condition
	Type_ string `json:"type,omitempty"`
}

//+kubebuilder:object:root=true

// TestTriggerList contains a list of TestTrigger
type TestTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestTrigger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestTrigger{}, &TestTriggerList{})
}
