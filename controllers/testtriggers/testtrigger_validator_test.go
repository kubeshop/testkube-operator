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
	"testing"

	testtriggerv1 "github.com/kubeshop/testkube-operator/apis/testtriggers/v1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

func TestValidator_validateAction(t *testing.T) {
	t.Parallel()

	v := NewValidator(buildFakeK8sClient(t))

	t.Run("no error for valid action", func(t *testing.T) {
		t.Parallel()

		err := v.validateAction("run")
		assert.Nil(t, err)
	})

	t.Run("error for invalid action", func(t *testing.T) {
		t.Parallel()

		err := v.validateAction("kill")
		assert.ErrorContains(t, err, "spec.action: Unsupported value: \"kill\"")
	})
}

func TestValidator_validateExecution(t *testing.T) {
	t.Parallel()

	v := NewValidator(buildFakeK8sClient(t))

	t.Run("no error for supported execution", func(t *testing.T) {
		t.Parallel()

		err := v.validateExecution("test")
		assert.Nil(t, err)
	})

	t.Run("error for unsupported execution", func(t *testing.T) {
		t.Parallel()

		err := v.validateExecution("testspec")
		assert.ErrorContains(t, err, "spec.execution: Unsupported value: \"testspec\"")
	})
}

func TestValidator_validateResource(t *testing.T) {
	t.Parallel()

	v := NewValidator(buildFakeK8sClient(t))

	t.Run("no error for supported resource", func(t *testing.T) {
		t.Parallel()

		err := v.validateResource("pod")
		assert.Nil(t, err)
	})

	t.Run("error for unsupported resource", func(t *testing.T) {
		t.Parallel()

		err := v.validateResource("replicaset")
		assert.ErrorContains(t, err, "spec.resource: Unsupported value: \"replicaset\"")
	})
}

func TestValidator_validateSelector(t *testing.T) {
	t.Parallel()

	v := NewValidator(buildFakeK8sClient(t))

	t.Run("no error if only name selector is specified", func(t *testing.T) {
		t.Parallel()

		fld := field.NewPath("spec").Child("testSelector")
		selector := testtriggerv1.TestTriggerSelector{Name: "test"}
		verrs := v.validateSelector(fld, selector)
		assert.Empty(t, verrs)
	})

	t.Run("no error if valid label selector is specified", func(t *testing.T) {
		t.Parallel()

		fld := field.NewPath("spec").Child("testSelector")
		selector := testtriggerv1.TestTriggerSelector{
			LabelSelector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"tier": "test"},
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{
						Key:      "app",
						Operator: metav1.LabelSelectorOpExists,
					},
				},
			},
		}
		verrs := v.validateSelector(fld, selector)
		assert.Empty(t, verrs)
	})

	t.Run("error when neither name nor label selector are specified", func(t *testing.T) {
		t.Parallel()

		fld := field.NewPath("spec").Child("testSelector")
		selector := testtriggerv1.TestTriggerSelector{LabelSelector: &metav1.LabelSelector{}}
		verrs := v.validateSelector(fld, selector)
		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "neither name nor label selector is specified")
	})

	t.Run("error when invalid labels are specified", func(t *testing.T) {
		t.Parallel()

		fld := field.NewPath("spec").Child("testSelector")
		selector := testtriggerv1.TestTriggerSelector{
			LabelSelector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"invalid=label": "value"},
			},
		}
		verrs := v.validateSelector(fld, selector)
		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "Invalid value: \"invalid=label\"")
	})

	t.Run("error when invalid expressions are specified", func(t *testing.T) {
		t.Parallel()

		fld := field.NewPath("spec").Child("testSelector")
		selector := testtriggerv1.TestTriggerSelector{
			LabelSelector: &metav1.LabelSelector{
				MatchExpressions: []metav1.LabelSelectorRequirement{{
					Key:      "test",
					Operator: "invalid",
					Values:   []string{"true"},
				}},
			},
		}
		verrs := v.validateSelector(fld, selector)
		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "\"invalid\" is not a valid pod selector operator")
	})
}

func buildFakeK8sClient(t *testing.T) client.Client {
	t.Helper()

	clientBuilder := fake.NewClientBuilder()

	groupVersion := schema.GroupVersion{Group: "tests.testkube.io", Version: "v1"}
	schemaBuilder := scheme.Builder{GroupVersion: groupVersion}
	schemaBuilder.Register(&testtriggerv1.TestTriggerList{})
	schemaBuilder.Register(&testtriggerv1.TestTrigger{})

	schema, err := schemaBuilder.Build()
	if err != nil {
		assert.Fail(t, "error building schema for TestTrigger")
	}
	assert.NotEmpty(t, schema)
	clientBuilder.WithScheme(schema)

	kClient := clientBuilder.Build()
	return kClient
}

func TestValidator_validateConditions(t *testing.T) {
	t.Parallel()

	v := NewValidator(buildFakeK8sClient(t))

	t.Run("no error for valid condition spec", func(t *testing.T) {
		t.Parallel()

		status := testtriggerv1.TRUE_TestTriggerConditionStatuses
		verrs := v.validateConditions(&testtriggerv1.TestTriggerConditionSpec{
			Timeout: 100,
			Conditions: []testtriggerv1.TestTriggerCondition{
				{Status: &status, Type_: "Progressing"},
			}})

		assert.Nil(t, verrs)
	})

	t.Run("error for invalid timeout", func(t *testing.T) {
		t.Parallel()

		verrs := v.validateConditions(&testtriggerv1.TestTriggerConditionSpec{
			Timeout: -100,
		})

		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "spec.conditionSpec.timeout: Invalid value: -100: timeout is negative")
	})

	t.Run("error for invalid delay", func(t *testing.T) {
		t.Parallel()

		verrs := v.validateConditions(&testtriggerv1.TestTriggerConditionSpec{
			Delay: -1,
		})

		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "spec.conditionSpec.delay: Invalid value: -1: delay is negative")
	})

	t.Run("error for invalid condition type", func(t *testing.T) {
		t.Parallel()

		status := testtriggerv1.TRUE_TestTriggerConditionStatuses
		verrs := v.validateConditions(&testtriggerv1.TestTriggerConditionSpec{
			Conditions: []testtriggerv1.TestTriggerCondition{
				{Status: &status},
			}})

		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "spec.conditionSpec.conditions.condition: Invalid value: \"\": condition type is not specified")
	})

	t.Run("error for invalid condition status", func(t *testing.T) {
		t.Parallel()

		verrs := v.validateConditions(&testtriggerv1.TestTriggerConditionSpec{
			Conditions: []testtriggerv1.TestTriggerCondition{
				{Type_: "Progressing"},
			}})

		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "spec.conditionSpec.conditions.condition: Invalid value: \"null\": condition status is not specified")
	})

	t.Run("error for unsupported condition status", func(t *testing.T) {
		t.Parallel()

		status := testtriggerv1.TestTriggerConditionStatuses("")
		verrs := v.validateConditions(&testtriggerv1.TestTriggerConditionSpec{
			Conditions: []testtriggerv1.TestTriggerCondition{
				{Status: &status, Type_: "Progressing"},
			}})

		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "spec.conditionSpec.conditions.condition: Unsupported value: \"\": supported values: \"True\", \"False\", \"Unknown\"")
	})
}

func TestValidator_validateProbes(t *testing.T) {
	t.Parallel()

	v := NewValidator(buildFakeK8sClient(t))

	t.Run("no error for valid probe spec", func(t *testing.T) {
		t.Parallel()

		verrs := v.validateProbes(&testtriggerv1.TestTriggerProbeSpec{
			Timeout: 50,
			Probes: []testtriggerv1.TestTriggerProbe{
				{
					Host: "testkube-api-server",
					Path: "/health",
					Port: 8088,
				},
			}})

		assert.Nil(t, verrs)
	})

	t.Run("error for invalid timeout", func(t *testing.T) {
		t.Parallel()

		verrs := v.validateProbes(&testtriggerv1.TestTriggerProbeSpec{
			Timeout: -100,
		})

		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "spec.probeSpec.timeout: Invalid value: -100: timeout is negative")
	})

	t.Run("error for invalid delay", func(t *testing.T) {
		t.Parallel()

		verrs := v.validateProbes(&testtriggerv1.TestTriggerProbeSpec{
			Delay: -1,
		})

		assert.Len(t, verrs, 1)
		assert.ErrorContains(t, verrs[0], "spec.probeSpec.delay: Invalid value: -1: delay is negative")
	})
}
