package testexecutions

import (
	"testing"

	testexecutionv1 "github.com/kubeshop/testkube-operator/apis/testexecution/v1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

func TestTestExecutions(t *testing.T) {
	var teClient *TestExecutionsClient
	testTestExecutions := []*testexecutionv1.TestExecution{
		{
			ObjectMeta: v1.ObjectMeta{
				Name:      "test-testexecution1",
				Namespace: "test-ns",
			},
			Spec:   testexecutionv1.TestExecutionSpec{},
			Status: testexecutionv1.TestExecutionStatus{},
		},
		{
			ObjectMeta: v1.ObjectMeta{
				Name:      "test-testexecution2",
				Namespace: "test-ns",
			},
			Spec:   testexecutionv1.TestExecutionSpec{},
			Status: testexecutionv1.TestExecutionStatus{},
		},
		{
			ObjectMeta: v1.ObjectMeta{
				Name:      "test-testexecution3",
				Namespace: "test-ns",
			},
			Spec:   testexecutionv1.TestExecutionSpec{},
			Status: testexecutionv1.TestExecutionStatus{},
		},
	}

	t.Run("NewTestExecutionsClient", func(t *testing.T) {
		clientBuilder := fake.NewClientBuilder()

		groupVersion := schema.GroupVersion{Group: "tests.testkube.io", Version: "v1"}
		schemaBuilder := scheme.Builder{GroupVersion: groupVersion}
		schemaBuilder.Register(&testexecutionv1.TestExecutionList{})
		schemaBuilder.Register(&testexecutionv1.TestExecution{})

		schema, err := schemaBuilder.Build()
		assert.NoError(t, err)
		assert.NotEmpty(t, schema)
		clientBuilder.WithScheme(schema)

		kClient := clientBuilder.Build()
		testNamespace := "test-ns"
		teClient = NewClient(kClient, testNamespace)
		assert.NotEmpty(t, teClient)
		assert.Equal(t, testNamespace, teClient.namespace)
	})
	t.Run("TestExecutionCreate", func(t *testing.T) {
		t.Run("Create new testexecutions", func(t *testing.T) {
			for _, w := range testTestExecutions {
				created, err := teClient.Create(w)
				assert.NoError(t, err)
				assert.Equal(t, w.Name, created.Name)

				res, err := teClient.Get(w.ObjectMeta.Name)
				assert.NoError(t, err)
				assert.Equal(t, w.Name, res.Name)
			}
		})
	})
	t.Run("TestExecutionUpdate", func(t *testing.T) {
		t.Run("Update new testexecutions", func(t *testing.T) {
			for _, w := range testTestExecutions {
				res, err := teClient.Get(w.ObjectMeta.Name)
				assert.NoError(t, err)
				assert.Equal(t, w.Name, res.Name)

				updated, err := teClient.Update(w)
				assert.NoError(t, err)
				assert.Equal(t, w.Name, updated.Name)
			}
		})
	})
	t.Run("TestExecutionGet", func(t *testing.T) {
		t.Run("Get testexecution with empty name", func(t *testing.T) {
			t.Parallel()
			_, err := teClient.Get("")
			assert.Error(t, err)
		})

		t.Run("Get testexecution with non existent name", func(t *testing.T) {
			t.Parallel()
			_, err := teClient.Get("no-testexecution")
			assert.Error(t, err)
		})

		t.Run("Get existing testexecution", func(t *testing.T) {
			res, err := teClient.Get(testTestExecutions[0].Name)
			assert.NoError(t, err)
			assert.Equal(t, testTestExecutions[0].Name, res.Name)
		})
	})
	t.Run("TestExecutionDelete", func(t *testing.T) {
		t.Run("Delete items", func(t *testing.T) {
			for _, testexecution := range testTestExecutions {
				w, err := teClient.Get(testexecution.Name)
				assert.NoError(t, err)
				assert.Equal(t, w.Name, testexecution.Name)

				err = teClient.Delete(testexecution.Name)
				assert.NoError(t, err)

				_, err = teClient.Get(testexecution.Name)
				assert.Error(t, err)
			}
		})

		t.Run("Delete non-existent item", func(t *testing.T) {
			_, err := teClient.Get("no-testexecution")
			assert.Error(t, err)

			err = teClient.Delete("no-testexecution")
			assert.Error(t, err)
		})
	})
}
