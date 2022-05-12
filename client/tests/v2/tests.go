package tests

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testsv2 "github.com/kubeshop/testkube-operator/apis/tests/v2"
	"k8s.io/apimachinery/pkg/labels"
)

const testkubeTestSecretLabel = "tests-secrets"

var testSecretDefaultLabels = map[string]string{
	"testkube":           testkubeTestSecretLabel,
	"testkubeSecretType": "variables",
}

// NewClent creates new Test client
func NewClient(client client.Client, namespace string) *TestsClient {
	return &TestsClient{
		Client:    client,
		Namespace: namespace,
	}
}

// TestClient implements methods to work with Test
type TestsClient struct {
	Client    client.Client
	Namespace string
}

// List lists Tests
func (s TestsClient) List(selector string) (*testsv2.TestList, error) {
	list := &testsv2.TestList{}
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

// ListLabels labels for Tests
func (s TestsClient) ListLabels() (map[string][]string, error) {
	labels := map[string][]string{}
	list := &testsv2.TestList{}
	if err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: s.Namespace}); err != nil {
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

// Get returns Test, loads and decodes secrets data
func (s TestsClient) Get(name string) (*testsv2.Test, error) {
	test := &testsv2.Test{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: name}, test)
	if err != nil {
		return nil, err
	}

	secret, err := s.LoadTestVariablesSecret(test)
	if err != nil {
		return nil, err
	}

	secretToTestVars(secret, test)

	return test, nil
}

// Create creates new Test and coupled resources
func (s TestsClient) Create(test *testsv2.Test) (*testsv2.Test, error) {
	err := s.CreateTestSecrets(test)
	if err != nil {
		return nil, err
	}
	err = s.Client.Create(context.Background(), test)
	return test, err
}

// Update updates existing Test and coupled resources
func (s TestsClient) Update(test *testsv2.Test) (*testsv2.Test, error) {
	err := s.UpdateTestSecrets(test)
	if err != nil {
		return nil, err
	}
	err = s.Client.Update(context.Background(), test)
	return test, err
}

// Delete deletes existing Test and coupled resources (secrets)
func (s TestsClient) Delete(name string) error {
	test, err := s.Get(name)
	if err != nil {
		return err
	}

	secret, err := s.LoadTestVariablesSecret(test)
	if err != nil {
		return err
	}

	err = s.Client.Delete(context.Background(), test)
	if err != nil {
		return err
	}

	return s.Client.Delete(context.Background(), secret)
}

// DeleteAll deletes all Tests
func (s TestsClient) DeleteAll() error {

	u := &unstructured.Unstructured{}
	u.SetKind("Test")
	u.SetAPIVersion("tests.testkube.io/v2")
	err := s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace))
	if err != nil {
		return err
	}

	u = &unstructured.Unstructured{}
	u.SetKind("Secret")
	u.SetAPIVersion("v1")
	u.SetLabels(testSecretDefaultLabels)

	return s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace))
}

// CreateTestSecrets creates corresponding test vars secrets
func (s TestsClient) CreateTestSecrets(test *testsv2.Test) error {
	secretName := secretName(test.Name)
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: s.Namespace,
			Labels:    testSecretDefaultLabels,
		},
		Type:       corev1.SecretTypeOpaque,
		StringData: map[string]string{},
	}

	testVarsToSecret(test, secret)

	if len(secret.StringData) > 0 {
		err := s.Client.Create(context.Background(), secret)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TestsClient) UpdateTestSecrets(test *testsv2.Test) error {
	secret, err := s.LoadTestVariablesSecret(test)
	if err != nil {
		return err
	}

	testVarsToSecret(test, secret)

	if len(secret.StringData) > 0 {
		err := s.Client.Update(context.Background(), secret)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TestsClient) LoadTestVariablesSecret(test *testsv2.Test) (*corev1.Secret, error) {
	secret := &corev1.Secret{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: secretName(test.Name)}, secret)
	return secret, err
}

// testVarsToSecret loads secrets data passed into Test CRD and remove plain text data
func testVarsToSecret(test *testsv2.Test, secret *corev1.Secret) {
	if secret.StringData == nil {
		secret.StringData = map[string]string{}
	}
	for k := range test.Spec.Variables {
		v := test.Spec.Variables[k]
		if v.Type_ == testsv2.VariableTypeSecret {
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
		test.Spec.Variables[k] = v
	}
}

// secretToTestVars loads secrets data passed into Test CRD and remove plain text data
func secretToTestVars(secret *corev1.Secret, test *testsv2.Test) {

	if secret.Data == nil {
		return
	}

	for k, v := range secret.Data {
		if variable, ok := test.Spec.Variables[k]; ok {
			variable.Value = string(v)
			test.Spec.Variables[k] = variable
		}
	}
}

func secretName(testName string) string {
	return fmt.Sprintf("%s-testvars", testName)
}

// DeleteByLabels deletes tests by labels
func (s TestsClient) DeleteByLabels(selector string) error {
	reqs, err := labels.ParseToRequirements(selector)
	if err != nil {
		return err
	}

	u := &unstructured.Unstructured{}
	u.SetKind("Test")
	u.SetAPIVersion("tests.testkube.io/v2")
	err = s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace),
		client.MatchingLabelsSelector{Selector: labels.NewSelector().Add(reqs...)})
	return err
}
