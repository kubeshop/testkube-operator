package events

import (
	"strings"

	"github.com/google/uuid"
	testexecutionv1 "github.com/kubeshop/testkube-operator/api/testexecution/v1"
	testsuiteexecutionv1 "github.com/kubeshop/testkube-operator/api/testsuiteexecution/v1"

	"k8s.io/apimachinery/pkg/labels"
)

func NewEvent(t *EventType, resource *EventResource, id string) Event {
	return Event{
		Id:         uuid.NewString(),
		ResourceId: id,
		Resource:   resource,
		Type_:      t,
	}
}

func NewEventStartTest(execution *testexecutionv1.TestExecution) Event {
	return Event{
		Id:            uuid.NewString(),
		Type_:         EventStartTest,
		TestExecution: execution,
	}
}

func NewEventEndTestSuccess(execution *testexecutionv1.TestExecution) Event {
	return Event{
		Id:            uuid.NewString(),
		Type_:         EventEndTestSuccess,
		TestExecution: execution,
	}
}

func NewEventEndTestFailed(execution *testexecutionv1.TestExecution) Event {
	return Event{
		Id:            uuid.NewString(),
		Type_:         EventEndTestFailed,
		TestExecution: execution,
	}
}

func NewEventEndTestAborted(execution *testexecutionv1.TestExecution) Event {
	return Event{
		Id:            uuid.NewString(),
		Type_:         EventEndTestAborted,
		TestExecution: execution,
	}
}

func NewEventEndTestTimeout(execution *testexecutionv1.TestExecution) Event {
	return Event{
		Id:            uuid.NewString(),
		Type_:         EventEndTestTimeout,
		TestExecution: execution,
	}
}

func NewEventStartTestSuite(execution *testsuiteexecutionv1.TestSuiteExecution) Event {
	return Event{
		Id:                 uuid.NewString(),
		Type_:              EventStartTestSuite,
		TestSuiteExecution: execution,
	}
}

func NewEventEndTestSuiteSuccess(execution *testsuiteexecutionv1.TestSuiteExecution) Event {
	return Event{
		Id:                 uuid.NewString(),
		Type_:              EventEndTestSuiteSuccess,
		TestSuiteExecution: execution,
	}
}

func NewEventEndTestSuiteFailed(execution *testsuiteexecutionv1.TestSuiteExecution) Event {
	return Event{
		Id:                 uuid.NewString(),
		Type_:              EventEndTestSuiteFailed,
		TestSuiteExecution: execution,
	}
}

func NewEventEndTestSuiteAborted(execution *testsuiteexecutionv1.TestSuiteExecution) Event {
	return Event{
		Id:                 uuid.NewString(),
		Type_:              EventEndTestSuiteAborted,
		TestSuiteExecution: execution,
	}
}

func NewEventEndTestSuiteTimeout(execution *testsuiteexecutionv1.TestSuiteExecution) Event {
	return Event{
		Id:                 uuid.NewString(),
		Type_:              EventEndTestSuiteTimeout,
		TestSuiteExecution: execution,
	}
}

func (e Event) Type() EventType {
	if e.Type_ != nil {
		return *e.Type_
	}
	return EventType("")
}

func (e Event) IsSuccess() bool {
	return strings.Contains(e.Type().String(), "success")
}

func (e Event) Log() []any {
	var name, eventType, labelsStr string
	var labels map[string]string

	if e.TestSuiteExecution != nil {
		name = e.TestSuiteExecution.Name
		labels = e.TestSuiteExecution.Labels
	} else if e.TestExecution != nil {
		name = e.TestExecution.Name
		labels = e.TestExecution.Labels
	}

	if e.Type_ != nil {
		eventType = e.Type_.String()
	}

	for k, v := range labels {
		labelsStr += k + "=" + v + " "
	}

	resource := ""
	if e.Resource != nil {
		resource = string(*e.Resource)
	}

	return []any{
		"id", e.Id,
		"type", eventType,
		"resource", resource,
		"resourceId", e.ResourceId,
		"executionName", name,
		"labels", labelsStr,
		"topic", e.Topic(),
	}
}

func (e Event) Valid(selector string, types []EventType) (valid bool) {
	var executionLabels map[string]string

	// load labels from event test execution or test-suite execution
	if e.TestSuiteExecution != nil {
		executionLabels = e.TestSuiteExecution.Labels
	} else if e.TestExecution != nil {
		executionLabels = e.TestExecution.Labels
	}

	typesMatch := false
	for _, t := range types {
		if t == e.Type() {
			typesMatch = true
			break
		}
	}

	if !typesMatch {
		return false
	}

	valid = selector == ""
	if !valid {
		selector, err := labels.Parse(selector)
		if err != nil {
			return false
		}

		valid = selector.Matches(labels.Set(executionLabels))
	}

	return
}

// Topic returns topic for event based on resource and resource id
// or fallback to global "events" topic
func (e Event) Topic() string {
	if e.Resource == nil {
		return "events.all"
	}

	if e.ResourceId == "" {
		return "events." + string(*e.Resource)
	}

	return "events." + string(*e.Resource) + "." + e.ResourceId
}
