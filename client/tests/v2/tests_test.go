//go:build kubernetesIntegrationTest

// TODO set-up workflows which can run kubernetes related tests

package tests

import (
	"fmt"
	"testing"

	testsv2 "github.com/kubeshop/testkube-operator/apis/tests/v2"
	kubeclient "github.com/kubeshop/testkube-operator/client"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestClient_IntegrationWithSecrets(t *testing.T) {
	// given test client and example test
	client, err := kubeclient.GetClient()
	assert.NoError(t, err)

	c := NewClient(client, "testkube")
	tst0, err := c.Create(&testsv2.Test{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-example-with-secrets",
			Namespace: "testkube",
		},
		Spec: testsv2.TestSpec{
			Type_: "postman/collection",
			Content: &testsv2.TestContent{
				Data: "{}",
			},
			Variables: map[string]testsv2.Variable{
				"secretVar1": {
					Type_: testsv2.VariableTypeSecret,
					Name:  "secretVar1",
					Value: "SECR3t",
				},
				"secretVar2": {
					Type_: testsv2.VariableTypeSecret,
					Name:  "secretVar2",
					Value: "SomeOtherSecretVar",
				},
			},
		},
	})

	assert.NoError(t, err)

	// when update test secret variable
	secret := tst0.Spec.Variables["secretVar1"]
	secret.Value = "UpdatedSecretValue"
	tst0.Spec.Variables["secretVar1"] = secret
	tstUpdated, err := c.Update(tst0)
	assert.NoError(t, err)

	// then value should be updated
	tst1, err := c.Get(tst0.Name)
	assert.NoError(t, err)
	fmt.Printf("%+v\n", tst1.Spec.Variables)

	assert.Equal(t, "UpdatedSecretValue", tst1.Spec.Variables["secretVar1"].Value)
	assert.Equal(t, "SomeOtherSecretVar", tst1.Spec.Variables["secretVar2"].Value)

	// when test is deleted
	err = c.Delete(tstUpdated.Name)
	assert.NoError(t, err)

	// then there should be no test anymore
	tst2, err := c.Get(tst0.Name)
	assert.Nil(t, tst2)
	assert.Error(t, err)
}
