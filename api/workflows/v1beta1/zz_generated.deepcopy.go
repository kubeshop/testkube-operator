//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"encoding/json"
	"github.com/kubeshop/testkube-operator/api/tests/v3"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerConfig) DeepCopyInto(out *ContainerConfig) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(Resources)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Entrypoint != nil {
		in, out := &in.Entrypoint, &out.Entrypoint
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerConfig.
func (in *ContainerConfig) DeepCopy() *ContainerConfig {
	if in == nil {
		return nil
	}
	out := new(ContainerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Content) DeepCopyInto(out *Content) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(ContentGit)
		(*in).DeepCopyInto(*out)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]ContentFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Content.
func (in *Content) DeepCopy() *Content {
	if in == nil {
		return nil
	}
	out := new(Content)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentFile) DeepCopyInto(out *ContentFile) {
	*out = *in
	in.ContentFrom.DeepCopyInto(&out.ContentFrom)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentFile.
func (in *ContentFile) DeepCopy() *ContentFile {
	if in == nil {
		return nil
	}
	out := new(ContentFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentGit) DeepCopyInto(out *ContentGit) {
	*out = *in
	if in.UsernameSecret != nil {
		in, out := &in.UsernameSecret, &out.UsernameSecret
		*out = new(v3.SecretRef)
		**out = **in
	}
	if in.TokenSecret != nil {
		in, out := &in.TokenSecret, &out.TokenSecret
		*out = new(v3.SecretRef)
		**out = **in
	}
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentGit.
func (in *ContentGit) DeepCopy() *ContentGit {
	if in == nil {
		return nil
	}
	out := new(ContentGit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectRef) DeepCopyInto(out *ObjectRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectRef.
func (in *ObjectRef) DeepCopy() *ObjectRef {
	if in == nil {
		return nil
	}
	out := new(ObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
	if in.Const != nil {
		in, out := &in.Const, &out.Const
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Enum != nil {
		in, out := &in.Enum, &out.Enum
		*out = make([]json.RawMessage, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(json.RawMessage, len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.MinLength != nil {
		in, out := &in.MinLength, &out.MinLength
		*out = new(int64)
		**out = **in
	}
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(int64)
		**out = **in
	}
	if in.Minimum != nil {
		in, out := &in.Minimum, &out.Minimum
		*out = new(int64)
		**out = **in
	}
	if in.Maximum != nil {
		in, out := &in.Maximum, &out.Maximum
		*out = new(int64)
		**out = **in
	}
	if in.ExclusiveMinimum != nil {
		in, out := &in.ExclusiveMinimum, &out.ExclusiveMinimum
		*out = new(int64)
		**out = **in
	}
	if in.ExclusiveMaximum != nil {
		in, out := &in.ExclusiveMaximum, &out.ExclusiveMaximum
		*out = new(int64)
		**out = **in
	}
	if in.MultipleOf != nil {
		in, out := &in.MultipleOf, &out.MultipleOf
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryConfig) DeepCopyInto(out *RetryConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryConfig.
func (in *RetryConfig) DeepCopy() *RetryConfig {
	if in == nil {
		return nil
	}
	out := new(RetryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpawnPod) DeepCopyInto(out *SpawnPod) {
	*out = *in
	in.ContainerConfig.DeepCopyInto(&out.ContainerConfig)
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
	in.LivenessProbe.DeepCopyInto(&out.LivenessProbe)
	in.ReadinessProbe.DeepCopyInto(&out.ReadinessProbe)
	in.StartupProbe.DeepCopyInto(&out.StartupProbe)
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ContainerPort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpawnPod.
func (in *SpawnPod) DeepCopy() *SpawnPod {
	if in == nil {
		return nil
	}
	out := new(SpawnPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	if in.Retry != nil {
		in, out := &in.Retry, &out.Retry
		*out = new(RetryConfig)
		**out = **in
	}
	if in.Spawn != nil {
		in, out := &in.Spawn, &out.Spawn
		*out = new(map[string]StepSpawn)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]StepSpawn, len(*in))
			for key, val := range *in {
				var outVal *map[string]SpawnPod
				if val == nil {
					(*out)[key] = nil
				} else {
					in, out := &val, &outVal
					*out = new(map[string]SpawnPod)
					if **in != nil {
						in, out := *in, *out
						*out = make(map[string]SpawnPod, len(*in))
						for key, val := range *in {
							(*out)[key] = *val.DeepCopy()
						}
					}
				}
				(*out)[key] = outVal
			}
		}
	}
	if in.Run != nil {
		in, out := &in.Run, &out.Run
		*out = new(StepRun)
		(*in).DeepCopyInto(*out)
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(StepCache)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadCache != nil {
		in, out := &in.ReadCache, &out.ReadCache
		*out = new(StepReadCache)
		**out = **in
	}
	if in.Junit != nil {
		in, out := &in.Junit, &out.Junit
		*out = new(StepJunit)
		(*in).DeepCopyInto(*out)
	}
	if in.Execute != nil {
		in, out := &in.Execute, &out.Execute
		*out = new(StepExecute)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Step.
func (in *Step) DeepCopy() *Step {
	if in == nil {
		return nil
	}
	out := new(Step)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepArtifacts) DeepCopyInto(out *StepArtifacts) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepArtifacts.
func (in *StepArtifacts) DeepCopy() *StepArtifacts {
	if in == nil {
		return nil
	}
	out := new(StepArtifacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepCache) DeepCopyInto(out *StepCache) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepCache.
func (in *StepCache) DeepCopy() *StepCache {
	if in == nil {
		return nil
	}
	out := new(StepCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepExecute) DeepCopyInto(out *StepExecute) {
	*out = *in
	out.Duration = in.Duration
	if in.Workflows != nil {
		in, out := &in.Workflows, &out.Workflows
		*out = make([]StepExecuteWorkflow, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepExecute.
func (in *StepExecute) DeepCopy() *StepExecute {
	if in == nil {
		return nil
	}
	out := new(StepExecute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepExecuteDuration) DeepCopyInto(out *StepExecuteDuration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepExecuteDuration.
func (in *StepExecuteDuration) DeepCopy() *StepExecuteDuration {
	if in == nil {
		return nil
	}
	out := new(StepExecuteDuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepExecuteWorkflow) DeepCopyInto(out *StepExecuteWorkflow) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepExecuteWorkflow.
func (in *StepExecuteWorkflow) DeepCopy() *StepExecuteWorkflow {
	if in == nil {
		return nil
	}
	out := new(StepExecuteWorkflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepJunit) DeepCopyInto(out *StepJunit) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepJunit.
func (in *StepJunit) DeepCopy() *StepJunit {
	if in == nil {
		return nil
	}
	out := new(StepJunit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepReadCache) DeepCopyInto(out *StepReadCache) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepReadCache.
func (in *StepReadCache) DeepCopy() *StepReadCache {
	if in == nil {
		return nil
	}
	out := new(StepReadCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepRun) DeepCopyInto(out *StepRun) {
	*out = *in
	in.ContainerConfig.DeepCopyInto(&out.ContainerConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepRun.
func (in *StepRun) DeepCopy() *StepRun {
	if in == nil {
		return nil
	}
	out := new(StepRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Variable) DeepCopyInto(out *Variable) {
	*out = *in
	in.ValueFrom.DeepCopyInto(&out.ValueFrom)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Variable.
func (in *Variable) DeepCopy() *Variable {
	if in == nil {
		return nil
	}
	out := new(Variable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workflow) DeepCopyInto(out *Workflow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workflow.
func (in *Workflow) DeepCopy() *Workflow {
	if in == nil {
		return nil
	}
	out := new(Workflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workflow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowList) DeepCopyInto(out *WorkflowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowList.
func (in *WorkflowList) DeepCopy() *WorkflowList {
	if in == nil {
		return nil
	}
	out := new(WorkflowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowSpec) DeepCopyInto(out *WorkflowSpec) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(Content)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(WorkflowTemplateRef)
		(*in).DeepCopyInto(*out)
	}
	in.ContainerConfig.DeepCopyInto(&out.ContainerConfig)
	if in.Before != nil {
		in, out := &in.Before, &out.Before
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.After != nil {
		in, out := &in.After, &out.After
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowSpec.
func (in *WorkflowSpec) DeepCopy() *WorkflowSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTemplate) DeepCopyInto(out *WorkflowTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTemplate.
func (in *WorkflowTemplate) DeepCopy() *WorkflowTemplate {
	if in == nil {
		return nil
	}
	out := new(WorkflowTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTemplateList) DeepCopyInto(out *WorkflowTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkflowTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTemplateList.
func (in *WorkflowTemplateList) DeepCopy() *WorkflowTemplateList {
	if in == nil {
		return nil
	}
	out := new(WorkflowTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkflowTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTemplateRef) DeepCopyInto(out *WorkflowTemplateRef) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]json.RawMessage, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(json.RawMessage, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTemplateRef.
func (in *WorkflowTemplateRef) DeepCopy() *WorkflowTemplateRef {
	if in == nil {
		return nil
	}
	out := new(WorkflowTemplateRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTemplateSpec) DeepCopyInto(out *WorkflowTemplateSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]Parameter, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.ContainerConfig.DeepCopyInto(&out.ContainerConfig)
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTemplateSpec.
func (in *WorkflowTemplateSpec) DeepCopy() *WorkflowTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowTemplateSpec)
	in.DeepCopyInto(out)
	return out
}
