package testtrigger

func GetSupportedResources() []string {
	return []string{
		ResourcePod,
		ResourceDeployment,
		ResourceStatefulSet,
		ResourceDaemonSet,
		ResourceService,
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
