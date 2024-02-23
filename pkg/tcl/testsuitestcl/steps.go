// Copyright 2024 Testkube.
//
// Licensed as a Testkube Pro file under the Testkube Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/kubeshop/testkube-operator/blob/main/licenses/TCL.txt

package testsuitestcl

import (
	commonv1 "github.com/kubeshop/testkube-operator/api/common/v1"
)

// +kubebuilder:object:generate=true
type Variable commonv1.Variable

// +kubebuilder:object:generate=true
type ArgsModeType commonv1.ArgsModeType

// TestSuiteStepExecutionRequest contains parameters to be used by the executions.
// These fields will be passed to the execution when a Test Suite is queued for execution.
// TestSuiteStepExecutionRequest parameters have the highest priority. They override the
// values coming from Test Suites, Tests, and Test Executions.
// +kubebuilder:object:generate=true
type TestSuiteStepExecutionRequest struct {
	// test execution custom name
	Name string `json:"name,omitempty"`
	// test execution labels
	ExecutionLabels map[string]string `json:"executionLabels,omitempty"`
	// test kubernetes namespace (\"testkube\" when not set)
	Namespace string              `json:"namespace,omitempty"`
	Variables map[string]Variable `json:"variables,omitempty"`
	// additional executor binary arguments
	Args []string `json:"args,omitempty"`
	// usage mode for arguments
	ArgsMode ArgsModeType `json:"argsMode,omitempty"`
	// executor binary command
	Command []string `json:"command,omitempty"`
	// whether to start execution sync or async
	Sync bool `json:"sync,omitempty"`
	// http proxy for executor containers
	HttpProxy string `json:"httpProxy,omitempty"`
	// https proxy for executor containers
	HttpsProxy string `json:"httpsProxy,omitempty"`
	// negative test will fail the execution if it is a success and it will succeed if it is a failure
	NegativeTest bool `json:"negativeTest,omitempty"`
	// job template extensions
	JobTemplate string `json:"jobTemplate,omitempty"`
	// job template extensions reference
	JobTemplateReference string `json:"jobTemplateReference,omitempty"`
	// cron job template extensions
	CronJobTemplate string `json:"cronJobTemplate,omitempty"`
	// cron job template extensions reference
	CronJobTemplateReference string `json:"cronJobTemplateReference,omitempty"`
	// scraper template extensions
	ScraperTemplate string `json:"scraperTemplate,omitempty"`
	// scraper template extensions reference
	ScraperTemplateReference string `json:"scraperTemplateReference,omitempty"`
	// pvc template extensions
	PvcTemplate string `json:"pvcTemplate,omitempty"`
	// pvc template extensions reference
	PvcTemplateReference string                   `json:"pvcTemplateReference,omitempty"`
	RunningContext       *commonv1.RunningContext `json:"runningContext,omitempty"`
}
