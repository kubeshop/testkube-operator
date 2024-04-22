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

package v1

import (
	"github.com/kubeshop/testkube-operator/api/tests/v3"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactCompression) DeepCopyInto(out *ArtifactCompression) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactCompression.
func (in *ArtifactCompression) DeepCopy() *ArtifactCompression {
	if in == nil {
		return nil
	}
	out := new(ArtifactCompression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerConfig) DeepCopyInto(out *ContainerConfig) {
	*out = *in
	if in.WorkingDir != nil {
		in, out := &in.WorkingDir, &out.WorkingDir
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
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
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(Resources)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.Tarball != nil {
		in, out := &in.Tarball, &out.Tarball
		*out = make([]ContentTarball, len(*in))
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
	if in.ContentFrom != nil {
		in, out := &in.ContentFrom, &out.ContentFrom
		*out = new(corev1.EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(int32)
		**out = **in
	}
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
	if in.UsernameFrom != nil {
		in, out := &in.UsernameFrom, &out.UsernameFrom
		*out = new(corev1.EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
	if in.TokenFrom != nil {
		in, out := &in.TokenFrom, &out.TokenFrom
		*out = new(corev1.EnvVarSource)
		(*in).DeepCopyInto(*out)
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
func (in *ContentTarball) DeepCopyInto(out *ContentTarball) {
	*out = *in
	if in.Mount != nil {
		in, out := &in.Mount, &out.Mount
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentTarball.
func (in *ContentTarball) DeepCopy() *ContentTarball {
	if in == nil {
		return nil
	}
	out := new(ContentTarball)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobConfig) DeepCopyInto(out *CronJobConfig) {
	*out = *in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobConfig.
func (in *CronJobConfig) DeepCopy() *CronJobConfig {
	if in == nil {
		return nil
	}
	out := new(CronJobConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicList) DeepCopyInto(out *DynamicList) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicList.
func (in *DynamicList) DeepCopy() *DynamicList {
	if in == nil {
		return nil
	}
	out := new(DynamicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Event) DeepCopyInto(out *Event) {
	*out = *in
	if in.Cronjob != nil {
		in, out := &in.Cronjob, &out.Cronjob
		*out = new(CronJobConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Event.
func (in *Event) DeepCopy() *Event {
	if in == nil {
		return nil
	}
	out := new(Event)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndependentStep) DeepCopyInto(out *IndependentStep) {
	*out = *in
	in.StepBase.DeepCopyInto(&out.StepBase)
	if in.Setup != nil {
		in, out := &in.Setup, &out.Setup
		*out = make([]IndependentStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]IndependentStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndependentStep.
func (in *IndependentStep) DeepCopy() *IndependentStep {
	if in == nil {
		return nil
	}
	out := new(IndependentStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobConfig) DeepCopyInto(out *JobConfig) {
	*out = *in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobConfig.
func (in *JobConfig) DeepCopy() *JobConfig {
	if in == nil {
		return nil
	}
	out := new(JobConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterNumberSchema) DeepCopyInto(out *ParameterNumberSchema) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterNumberSchema.
func (in *ParameterNumberSchema) DeepCopy() *ParameterNumberSchema {
	if in == nil {
		return nil
	}
	out := new(ParameterNumberSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterSchema) DeepCopyInto(out *ParameterSchema) {
	*out = *in
	if in.Enum != nil {
		in, out := &in.Enum, &out.Enum
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Example != nil {
		in, out := &in.Example, &out.Example
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(intstr.IntOrString)
		**out = **in
	}
	in.ParameterStringSchema.DeepCopyInto(&out.ParameterStringSchema)
	in.ParameterNumberSchema.DeepCopyInto(&out.ParameterNumberSchema)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterSchema.
func (in *ParameterSchema) DeepCopy() *ParameterSchema {
	if in == nil {
		return nil
	}
	out := new(ParameterSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterStringSchema) DeepCopyInto(out *ParameterStringSchema) {
	*out = *in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterStringSchema.
func (in *ParameterStringSchema) DeepCopy() *ParameterStringSchema {
	if in == nil {
		return nil
	}
	out := new(ParameterStringSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PauseConfig) DeepCopyInto(out *PauseConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PauseConfig.
func (in *PauseConfig) DeepCopy() *PauseConfig {
	if in == nil {
		return nil
	}
	out := new(PauseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodConfig) DeepCopyInto(out *PodConfig) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodConfig.
func (in *PodConfig) DeepCopy() *PodConfig {
	if in == nil {
		return nil
	}
	out := new(PodConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(map[corev1.ResourceName]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(map[corev1.ResourceName]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
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
func (in *RetryPolicy) DeepCopyInto(out *RetryPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryPolicy.
func (in *RetryPolicy) DeepCopy() *RetryPolicy {
	if in == nil {
		return nil
	}
	out := new(RetryPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	in.StepBase.DeepCopyInto(&out.StepBase)
	if in.Use != nil {
		in, out := &in.Use, &out.Use
		*out = make([]TemplateRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(TemplateRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Setup != nil {
		in, out := &in.Setup, &out.Setup
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
	if in.WorkingDir != nil {
		in, out := &in.WorkingDir, &out.WorkingDir
		*out = new(string)
		**out = **in
	}
	if in.Compress != nil {
		in, out := &in.Compress, &out.Compress
		*out = new(ArtifactCompression)
		**out = **in
	}
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
func (in *StepBase) DeepCopyInto(out *StepBase) {
	*out = *in
	if in.Pause != nil {
		in, out := &in.Pause, &out.Pause
		*out = new(PauseConfig)
		**out = **in
	}
	if in.Retry != nil {
		in, out := &in.Retry, &out.Retry
		*out = new(RetryPolicy)
		**out = **in
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(Content)
		(*in).DeepCopyInto(*out)
	}
	if in.Run != nil {
		in, out := &in.Run, &out.Run
		*out = new(StepRun)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkingDir != nil {
		in, out := &in.WorkingDir, &out.WorkingDir
		*out = new(string)
		**out = **in
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(ContainerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Execute != nil {
		in, out := &in.Execute, &out.Execute
		*out = new(StepExecute)
		(*in).DeepCopyInto(*out)
	}
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = new(StepArtifacts)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepBase.
func (in *StepBase) DeepCopy() *StepBase {
	if in == nil {
		return nil
	}
	out := new(StepBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepExecute) DeepCopyInto(out *StepExecute) {
	*out = *in
	if in.Tests != nil {
		in, out := &in.Tests, &out.Tests
		*out = make([]StepExecuteTest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Workflows != nil {
		in, out := &in.Workflows, &out.Workflows
		*out = make([]StepExecuteWorkflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
func (in *StepExecuteStrategy) DeepCopyInto(out *StepExecuteStrategy) {
	*out = *in
	if in.Matrix != nil {
		in, out := &in.Matrix, &out.Matrix
		*out = make(map[string]DynamicList, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = make(map[string]DynamicList, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepExecuteStrategy.
func (in *StepExecuteStrategy) DeepCopy() *StepExecuteStrategy {
	if in == nil {
		return nil
	}
	out := new(StepExecuteStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepExecuteTest) DeepCopyInto(out *StepExecuteTest) {
	*out = *in
	in.StepExecuteStrategy.DeepCopyInto(&out.StepExecuteStrategy)
	if in.Tarball != nil {
		in, out := &in.Tarball, &out.Tarball
		*out = make(map[string]TarballRequest, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ExecutionRequest != nil {
		in, out := &in.ExecutionRequest, &out.ExecutionRequest
		*out = new(TestExecutionRequest)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepExecuteTest.
func (in *StepExecuteTest) DeepCopy() *StepExecuteTest {
	if in == nil {
		return nil
	}
	out := new(StepExecuteTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepExecuteWorkflow) DeepCopyInto(out *StepExecuteWorkflow) {
	*out = *in
	in.StepExecuteStrategy.DeepCopyInto(&out.StepExecuteStrategy)
	if in.Tarball != nil {
		in, out := &in.Tarball, &out.Tarball
		*out = make(map[string]TarballRequest, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
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
func (in *StepRun) DeepCopyInto(out *StepRun) {
	*out = *in
	in.ContainerConfig.DeepCopyInto(&out.ContainerConfig)
	if in.Shell != nil {
		in, out := &in.Shell, &out.Shell
		*out = new(string)
		**out = **in
	}
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
func (in *TarballRequest) DeepCopyInto(out *TarballRequest) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = new(DynamicList)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TarballRequest.
func (in *TarballRequest) DeepCopy() *TarballRequest {
	if in == nil {
		return nil
	}
	out := new(TarballRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateRef) DeepCopyInto(out *TemplateRef) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateRef.
func (in *TemplateRef) DeepCopy() *TemplateRef {
	if in == nil {
		return nil
	}
	out := new(TemplateRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestExecutionRequest) DeepCopyInto(out *TestExecutionRequest) {
	*out = *in
	if in.ExecutionLabels != nil {
		in, out := &in.ExecutionLabels, &out.ExecutionLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]v3.Variable, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.ArtifactRequest != nil {
		in, out := &in.ArtifactRequest, &out.ArtifactRequest
		*out = new(v3.ArtifactRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvConfigMaps != nil {
		in, out := &in.EnvConfigMaps, &out.EnvConfigMaps
		*out = make([]v3.EnvReference, len(*in))
		copy(*out, *in)
	}
	if in.EnvSecrets != nil {
		in, out := &in.EnvSecrets, &out.EnvSecrets
		*out = make([]v3.EnvReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestExecutionRequest.
func (in *TestExecutionRequest) DeepCopy() *TestExecutionRequest {
	if in == nil {
		return nil
	}
	out := new(TestExecutionRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestWorkflow) DeepCopyInto(out *TestWorkflow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestWorkflow.
func (in *TestWorkflow) DeepCopy() *TestWorkflow {
	if in == nil {
		return nil
	}
	out := new(TestWorkflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestWorkflow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestWorkflowList) DeepCopyInto(out *TestWorkflowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TestWorkflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestWorkflowList.
func (in *TestWorkflowList) DeepCopy() *TestWorkflowList {
	if in == nil {
		return nil
	}
	out := new(TestWorkflowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestWorkflowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestWorkflowSpec) DeepCopyInto(out *TestWorkflowSpec) {
	*out = *in
	if in.Use != nil {
		in, out := &in.Use, &out.Use
		*out = make([]TemplateRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.TestWorkflowSpecBase.DeepCopyInto(&out.TestWorkflowSpecBase)
	if in.Setup != nil {
		in, out := &in.Setup, &out.Setup
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestWorkflowSpec.
func (in *TestWorkflowSpec) DeepCopy() *TestWorkflowSpec {
	if in == nil {
		return nil
	}
	out := new(TestWorkflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestWorkflowSpecBase) DeepCopyInto(out *TestWorkflowSpecBase) {
	*out = *in
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]Event, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]ParameterSchema, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(Content)
		(*in).DeepCopyInto(*out)
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(ContainerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Job != nil {
		in, out := &in.Job, &out.Job
		*out = new(JobConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(PodConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestWorkflowSpecBase.
func (in *TestWorkflowSpecBase) DeepCopy() *TestWorkflowSpecBase {
	if in == nil {
		return nil
	}
	out := new(TestWorkflowSpecBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestWorkflowTemplate) DeepCopyInto(out *TestWorkflowTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestWorkflowTemplate.
func (in *TestWorkflowTemplate) DeepCopy() *TestWorkflowTemplate {
	if in == nil {
		return nil
	}
	out := new(TestWorkflowTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestWorkflowTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestWorkflowTemplateList) DeepCopyInto(out *TestWorkflowTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TestWorkflowTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestWorkflowTemplateList.
func (in *TestWorkflowTemplateList) DeepCopy() *TestWorkflowTemplateList {
	if in == nil {
		return nil
	}
	out := new(TestWorkflowTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestWorkflowTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestWorkflowTemplateSpec) DeepCopyInto(out *TestWorkflowTemplateSpec) {
	*out = *in
	in.TestWorkflowSpecBase.DeepCopyInto(&out.TestWorkflowSpecBase)
	if in.Setup != nil {
		in, out := &in.Setup, &out.Setup
		*out = make([]IndependentStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]IndependentStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.After != nil {
		in, out := &in.After, &out.After
		*out = make([]IndependentStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestWorkflowTemplateSpec.
func (in *TestWorkflowTemplateSpec) DeepCopy() *TestWorkflowTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TestWorkflowTemplateSpec)
	in.DeepCopyInto(out)
	return out
}
