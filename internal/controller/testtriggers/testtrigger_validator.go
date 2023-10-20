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

package testtriggers

import (
	"context"

	testtriggerv1 "github.com/kubeshop/testkube-operator/api/testtriggers/v1"
	"github.com/kubeshop/testkube-operator/internal/utils"
	"github.com/kubeshop/testkube-operator/pkg/validation/tests/v1/testtrigger"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Validator struct {
	c client.Client
}

func NewValidator(c client.Client) *Validator {
	return &Validator{c: c}
}

func (v *Validator) ValidateCreate(ctx context.Context, t *testtriggerv1.TestTrigger) error {
	var allErrs field.ErrorList

	if err := v.validateResource(t.Spec.Resource); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateAction(t.Spec.Action); err != nil {
		allErrs = append(allErrs, err)
	}

	if errs := v.validateConditions(t.Spec.ConditionSpec); errs != nil {
		allErrs = append(allErrs, errs...)
	}

	if errs := v.validateProbes(t.Spec.ProbeSpec); errs != nil {
		allErrs = append(allErrs, errs...)
	}

	if err := v.validateExecution(t.Spec.Execution); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateConcurrencyPolicy(t.Spec.ConcurrencyPolicy); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateResourceSelector(t.Spec.ResourceSelector); err != nil {
		allErrs = append(allErrs, err...)
	}

	if err := v.validateTestSelector(t.Spec.TestSelector); err != nil {
		allErrs = append(allErrs, err...)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return k8serrors.NewInvalid(
		schema.GroupKind{
			Group: testtriggerv1.GroupVersion.Group,
			Kind:  "TestTrigger",
		},
		t.Name,
		allErrs,
	)
}

func (v *Validator) ValidateUpdate(ctx context.Context, old runtime.Object, new *testtriggerv1.TestTrigger) error {
	var allErrs field.ErrorList

	if err := v.validateResource(new.Spec.Resource); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateAction(new.Spec.Action); err != nil {
		allErrs = append(allErrs, err)
	}

	if errs := v.validateConditions(new.Spec.ConditionSpec); errs != nil {
		allErrs = append(allErrs, errs...)
	}

	if errs := v.validateProbes(new.Spec.ProbeSpec); errs != nil {
		allErrs = append(allErrs, errs...)
	}

	if err := v.validateExecution(new.Spec.Execution); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateConcurrencyPolicy(new.Spec.ConcurrencyPolicy); err != nil {
		allErrs = append(allErrs, err)
	}

	if errs := v.validateResourceSelector(new.Spec.ResourceSelector); errs != nil {
		allErrs = append(allErrs, errs...)
	}

	if err := v.validateTestSelector(new.Spec.TestSelector); err != nil {
		allErrs = append(allErrs, err...)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return k8serrors.NewInvalid(
		schema.GroupKind{
			Group: testtriggerv1.GroupVersion.Group,
			Kind:  testtriggerv1.Resource,
		},
		new.Name,
		allErrs,
	)
}

func (v *Validator) ValidateDelete(ctx context.Context, trigger *testtriggerv1.TestTrigger) error {
	return nil
}

func (v *Validator) validateResourceSelector(resourceSelector testtriggerv1.TestTriggerSelector) field.ErrorList {
	fld := field.NewPath("spec").Child("testSelector")
	return v.validateSelector(fld, resourceSelector)
}

func (v *Validator) validateTestSelector(testSelector testtriggerv1.TestTriggerSelector) field.ErrorList {
	var allErrs field.ErrorList

	fld := field.NewPath("spec").Child("testSelector")
	if err := v.validateSelector(fld, testSelector); err != nil {
		allErrs = append(allErrs, err...)
	}

	return allErrs
}

func (v *Validator) validateSelector(fld *field.Path, selector testtriggerv1.TestTriggerSelector) field.ErrorList {
	var allErrs field.ErrorList

	isLabelSelectorEmpty := true

	if selector.LabelSelector != nil {
		isEmpty, verr := validateLabelSelector(selector.LabelSelector, fld.Child("labelSelector"))
		if verr != nil {
			allErrs = append(allErrs, verr)
		}
		isLabelSelectorEmpty = isEmpty
	}

	if (selector.Name != "" && selector.LabelSelector != nil) || (selector.NameRegex != "" && selector.LabelSelector != nil) ||
		(selector.Name != "" && selector.NameRegex != "") {
		verr := field.Duplicate(fld, "either name, name regex or label selector can be used")
		allErrs = append(allErrs, verr)
	}

	if selector.Name == "" && selector.NameRegex == "" && isLabelSelectorEmpty {
		verr := field.Invalid(fld, selector, "neither name, name regex nor label selector is specified")
		allErrs = append(allErrs, verr)
	}

	return allErrs
}

func validateLabelSelector(labelSelector *v1.LabelSelector, fld *field.Path) (empty bool, verr *field.Error) {
	s, err := v1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		isEmpty := len(labelSelector.MatchLabels) == 0 && len(labelSelector.MatchExpressions) == 0
		return isEmpty, field.Invalid(fld, labelSelector, err.Error())
	}

	return s.Empty(), nil
}

func (v *Validator) validateResource(resource testtriggerv1.TestTriggerResource) *field.Error {
	if !utils.In(string(resource), testtrigger.GetSupportedResources()) {
		fld := field.NewPath("spec").Child("resource")
		return field.NotSupported(fld, resource, testtrigger.GetSupportedResources())
	}
	return nil
}

func (v *Validator) validateAction(action testtriggerv1.TestTriggerAction) *field.Error {
	if !utils.In(string(action), testtrigger.GetSupportedActions()) {
		fld := field.NewPath("spec").Child("action")
		return field.NotSupported(fld, action, testtrigger.GetSupportedActions())
	}
	return nil
}

func (v *Validator) validateExecution(execution testtriggerv1.TestTriggerExecution) *field.Error {
	if !utils.In(string(execution), testtrigger.GetSupportedExecutions()) {
		fld := field.NewPath("spec").Child("execution")
		return field.NotSupported(fld, execution, testtrigger.GetSupportedExecutions())
	}
	return nil
}

func (v *Validator) validateConcurrencyPolicy(concurrencyPolicy testtriggerv1.TestTriggerConcurrencyPolicy) *field.Error {
	if concurrencyPolicy != "" && !utils.In(string(concurrencyPolicy), testtrigger.GetSupportedConcurrencyPolicies()) {
		fld := field.NewPath("spec").Child("concurrencyPolicy")
		return field.NotSupported(fld, concurrencyPolicy, testtrigger.GetSupportedConcurrencyPolicies())
	}
	return nil
}

func (v *Validator) validateConditions(conditionSpec *testtriggerv1.TestTriggerConditionSpec) field.ErrorList {
	var allErrs field.ErrorList
	if conditionSpec == nil {
		return allErrs
	}

	if conditionSpec.Timeout < 0 {
		fld := field.NewPath("spec").Child("conditionSpec").Child("timeout")
		verr := field.Invalid(fld, conditionSpec.Timeout, "timeout is negative")
		allErrs = append(allErrs, verr)
	}

	if conditionSpec.Delay < 0 {
		fld := field.NewPath("spec").Child("conditionSpec").Child("delay")
		verr := field.Invalid(fld, conditionSpec.Delay, "delay is negative")
		allErrs = append(allErrs, verr)
	}

	for _, condition := range conditionSpec.Conditions {
		if condition.Type_ == "" {
			fld := field.NewPath("spec").Child("conditionSpec").Child("conditions").Child("condition")
			verr := field.Invalid(fld, condition.Type_, "condition type is not specified")
			allErrs = append(allErrs, verr)
		}

		if condition.Status == nil {
			fld := field.NewPath("spec").Child("conditionSpec").Child("conditions").Child("condition")
			verr := field.Invalid(fld, condition.Status, "condition status is not specified")
			allErrs = append(allErrs, verr)
			continue
		}

		if !utils.In(string(*condition.Status), testtrigger.GetSupportedConditionStatuses()) {
			fld := field.NewPath("spec").Child("conditionSpec").Child("conditions").Child("condition")
			allErrs = append(allErrs, field.NotSupported(fld, string(*condition.Status), testtrigger.GetSupportedConditionStatuses()))
		}
	}

	return allErrs
}

func (v *Validator) validateProbes(probeSpec *testtriggerv1.TestTriggerProbeSpec) field.ErrorList {
	var allErrs field.ErrorList
	if probeSpec == nil {
		return allErrs
	}

	if probeSpec.Timeout < 0 {
		fld := field.NewPath("spec").Child("probeSpec").Child("timeout")
		verr := field.Invalid(fld, probeSpec.Timeout, "timeout is negative")
		allErrs = append(allErrs, verr)
	}

	if probeSpec.Delay < 0 {
		fld := field.NewPath("spec").Child("probeSpec").Child("delay")
		verr := field.Invalid(fld, probeSpec.Delay, "delay is negative")
		allErrs = append(allErrs, verr)
	}

	return allErrs
}

var _ testtriggerv1.TestTriggerValidator = &Validator{}
