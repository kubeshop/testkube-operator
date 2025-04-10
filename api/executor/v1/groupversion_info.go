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

// Package v1 contains API Schema definitions for the executor v1 API group
// +kubebuilder:object:generate=true
// +groupName=executor.testkube.io
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// Group represents the API Group
	Group = "executor.testkube.io"

	// Version represents the Resource version
	Version = "v1"

	// ExecutorResource corresponds to the CRD Kind
	ExecutorResource = "Executor"

	// WebhookResource corresponds to the CRD Kind
	WebhookResource = "Webhook"

	// WebhookTemplateResource corresponds to the CRD Kind
	WebhookTemplateResource = "WebhookTemplate"

	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// ExecutorGroupVersionResource is group, version and resource used to register these objects
	ExecutorGroupVersionResource = schema.GroupVersionResource{Group: Group, Version: Version, Resource: ExecutorResource}

	// WebhookGroupVersionResource is group, version and resource used to register these objects
	WebhookGroupVersionResource = schema.GroupVersionResource{Group: Group, Version: Version, Resource: WebhookResource}

	// WebhookTemplateGroupVersionResource is group, version and resource used to register these objects
	WebhookTemplateGroupVersionResource = schema.GroupVersionResource{Group: Group, Version: Version, Resource: WebhookTemplateResource}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
