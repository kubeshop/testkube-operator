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

import "github.com/kubeshop/testkube-operator/pkg/informers/externalversions/internalinterfaces"

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Executor returns an ExecutorInformer.
	Executor() ExecutorInformer
	// Webhook returns a WebhookInformer.
	Webhook() WebhookInformer
	// WebhookTemplate returns a WebhookTemplateInformer.
	WebhookTemplate() WebhookTemplateInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(
	f internalinterfaces.SharedInformerFactory,
	namespace string,
	tweakListOptions internalinterfaces.TweakListOptionsFunc,
) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Executor returns an ExecutorInformer.
func (v *version) Executor() ExecutorInformer {
	return &executorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Webhook returns a WebhookInformer.
func (v *version) Webhook() WebhookInformer {
	return &webhookInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WebhookTemplate returns a WebhookTemplateInformer.
func (v *version) WebhookTemplate() WebhookTemplateInformer {
	return &webhookTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
