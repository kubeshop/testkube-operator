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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
	"testing"
)

func TestValidator_validateAction(t *testing.T) {
	v := NewValidator(buildFakeK8sClient(t))

	t.Run("no error on valid action", func(t *testing.T) {
		err := v.validateAction("run")
		assert.Nil(t, err)
	})

	t.Run("error on invalid action", func(t *testing.T) {
		err := v.validateAction("kill")
		assert.Error(t, err)
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
