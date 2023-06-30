package testtrigger

import (
	testtriggerv1 "github.com/kubeshop/testkube-operator/apis/testtriggers/v1"
)

func GetSupportedResources() []string {
	return []string{
		ResourcePod,
		ResourceDeployment,
		ResourceStatefulSet,
		ResourceDaemonSet,
		ResourceService,
		ResourceCustomResource,
		ResourceIngress,
		ResourceEvent,
		ResourceConfigMap,
	}
}

func GetSupportedActions() []string {
	return []string{ActionRun}
}

func GetSupportedExecutions() []string {
	return []string{ExecutionTest, ExecutionTestsuite}
}

func GetSupportedConditionStatuses() []string {
	return []string{
		string(testtriggerv1.TRUE_TestTriggerConditionStatuses),
		string(testtriggerv1.FALSE_TestTriggerConditionStatuses),
		string(testtriggerv1.UNKNOWN_TestTriggerConditionStatuses),
	}
}

func GetSupportedConditions() []string {
	return []string{ConditionAvailable, ConditionContainersReady, ConditionInitialized, ConditionPodHasNetwork,
		ConditionPodScheduled, ConditionProgressing, ConditionReady, ConditionReplicaFailure}
}
