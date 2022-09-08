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
	testsv1 "github.com/kubeshop/testkube-operator/apis/tests/v1"
	testsuitev1 "github.com/kubeshop/testkube-operator/apis/testsuite/v2"
	v12 "github.com/kubeshop/testkube-operator/apis/testtriggers/v1"
	pkgerrors "github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	testTypeTest        = "test"
	testTypeTestsuite   = "testsuite"
	actionRun           = "run"
	resourcePod         = "pod"
	resourceDeployment  = "deployment"
	resourceStatefulSet = "statefulset"
	resourceDaemonSet   = "daemonset"
	resourceService     = "service"
	resourceIngress     = "ingress"
)

type Validator struct {
	c client.Client
}

func (v *Validator) ValidateCreate(ctx context.Context, t *v12.TestTrigger) error {
	var allErrs field.ErrorList

	if err := v.validateResource(t.Spec.Resource); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateAction(t.Spec.Action); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateTestType(t.Spec.TestType); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateResourceSelector(t.Spec.ResourceSelector); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := v.validateTestSelector(ctx, t.Spec.TestSelector, t.Spec.TestType, t.Namespace); err != nil {
		allErrs = append(allErrs, err...)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return k8serrors.NewInvalid(
		schema.GroupKind{
			Group: v12.GroupVersion.Group,
			Kind:  "TestTrigger",
		},
		t.Name,
		allErrs,
	)
}

func (v *Validator) validateResourceSelector(resourceSelector v12.TestTriggerSelector) *field.Error {
	fld := field.NewPath("spec").Child("testSelector")
	if err := v.validateSelector(fld, resourceSelector); err != nil {
		return err
	}
	return nil
}

func (v *Validator) validateTestSelector(
	ctx context.Context,
	testSelector v12.TestTriggerSelector,
	testType, namespace string,
) field.ErrorList {
	var allErrs field.ErrorList

	fld := field.NewPath("spec").Child("testSelector").Child("name")
	if err := v.validateSelector(fld, testSelector); err != nil {
		allErrs = append(allErrs, err)
	}

	if testSelector.Name != "" {
		if err := v.getTestResource(ctx, fld, testType, namespace, testSelector.Name); err != nil {
			allErrs = append(allErrs, err)
		}
	}

	return allErrs
}

func (v *Validator) validateSelector(fld *field.Path, selector v12.TestTriggerSelector) *field.Error {
	if selector.Name != "" && len(selector.Labels) > 0 {
		verr := field.Invalid(fld, "", "either name or labels selector can be used")
		return verr
	}
	return nil
}

func (v *Validator) getTestResource(
	ctx context.Context,
	fld *field.Path,
	testType string,
	namespace, name string,
) *field.Error {
	switch testType {
	case testTypeTestsuite:
		var testsuite testsuitev1.TestSuite
		err := v.c.Get(ctx, client.ObjectKey{Name: name, Namespace: namespace}, &testsuite)
		if k8serrors.IsNotFound(err) {
			return field.Invalid(fld, name, fmt.Sprintf("TestSuite %s/%s not does not exist", namespace, name))
		} else {
			return field.InternalError(fld, pkgerrors.Errorf("error fetching TestSuite %s/%s", namespace, name))
		}
	case testTypeTest:
		var test testsv1.Test
		err := v.c.Get(ctx, client.ObjectKey{Name: name, Namespace: namespace}, &test)
		if k8serrors.IsNotFound(err) {
			return field.Invalid(fld, name, fmt.Sprintf("TestSuite %s/%s not does not exist", namespace, name))
		} else {
			return field.InternalError(fld, pkgerrors.Errorf("error fetching TestSuite %s/%s", namespace, name))
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

func in(target string, arr []string) bool {
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

func (v *Validator) validateTestType(testType string) *field.Error {
	if testType != testTypeTest && testType != testTypeTestsuite {
		fld := field.NewPath("spec").Child("testType")
		verr := field.NotSupported(fld, testType, []string{"test", "testsuite"})
		return verr
	}
	return nil
}
