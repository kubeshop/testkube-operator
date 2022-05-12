package v1

import (
	"context"
	"fmt"

	testsuitev1 "github.com/kubeshop/testkube-operator/apis/testsuite/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const testkubeTestsuiteSecretLabel = "testsuites-secrets"

var testsuiteSecretDefaultLabels = map[string]string{
	"testkube":           testkubeTestsuiteSecretLabel,
	"testkubeSecretType": "variables",
}

// NewClient creates new TestSuite client
func NewClient(client client.Client, namespace string) *TestSuitesClient {
	return &TestSuitesClient{
		Client:    client,
		Namespace: namespace,
	}
}

// TestSuitesClient implements methods to work with TestSuites
type TestSuitesClient struct {
	Client    client.Client
	Namespace string
}

// List lists TestSuites
func (s TestSuitesClient) List(selector string) (*testsuitev1.TestSuiteList, error) {
	list := &testsuitev1.TestSuiteList{}
	reqs, err := labels.ParseToRequirements(selector)
	if err != nil {
		return list, err
	}

	options := &client.ListOptions{
		Namespace:     s.Namespace,
		LabelSelector: labels.NewSelector().Add(reqs...),
	}

	if err = s.Client.List(context.Background(), list, options); err != nil {
		return list, err
	}

	return list, nil
}

// ListLabelslists labels for TestSuites
func (s TestSuitesClient) ListLabels() (map[string][]string, error) {
	labels := map[string][]string{}
	list := &testsuitev1.TestSuiteList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: s.Namespace})
	if err != nil {
		return labels, err
	}

	for _, test := range list.Items {
		for key, value := range test.Labels {
			if values, ok := labels[key]; !ok {
				labels[key] = []string{value}
			} else {
				for _, v := range values {
					if v == value {
						continue
					}
				}
				labels[key] = append(labels[key], value)
			}
		}
	}

	return labels, nil
}

// Get returns TestSuite
func (s TestSuitesClient) Get(name string) (*testsuitev1.TestSuite, error) {
	testsuite := &testsuitev1.TestSuite{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: name}, testsuite)
	if err != nil {
		return nil, err
	}

	secret, err := s.LoadTestVariablesSecret(testsuite)
	if err != nil {
		return nil, err
	}

	secretToTestVars(secret, testsuite)

	return testsuite, err
}

// Create creates new TestSuite
func (s TestSuitesClient) Create(testsuite *testsuitev1.TestSuite) (*testsuitev1.TestSuite, error) {
	err := s.CreateTestsuiteSecrets(testsuite)
	if err != nil {
		return nil, err
	}

	err = s.Client.Create(context.Background(), testsuite)
	return testsuite, err
}

// Update updates existing TestSuite
func (s TestSuitesClient) Update(testsuite *testsuitev1.TestSuite) (*testsuitev1.TestSuite, error) {
	err := s.UpdateTestsuiteSecrets(testsuite)
	if err != nil {
		return nil, err
	}

	err = s.Client.Update(context.Background(), testsuite)
	return testsuite, err
}

// Delete deletes existing TestSuite
func (s TestSuitesClient) Delete(name string) error {
	testsuite, err := s.Get(name)
	if err != nil {
		return err
	}

	secret, err := s.LoadTestVariablesSecret(testsuite)
	if err != nil {
		return err
	}

	err = s.Client.Delete(context.Background(), testsuite)
	if err != nil {
		return err
	}

	return s.Client.Delete(context.Background(), secret)
}

// DeleteAll delete all TestSuites
func (s TestSuitesClient) DeleteAll() error {
	u := &unstructured.Unstructured{}
	u.SetKind("TestSuite")
	u.SetAPIVersion("tests.testkube.io/v1")
	err := s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace))
	if err != nil {
		return err
	}

	u = &unstructured.Unstructured{}
	u.SetKind("Secret")
	u.SetAPIVersion("v1")
	u.SetLabels(testsuiteSecretDefaultLabels)

	return s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace))
}

// CreateTestSecrets creates corresponding test vars secrets
func (s TestSuitesClient) CreateTestsuiteSecrets(testsuite *testsuitev1.TestSuite) error {
	secretName := secretName(testsuite.Name)
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: s.Namespace,
			Labels:    testsuiteSecretDefaultLabels,
		},
		Type:       corev1.SecretTypeOpaque,
		StringData: map[string]string{},
	}

	testVarsToSecret(testsuite, secret)

	if len(secret.StringData) > 0 {
		err := s.Client.Create(context.Background(), secret)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TestSuitesClient) UpdateTestsuiteSecrets(testsuite *testsuitev1.TestSuite) error {
	secret, err := s.LoadTestVariablesSecret(testsuite)
	if err != nil {
		return err
	}

	testVarsToSecret(testsuite, secret)

	if len(secret.StringData) > 0 {
		err := s.Client.Update(context.Background(), secret)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TestSuitesClient) LoadTestVariablesSecret(testsuite *testsuitev1.TestSuite) (*corev1.Secret, error) {
	secret := &corev1.Secret{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: secretName(testsuite.Name)}, secret)
	return secret, err
}

// testVarsToSecret loads secrets data passed into Test CRD and remove plain text data
func testVarsToSecret(testsuite *testsuitev1.TestSuite, secret *corev1.Secret) {
	if secret.StringData == nil {
		secret.StringData = map[string]string{}
	}
	for k := range testsuite.Spec.Variables {
		v := testsuite.Spec.Variables[k]
		if v.Type_ == testsuitev1.VariableTypeSecret {
			secret.StringData[v.Name] = v.Value
			// clear passed test variable secret value and save as reference o secret
			v.Value = ""
			v.ValueFrom = corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: v.Name,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secret.Name,
					},
				},
			}
		}
		testsuite.Spec.Variables[k] = v
	}
}

// secretToTestVars loads secrets data passed into Test CRD and remove plain text data
func secretToTestVars(secret *corev1.Secret, testsuite *testsuitev1.TestSuite) {

	if secret.Data == nil {
		return
	}

	for k, v := range secret.Data {
		if variable, ok := testsuite.Spec.Variables[k]; ok {
			variable.Value = string(v)
			testsuite.Spec.Variables[k] = variable
		}
	}
}

func secretName(testsuiteName string) string {
	return fmt.Sprintf("%s-testsuitevars", testsuiteName)
}
