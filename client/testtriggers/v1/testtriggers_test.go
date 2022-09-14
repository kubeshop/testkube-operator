// TODO set-up workflows which can run kubernetes related tests

package v1

import (
	"context"
	testtriggerv1 "github.com/kubeshop/testkube-operator/apis/testtriggers/v1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
	"testing"
)

func TestTestTriggers(t *testing.T) {
	ctx := context.Background()
	var tClient *TestTriggersClient
	testTestTriggers := []*testtriggerv1.TestTrigger{
		{
			ObjectMeta: v1.ObjectMeta{
				Name:      "test-testtrigger1",
				Namespace: "test-ns",
			},
			Spec: testtriggerv1.TestTriggerSpec{
				Resource:         "test-resource1",
				ResourceSelector: testtriggerv1.TestTriggerSelector{Name: "test-pod1"},
				Event:            "",
				Action:           "run",
				Execution:        "test",
				TestSelector:     testtriggerv1.TestTriggerSelector{Name: "test-test1"},
			},
			Status: testtriggerv1.TestTriggerStatus{},
		},
	}

	t.Run("NewTestTriggerClient", func(t *testing.T) {
		clientBuilder := fake.NewClientBuilder()

		groupVersion := schema.GroupVersion{Group: "tests.testkube.io", Version: "v1"}
		schemaBuilder := scheme.Builder{GroupVersion: groupVersion}
		schemaBuilder.Register(&testtriggerv1.TestTriggerList{})
		schemaBuilder.Register(&testtriggerv1.TestTrigger{})

		schema, err := schemaBuilder.Build()
		assert.NoError(t, err)
		assert.NotEmpty(t, schema)
		clientBuilder.WithScheme(schema)

		kClient := clientBuilder.Build()
		testNamespace := "test-ns"
		tClient = NewClient(kClient, testNamespace)
		assert.NotEmpty(t, tClient)
		assert.Equal(t, testNamespace, tClient.Namespace)
	})
	t.Run("TestCreate", func(t *testing.T) {
		t.Run("Create new TestTrigger", func(t *testing.T) {
			for _, tt := range testTestTriggers {
				created, err := tClient.Create(ctx, tt)
				assert.NoError(t, err)
				assert.Equal(t, tt.Name, created.Name)

				res, err := tClient.Get(ctx, tt.Name)
				assert.NoError(t, err)
				assert.Equal(t, tt.Name, res.Name)
			}
		})
	})
	t.Run("TestList", func(t *testing.T) {
		t.Run("List without selector", func(t *testing.T) {
			l, err := tClient.List(ctx, "")
			assert.NoError(t, err)
			assert.Equal(t, len(testTestTriggers), len(l.Items))
		})
	})
	t.Run("TestGet", func(t *testing.T) {
		t.Run("Get TestTrigger with empty name", func(t *testing.T) {
			t.Parallel()
			_, err := tClient.Get(ctx, "")
			assert.Error(t, err)
		})

		t.Run("Get TestTrigger with non existent name", func(t *testing.T) {
			t.Parallel()
			_, err := tClient.Get(ctx, "no-testtrigger")
			assert.Error(t, err)
		})

		t.Run("Get existing TestTrigger", func(t *testing.T) {
			res, err := tClient.Get(ctx, testTestTriggers[0].Name)
			assert.NoError(t, err)
			assert.Equal(t, testTestTriggers[0].Name, res.Name)
		})
	})
	t.Run("TestDelete", func(t *testing.T) {
		t.Run("Delete items", func(t *testing.T) {
			for _, trigger := range testTestTriggers {
				tt, err := tClient.Get(ctx, trigger.Name)
				assert.NoError(t, err)
				assert.Equal(t, tt.Name, trigger.Name)

				err = tClient.Delete(ctx, trigger.Name)
				assert.NoError(t, err)

				_, err = tClient.Get(ctx, trigger.Name)
				assert.Error(t, err)
			}
		})

		t.Run("Delete non-existent item", func(t *testing.T) {
			_, err := tClient.Get(ctx, "no-testtrigger")
			assert.Error(t, err)

			err = tClient.Delete(ctx, "no-testtrigger")
			assert.Error(t, err)
		})
	})
}
