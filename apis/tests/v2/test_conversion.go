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

package v2

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/kubeshop/testkube-operator/apis/tests/v1"
)

func convertTestStepSpecTo(step TestStepSpec) v1.TestStepSpec {
	var execute *v1.TestStepExecute
	var delay *v1.TestStepDelay

	if step.Execute != nil {
		execute = &v1.TestStepExecute{
			Namespace:     step.Execute.Namespace,
			Name:          step.Execute.Name,
			StopOnFailure: step.Execute.StopOnFailure,
		}
	}

	if step.Delay != nil {
		delay = &v1.TestStepDelay{
			Duration: step.Delay.Duration,
		}
	}

	return v1.TestStepSpec{
		Type:    step.Type,
		Execute: execute,
		Delay:   delay,
	}
}

func convertTestStepSpecFrom(step v1.TestStepSpec) TestStepSpec {
	var execute *TestStepExecute
	var delay *TestStepDelay

	if step.Execute != nil {
		execute = &TestStepExecute{
			Namespace:     step.Execute.Namespace,
			Name:          step.Execute.Name,
			StopOnFailure: step.Execute.StopOnFailure,
		}
	}

	if step.Delay != nil {
		delay = &TestStepDelay{
			Duration: step.Delay.Duration,
		}
	}

	return TestStepSpec{
		Type:    step.Type,
		Execute: execute,
		Delay:   delay,
	}
}

// ConvertTo converts this Test to the Hub version (v1).
func (src *Test) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.Test)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Repeats = src.Spec.Repeats
	dst.Spec.Description = src.Spec.Description
	dst.Spec.Tags = src.Spec.Tags

	for _, step := range src.Spec.Before {
		dst.Spec.Before = append(dst.Spec.Before, convertTestStepSpecTo(step))
	}

	for _, step := range src.Spec.Steps {
		dst.Spec.Steps = append(dst.Spec.Steps, convertTestStepSpecTo(step))
	}

	for _, step := range src.Spec.After {
		dst.Spec.After = append(dst.Spec.After, convertTestStepSpecTo(step))
	}

	// Status

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *Test) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.Test)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Repeats = src.Spec.Repeats
	dst.Spec.Description = src.Spec.Description
	dst.Spec.URI = "n/a"
	dst.Spec.Tags = src.Spec.Tags

	for _, step := range src.Spec.Before {
		dst.Spec.Before = append(dst.Spec.Before, convertTestStepSpecFrom(step))
	}

	for _, step := range src.Spec.Steps {
		dst.Spec.Steps = append(dst.Spec.Steps, convertTestStepSpecFrom(step))
	}

	for _, step := range src.Spec.After {
		dst.Spec.After = append(dst.Spec.After, convertTestStepSpecFrom(step))
	}

	// Status
	return nil
}
