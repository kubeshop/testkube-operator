package testtrigger

type EventType string
type ResourceType string
type Cause string

const (
	ExecutionTest                               = "test"
	ExecutionTestsuite                          = "testsuite"
	ExecutionTestWorkflow                       = "testworkflow"
	ActionRun                                   = "run"
	ConcurrencyPolicyAllow                      = "allow"
	ConcurrencyPolicyForbid                     = "forbid"
	ConcurrencyPolicyReplace                    = "replace"
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
	CauseDeploymentScaleUpdate        Cause     = "deployment-scale-update"
	CauseDeploymentImageUpdate        Cause     = "deployment-image-update"
	CauseDeploymentEnvUpdate          Cause     = "deployment-env-update"
	CauseDeploymentContainersModified Cause     = "deployment-containers-modified"
	ConditionAvailable                          = "Available"
	ConditionProgressing                        = "Progressing"
	ConditionReplicaFailure                     = "ReplicaFailure"
	ConditionPodScheduled                       = "PodScheduled"
	ConditionPodHasNetwork                      = "PodHasNetwork"
	ConditionContainersReady                    = "ContainersReady"
	ConditionInitialized                        = "Initialized"
	ConditionReady                              = "Ready"
)
