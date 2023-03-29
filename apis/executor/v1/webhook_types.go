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

// WebhookSpec defines the desired state of Webhook
type WebhookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Uri is address where webhook should be made
	Uri string `json:"uri,omitempty"`
	// Events declare list if events on which webhook should be called
	Events []EventType `json:"events,omitempty"`
	// Labels to filter for tests and test suites
	Selector string `json:"selector,omitempty"`
	// will load the generated payload for notification inside the object
	PayloadObjectField string `json:"payloadObjectField,omitempty"`
}

// +kubebuilder:validation:Enum=start-test;end-test-success;end-test-failed;end-test-aborted;end-test-timeout;start-testsuite;end-testsuite-success;end-testsuite-failed;end-testsuite-aborted;end-testsuite-timeout
type EventType string

// List of EventType
const (
	START_TEST_EventType            EventType = "start-test"
	END_TEST_SUCCESS_EventType      EventType = "end-test-success"
	END_TEST_FAILED_EventType       EventType = "end-test-failed"
	END_TEST_ABORTED_EventType      EventType = "end-test-aborted"
	END_TEST_TIMEOUT_EventType      EventType = "end-test-timeout"
	START_TESTSUITE_EventType       EventType = "start-testsuite"
	END_TESTSUITE_SUCCESS_EventType EventType = "end-testsuite-success"
	END_TESTSUITE_FAILED_EventType  EventType = "end-testsuite-failed"
	END_TESTSUITE_ABORTED_EventType EventType = "end-testsuite-aborted"
	END_TESTSUITE_TIMEOUT_EventType EventType = "end-testsuite-timeout"
)

// WebhookStatus defines the observed state of Webhook
type WebhookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Webhook is the Schema for the webhooks API
type Webhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebhookSpec   `json:"spec,omitempty"`
	Status WebhookStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WebhookList contains a list of Webhook
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webhook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Webhook{}, &WebhookList{})
}
