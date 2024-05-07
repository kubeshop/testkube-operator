package v1

import (
	"encoding/json"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ContainerConfig struct {
	// override default working directory in the image (empty string to default WORKDIR for the image)
	WorkingDir *string `json:"workingDir,omitempty" expr:"template"`

	// image to be used for the container
	Image string `json:"image,omitempty" expr:"template"`

	// pulling policy for the image
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty" expr:"template"`

	// environment variables to append to the container
	Env []corev1.EnvVar `json:"env,omitempty" expr:"force"`

	// external environment variables to append to the container
	EnvFrom []corev1.EnvFromSource `json:"envFrom,omitempty" expr:"force"`

	// override default command in the image (empty string to default ENTRYPOINT of the image)
	Command *[]string `json:"command,omitempty" expr:"template"`

	// override default command in the image (empty string to default CMD of the image)
	Args *[]string `json:"args,omitempty" expr:"template"`

	// expected resources for the container
	Resources *Resources `json:"resources,omitempty" expr:"include"`

	// security context for the container
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty" expr:"force"`

	// volume mounts to append to the container
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty" expr:"force"`
}

type Resources struct {
	// resource limits for the container
	Limits map[corev1.ResourceName]intstr.IntOrString `json:"limits,omitempty" expr:"template,template"`

	// resource requests for the container
	Requests map[corev1.ResourceName]intstr.IntOrString `json:"requests,omitempty" expr:"template,template"`
}

type JobConfig struct {
	// labels added to the scheduled job
	Labels map[string]string `json:"labels,omitempty" expr:"template,template"`

	// annotations added to the scheduled job
	Annotations map[string]string `json:"annotations,omitempty" expr:"template,template"`
}

type PodConfig struct {
	// default service account name for the scheduled pod
	ServiceAccountName string `json:"serviceAccountName,omitempty" expr:"template"`

	// references to secrets with credentials for pulling the images from registry
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" expr:"force"`

	// node selector to define on which node the pod should land
	NodeSelector map[string]string `json:"nodeSelector,omitempty" expr:"template,template"`

	// labels added to the scheduled pod
	Labels map[string]string `json:"labels,omitempty" expr:"template,template"`

	// annotations added to the scheduled pod
	Annotations map[string]string `json:"annotations,omitempty" expr:"template,template"`

	// volumes to include in the pod
	Volumes []corev1.Volume `json:"volumes,omitempty" expr:"force"`

	// duration in seconds the pod may be active on the node
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty" expr:"ignore"`

	// DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.
	DNSPolicy corev1.DNSPolicy `json:"dnsPolicy,omitempty" expr:"template"`

	// NodeName is a request to schedule this pod onto a specific node.
	NodeName string `json:"nodeName,omitempty" expr:"template"`

	// SecurityContext holds pod-level security attributes and common container settings.
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty" expr:"force"`

	// Specifies the hostname of the Pod
	Hostname string `json:"hostname,omitempty" expr:"template"`

	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>".
	Subdomain string `json:"subdomain,omitempty" expr:"template"`

	// If specified, the pod's scheduling constraints
	Affinity *corev1.Affinity `json:"affinity,omitempty" expr:"force"`

	// If specified, the pod's tolerations.
	Tolerations []corev1.Toleration `json:"tolerations,omitempty" expr:"force"`

	// HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified
	HostAliases []corev1.HostAlias `json:"hostAliases,omitempty" expr:"force"`

	// If specified, indicates the pod's priority.
	PriorityClassName string `json:"priorityClassName,omitempty" expr:"template"`

	// The priority value. Various system components use this field to find the priority of the pod.
	Priority *int32 `json:"priority,omitempty" expr:"ignore"`

	// Specifies the DNS parameters of a pod.
	DNSConfig *corev1.PodDNSConfig `json:"dnsConfig,omitempty" expr:"force"`

	// PreemptionPolicy is the Policy for preempting pods with lower priority.
	PreemptionPolicy *corev1.PreemptionPolicy `json:"preemptionPolicy,omitempty" expr:"template"`

	// TopologySpreadConstraints describes how a group of pods ought to spread across topology domains.
	TopologySpreadConstraints []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty" expr:"force"`

	// SchedulingGates is an opaque list of values that if specified will block scheduling the pod.
	SchedulingGates []corev1.PodSchedulingGate `json:"schedulingGates,omitempty" expr:"force"`

	// ResourceClaims defines which ResourceClaims must be allocated and reserved before the Pod is allowed to start.
	ResourceClaims []corev1.PodResourceClaim `json:"resourceClaims,omitempty" expr:"force"`

	// namespace for execution of test workflow
	Namespace string `json:"namespace,omitempty"  expr:"template"`
}

type DynamicList struct {
	Dynamic    bool     `expr:"ignore"`
	Static     []string `expr:"template"`
	Expression string   `expr:"expression"`
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (s *DynamicList) UnmarshalJSON(value []byte) error {
	if value[0] == '[' {
		result := make([]interface{}, 0)
		err := json.Unmarshal(value, &result)
		if err != nil {
			return err
		}
		isStringOnly := true
		for i := range result {
			if _, ok := result[i].(string); !ok {
				isStringOnly = false
				break
			}
		}
		if isStringOnly {
			s.Dynamic = false
			s.Static = make([]string, len(result))
			for i := range result {
				s.Static[i] = result[i].(string)
			}
		} else {
			s.Dynamic = true
			s.Expression = string(value)
		}
		return nil
	}
	if value[0] == '"' {
		s.Dynamic = true
		return json.Unmarshal(value, &s.Expression)
	}
	s.Dynamic = true
	s.Expression = string(value)
	return nil
}

// MarshalJSON implements the json.Marshaller interface.
func (s DynamicList) MarshalJSON() ([]byte, error) {
	if s.Dynamic {
		var v []interface{}
		err := json.Unmarshal([]byte(s.Expression), &v)
		if err != nil {
			return json.Marshal(s.Expression)
		} else {
			return []byte(s.Expression), nil
		}
	}
	return json.Marshal(s.Static)
}

type Event struct {
	Cronjob *CronJobConfig `json:"cronjob,omitempty"`
}

// cron job configuration
type CronJobConfig struct {
	// cron schedule to run a test workflow
	Cron string `json:"cron" expr:"template"`
	// labels to attach to the cron job
	Labels map[string]string `json:"labels,omitempty" expr:"template,template"`
	// annotations to attach to the cron job
	Annotations map[string]string `json:"annotations,omitempty" expr:"template,template"`
}
