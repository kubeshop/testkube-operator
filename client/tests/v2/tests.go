package tests

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testsv2 "github.com/kubeshop/testkube-operator/apis/tests/v2"
	"k8s.io/apimachinery/pkg/labels"
)

const testkubeTestSecretLabel = "tests-secrets"

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

// Get returns Test
func (s TestsClient) Get(name string) (*testsv2.Test, error) {
	test := &testsv2.Test{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: name}, test)
	return test, err
}

// Create creates new Test
func (s TestsClient) Create(test *testsv2.Test) (*testsv2.Test, error) {
	err := s.Client.Create(context.Background(), test)

	return test, err
}

// Update updates existing Test
func (s TestsClient) Update(test *testsv2.Test) (*testsv2.Test, error) {
	err := s.Client.Update(context.Background(), test)
	return test, err
}

// Delete deletes existing Test
func (s TestsClient) Delete(name string) error {
	test, err := s.Get(name)
	if err != nil {
		return err
	}

	err = s.Client.Delete(context.Background(), test)
	return err
}

// DeleteAll deletes all Tests
func (s TestsClient) DeleteAll() error {

	u := &unstructured.Unstructured{}
	u.SetKind("Test")
	u.SetAPIVersion("tests.testkube.io/v2")
	err := s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace))
	return err
}

// Delete deletes existing Test
func (s TestsClient) CreateTestSecrets(test *testsv2.Test) error {
	secretName := fmt.Sprintf("%s-vars", test.Name)
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: s.Namespace,
			Labels:    map[string]string{"testkube": testkubeTestSecretLabel},
		},
		Type: v1.SecretTypeOpaque,
	}

	for k := range test.Spec.Variables {
		v := test.Spec.Variables[k]
		if v.Type_ == testsv2.VariableTypeSecret {
			secret.StringData[v.Name] = v.Value
			v.Value = ""
			v.ValueFrom = v1.EnvVarSource{
				SecretKeyRef: &v1.SecretKeySelector{
					Key: v.Name,
					LocalObjectReference: v1.LocalObjectReference{
						Name: secretName,
					},
				},
			}
		}
		test.Spec.Variables[k] = v

		s.Client.Create(context.Background(), secret)
	}
	return nil
}

// NewSpec is a method to return secret spec
func NewSpec(id, namespace string, labels, stringData map[string]string) *v1.Secret {
	configuration := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      id,
			Namespace: namespace,
			Labels:    map[string]string{"testkube": testkubeTestSecretLabel},
		},
		Type:       v1.SecretTypeOpaque,
		StringData: stringData,
	}

	for key, value := range labels {
		configuration.Labels[key] = value
	}

	return configuration
}
