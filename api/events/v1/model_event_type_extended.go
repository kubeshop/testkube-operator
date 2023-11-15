package events

var AllEventTypes = []EventType{
	START_TEST_EventType,
	END_TEST_SUCCESS_EventType,
	END_TEST_FAILED_EventType,
	END_TEST_ABORTED_EventType,
	END_TEST_TIMEOUT_EventType,
	START_TESTSUITE_EventType,
	END_TESTSUITE_SUCCESS_EventType,
	END_TESTSUITE_FAILED_EventType,
	END_TESTSUITE_ABORTED_EventType,
	END_TESTSUITE_TIMEOUT_EventType,
	CREATED_EventType,
	DELETED_EventType,
	UPDATED_EventType,
}

func (t EventType) String() string {
	return string(t)
}

func EventTypePtr(t EventType) *EventType {
	return &t
}

var (
	EventStartTest           = EventTypePtr(START_TEST_EventType)
	EventEndTestSuccess      = EventTypePtr(END_TEST_SUCCESS_EventType)
	EventEndTestFailed       = EventTypePtr(END_TEST_FAILED_EventType)
	EventEndTestAborted      = EventTypePtr(END_TEST_ABORTED_EventType)
	EventEndTestTimeout      = EventTypePtr(END_TEST_TIMEOUT_EventType)
	EventStartTestSuite      = EventTypePtr(START_TESTSUITE_EventType)
	EventEndTestSuiteSuccess = EventTypePtr(END_TESTSUITE_SUCCESS_EventType)
	EventEndTestSuiteFailed  = EventTypePtr(END_TESTSUITE_FAILED_EventType)
	EventEndTestSuiteAborted = EventTypePtr(END_TESTSUITE_ABORTED_EventType)
	EventEndTestSuiteTimeout = EventTypePtr(END_TESTSUITE_TIMEOUT_EventType)
	EventCreated             = EventTypePtr(CREATED_EventType)
	EventDeleted             = EventTypePtr(DELETED_EventType)
	EventUpdated             = EventTypePtr(UPDATED_EventType)

	EventTestExecutionCreated = EventTypePtr(TEST_EXECUTION_CREATED_EventType)
	EventTestExecutionUpdated = EventTypePtr(TEST_EXECUTION_UPDATED_EventType)
	EventTestExecutionDeleted = EventTypePtr(TEST_EXECUTION_DELETED_EventType)

	EventTestCreated          = EventTypePtr(TEST_CREATED_EventType)
	EventTestUpdated          = EventTypePtr(TEST_UPDATED_EventType)
	EventTestDeleted          = EventTypePtr(TEST_DELETED_EventType)
	EventTestsDeletedAll      = EventTypePtr(TEST_DELETED_ALL_EventType)
	EventTestsDeletedFiltered = EventTypePtr(TEST_DELETED_FILTERED_EventType)

	EventTestSuiteExecutionCreated = EventTypePtr(TESTSUITE_EXECUTION_CREATED_EventType)
	EventTestSuiteExecutionUpdated = EventTypePtr(TESTSUITE_EXECUTION_UPDATED_EventType)
	EventTestSuiteExecutionDeleted = EventTypePtr(TESTSUITE_EXECUTION_DELETED_EventType)

	EventTestSuiteCreated          = EventTypePtr(TESTSUITE_CREATED_EventType)
	EventTestSuiteUpdated          = EventTypePtr(TESTSUITE_UPDATED_EventType)
	EventTestSuiteDeleted          = EventTypePtr(TESTSUITE_DELETED_EventType)
	EventTestSuitesDeletedAll      = EventTypePtr(TESTSUITE_DELETED_ALL_EventType)
	EventTestSuitesDeletedFiltered = EventTypePtr(TESTSUITE_DELETED_FILTERED_EventType)
)

func EventTypesFromSlice(types []string) []EventType {
	var t []EventType
	for _, v := range types {
		t = append(t, EventType(v))
	}
	return t
}

type EventOperationType string

const (
	CREATED EventOperationType = "created"
	UPDATED EventOperationType = "updated"
	DELETED EventOperationType = "deleted"
)
