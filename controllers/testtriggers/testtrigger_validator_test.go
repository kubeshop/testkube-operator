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
	testtriggerv1 "github.com/kubeshop/testkube-operator/apis/testtriggers/v1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
	"testing"
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
