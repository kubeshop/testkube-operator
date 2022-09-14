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
	"fmt"
	testsv3 "github.com/kubeshop/testkube-operator/apis/tests/v3"
	testsuitev2 "github.com/kubeshop/testkube-operator/apis/testsuite/v2"
	testtriggerv1 "github.com/kubeshop/testkube-operator/apis/testtriggers/v1"
	pkgerrors "github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	executionTest       = "test"
	executionTestsuite  = "testsuite"
	actionRun           = "run"
	resourcePod         = "pod"
	resourceDeployment  = "deployment"
	resourceStatefulSet = "statefulset"
	resourceDaemonSet   = "daemonset"
	resourceService     = "service"
	resourceIngress     = "ingress"
	defaultNamespace    = "testkube"
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

	if err := v.validateExecution(t.Spec.Execution); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateResourceSelector(t.Spec.ResourceSelector); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateTestSelector(ctx, t.Spec.TestSelector, t.Spec.Execution); err != nil {
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

	if err := v.validateExecution(new.Spec.Execution); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateResourceSelector(new.Spec.ResourceSelector); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateTestSelector(ctx, new.Spec.TestSelector, new.Spec.Execution); err != nil {
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

func (v *Validator) validateResourceSelector(resourceSelector testtriggerv1.TestTriggerSelector) *field.Error {
	fld := field.NewPath("spec").Child("testSelector")
	if err := v.validateSelector(fld, resourceSelector); err != nil {
		return err
	}
	return nil
}

func (v *Validator) validateTestSelector(
	ctx context.Context,
	testSelector testtriggerv1.TestTriggerSelector,
	execution string,
) field.ErrorList {
	var allErrs field.ErrorList

	fld := field.NewPath("spec").Child("testSelector")
	if err := v.validateSelector(fld, testSelector); err != nil {
		allErrs = append(allErrs, err)
	}

	if testSelector.Name != "" {
		namespace := testSelector.Namespace
		if namespace == "" {
			namespace = defaultNamespace
		}
		fld = fld.Child("name")
		if err := v.getTestResource(ctx, fld, execution, namespace, testSelector.Name); err != nil {
			allErrs = append(allErrs, err)
		}
	}

	return allErrs
}

func (v *Validator) validateSelector(fld *field.Path, selector testtriggerv1.TestTriggerSelector) *field.Error {
	if selector.Name != "" && len(selector.Labels) > 0 {
		verr := field.Duplicate(fld, "either name or labels selector can be used")
		return verr
	}
	return nil
}

func (v *Validator) getTestResource(
	ctx context.Context,
	fld *field.Path,
	execution, namespace, name string,
) *field.Error {
	switch execution {
	case executionTestsuite:
		var testsuite testsuitev2.TestSuite
		err := v.c.Get(ctx, client.ObjectKey{Name: name, Namespace: namespace}, &testsuite)
		if k8serrors.IsNotFound(err) {
			return field.NotFound(fld, fmt.Sprintf("testsuites.tests.testkube.io/v2 %s/%s", namespace, name))
		} else if err != nil {
			return field.InternalError(
				fld,
				pkgerrors.Errorf("error fetching TestSuite V2 %s/%s: %v", namespace, name, err),
			)
		}
	case executionTest:
		var test testsv3.Test
		err := v.c.Get(ctx, client.ObjectKey{Name: name, Namespace: namespace}, &test)
		if k8serrors.IsNotFound(err) {
			return field.NotFound(fld, fmt.Sprintf("tests.tests.testkube.io/v3 %s/%s", namespace, name))
		} else if err != nil {
			return field.InternalError(
				fld,
				pkgerrors.Errorf("error fetching Test V3 %s/%s: %v", namespace, name, err),
			)
		}
	}
	return nil
}

func (v *Validator) validateResource(resource string) *field.Error {
	allowedResources := []string{
		resourcePod,
		resourceDeployment,
		resourceStatefulSet,
		resourceDaemonSet,
		resourceService,
		resourceIngress,
	}
	if !in(resource, allowedResources) {
		fld := field.NewPath("spec").Child("resource")
		verr := field.NotSupported(
			fld,
			resource,
			[]string{"pod", "deployment", "statefulset", "daemonset", "service", "ingress"},
		)
		return verr
	}
	return nil
}

func in[T comparable](target T, arr []T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func (v *Validator) validateAction(action string) *field.Error {
	if action != actionRun {
		fld := field.NewPath("spec").Child("action")
		verr := field.NotSupported(fld, action, []string{"run"})
		return verr
	}
	return nil
}

func (v *Validator) validateExecution(execution string) *field.Error {
	if execution != executionTest && execution != executionTestsuite {
		fld := field.NewPath("spec").Child("execution")
		verr := field.NotSupported(fld, execution, []string{"test", "testsuite"})
		return verr
	}
	return nil
}

var _ testtriggerv1.TestTriggerValidator = &Validator{}
