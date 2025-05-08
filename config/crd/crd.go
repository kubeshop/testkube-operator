package crd

import "embed"

//go:embed bases
var SF embed.FS

// Schema is crd schema type
type Schema string

const (
	// SchemaTestWorkflow is test workflow crd schema
	SchemaTestWorkflow Schema = "testworkflows.testkube.io_testworkflows"
	// SchemaTestWorkflowTemplate is test workflow template crd schema
	SchemaTestWorkflowTemplate Schema = "testworkflows.testkube.io_testworkflowtemplates"
)
