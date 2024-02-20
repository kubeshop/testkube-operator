package cronjob

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/controller-runtime/pkg/client"
	kyaml "sigs.k8s.io/kustomize/kyaml/yaml"
	"sigs.k8s.io/kustomize/kyaml/yaml/merge2"
)

const (
	// TestResourceURI is test resource uri for cron job call
	TestResourceURI = "tests"
	// TestSuiteResourceURI is test suite resource uri for cron job call
	TestSuiteResourceURI = "test-suites"
)

// Client data struct for managing running cron jobs
type Client struct {
	client.Client
	serviceName     string
	servicePort     int
	cronJobTemplate string
	registry        string
	argoCDSync      bool
}

type CronJobOptions struct {
	Schedule                  string
	Resource                  string
	Version                   string
	ResourceURI               string
	Data                      string
	Labels                    map[string]string
	CronJobTemplate           string
	CronJobTemplateExtensions string
}

type templateParameters struct {
	Id                        string
	Name                      string
	Namespace                 string
	ServiceName               string
	ServicePort               int
	Schedule                  string
	Resource                  string
	Version                   string
	ResourceURI               string
	CronJobTemplate           string
	CronJobTemplateExtensions string
	Data                      string
	Labels                    map[string]string
	Registry                  string
	ArgoCDSync                bool
	UID                       string
}

// NewClient is a method to create new cron job client
func NewClient(cli client.Client, serviceName string, servicePort int, cronJobTemplate, registry string,
	argoCDSync bool) *Client {
	return &Client{
		Client:          cli,
		serviceName:     serviceName,
		servicePort:     servicePort,
		cronJobTemplate: cronJobTemplate,
		registry:        registry,
		argoCDSync:      argoCDSync,
	}
}

// Get is a method to retrieve an existing cron job
func (c *Client) Get(ctx context.Context, name, namespace string) (*batchv1.CronJob, error) {
	var cronJob batchv1.CronJob
	if err := c.Client.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, &cronJob); err != nil {
		return nil, err
	}

	return &cronJob, nil
}

// Create is a method to create a cron job
func (c *Client) Create(ctx context.Context, id, name, namespace, uid string, options CronJobOptions) error {
	template := c.cronJobTemplate
	if options.CronJobTemplate != "" {
		template = options.CronJobTemplate
	}

	parameters := templateParameters{
		Id:                        id,
		Name:                      name,
		Namespace:                 namespace,
		ServiceName:               c.serviceName,
		ServicePort:               c.servicePort,
		Schedule:                  options.Schedule,
		Resource:                  options.Resource,
		Version:                   options.Version,
		ResourceURI:               options.ResourceURI,
		CronJobTemplate:           template,
		CronJobTemplateExtensions: options.CronJobTemplateExtensions,
		Data:                      options.Data,
		Labels:                    options.Labels,
		Registry:                  c.registry,
		ArgoCDSync:                c.argoCDSync,
		UID:                       uid,
	}

	cronJobSpec, err := NewCronJobSpec(parameters)
	if err != nil {
		return err
	}

	if err := c.Client.Create(ctx, cronJobSpec); err != nil {
		return err
	}

	return nil
}

// Update is a method to update an existing cron job
func (c *Client) Update(ctx context.Context, cronJob *batchv1.CronJob, id, name, namespace, uid string, options CronJobOptions) error {
	template := c.cronJobTemplate
	if options.CronJobTemplate != "" {
		template = options.CronJobTemplate
	}

	parameters := templateParameters{
		Id:                        id,
		Name:                      name,
		Namespace:                 namespace,
		ServiceName:               c.serviceName,
		ServicePort:               c.servicePort,
		Schedule:                  options.Schedule,
		Resource:                  options.Resource,
		Version:                   options.Version,
		ResourceURI:               options.ResourceURI,
		CronJobTemplate:           template,
		CronJobTemplateExtensions: options.CronJobTemplateExtensions,
		Data:                      options.Data,
		Labels:                    options.Labels,
		Registry:                  c.registry,
		ArgoCDSync:                c.argoCDSync,
		UID:                       uid,
	}

	cronJobSpec, err := NewCronJobSpec(parameters)
	if err != nil {
		return err
	}

	cronJob.ObjectMeta = cronJobSpec.ObjectMeta
	cronJob.Spec = cronJobSpec.Spec
	if err := c.Client.Update(ctx, cronJob); err != nil {
		return err
	}

	return nil
}

// Delete is a method to delete a cron job if it exists
func (c *Client) Delete(ctx context.Context, name, namespace string) error {
	var cronJob batchv1.CronJob
	if err := c.Client.Get(context.Background(), types.NamespacedName{
		Name:      name,
		Namespace: namespace}, &cronJob); err != nil {
		if errors.IsNotFound(err) {
			return nil
		}

		return err
	}

	if err := c.Client.Delete(ctx, &cronJob); err != nil {
		return err
	}

	return nil
}

// NewCronJobSpec is a method to return cron job spec
func NewCronJobSpec(parameters templateParameters) (*batchv1.CronJob, error) {
	tmpl, err := template.New("cronjob").Parse(parameters.CronJobTemplate)
	if err != nil {
		return nil, fmt.Errorf("creating cron job spec from options.CronJobTemplate error: %w", err)
	}

	parameters.Data = strings.ReplaceAll(parameters.Data, "'", "''''")
	var buffer bytes.Buffer
	if err = tmpl.ExecuteTemplate(&buffer, "cronjob", parameters); err != nil {
		return nil, fmt.Errorf("executing cron job spec template: %w", err)
	}

	var cronJob batchv1.CronJob
	cronJobSpec := buffer.String()
	if parameters.CronJobTemplateExtensions != "" {
		tmplExt, err := template.New("cronJobExt").Parse(parameters.CronJobTemplateExtensions)
		if err != nil {
			return nil, fmt.Errorf("creating cron job extensions spec from default template error: %w", err)
		}

		var bufferExt bytes.Buffer
		if err = tmplExt.ExecuteTemplate(&bufferExt, "cronJobExt", parameters); err != nil {
			return nil, fmt.Errorf("executing cron job extensions spec default template: %w", err)
		}

		if cronJobSpec, err = merge2.MergeStrings(bufferExt.String(), cronJobSpec, false, kyaml.MergeOptions{}); err != nil {
			return nil, fmt.Errorf("merging cron job spec templates: %w", err)
		}
	}

	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewBufferString(cronJobSpec), len(cronJobSpec))
	if err := decoder.Decode(&cronJob); err != nil {
		return nil, fmt.Errorf("decoding cron job spec error: %w", err)
	}

	for key, value := range parameters.Labels {
		cronJob.Labels[key] = value
	}

	return &cronJob, nil
}

// GetMetadataName returns cron job metadata name
func GetMetadataName(name, resource string) string {
	result := fmt.Sprintf("%s-%s", name, resource)

	if len(result) > 52 {
		return result[:52]
	}

	return result
}
