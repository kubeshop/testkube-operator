package testtrigger

type EventType string
type ResourceType string
type Cause string

const (
	ExecutionTest                               = "test"
	ExecutionTestsuite                          = "testsuite"
	ActionRun                                   = "run"
	ResourcePod                                 = "pod"
	ResourceDeployment                          = "deployment"
	ResourceStatefulSet                         = "statefulset"
	ResourceDaemonSet                           = "daemonset"
	ResourceService                             = "service"
	ResourceIngress                             = "ingress"
	ResourceEvent                               = "event"
	ResourceConfigMap                           = "configmap"
	DefaultNamespace                            = "testkube"
	EventCreated                      EventType = "created"
	EventModified                     EventType = "modified"
	EventDeleted                      EventType = "deleted"
	CauseDeploymentScaleUpdate        Cause     = "deployment_scale_update"
	CauseDeploymentImageUpdate        Cause     = "deployment_image_update"
	CauseDeploymentEnvUpdate          Cause     = "deployment_env_update"
	CauseDeploymentContainersModified Cause     = "deployment_containers_modified"
)
