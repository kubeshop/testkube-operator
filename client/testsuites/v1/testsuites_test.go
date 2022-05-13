//go:build k8sIntegration

// TODO set-up workflows which can run kubernetes related tests

package v1

import (
	"testing"

	commonv1 "github.com/kubeshop/testkube-operator/apis/common/v1"
	testsuitev1 "github.com/kubeshop/testkube-operator/apis/testsuite/v1"
	kubeclient "github.com/kubeshop/testkube-operator/client"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const testsuiteName = "testsuite-example-with-secrets"

func TestClient_IntegrationWithSecrets(t *testing.T) {
	// given test client and example test
	client, err := kubeclient.GetClient()
	assert.NoError(t, err)

	c := NewClient(client, "testkube")

	tst0, err := c.Create(&testsuitev1.TestSuite{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testsuiteName,
			Namespace: "testkube",
		},
		Spec: testsuitev1.TestSuiteSpec{
			Variables: map[string]testsuitev1.Variable{
				"secretVar1": {
					Type_: commonv1.VariableTypeSecret,
					Name:  "secretVar1",
					Value: "SECR3t",
				},
				"secretVar2": {
					Type_: commonv1.VariableTypeSecret,
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

	secret = tst0.Spec.Variables["secretVar2"]
	secret.Value = "SomeOtherSecretVar"
	tst0.Spec.Variables["secretVar2"] = secret

	tstUpdated, err := c.Update(tst0)
	assert.NoError(t, err)

	// then value should be updated
	tst1, err := c.Get(tst0.Name)
	assert.NoError(t, err)

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
