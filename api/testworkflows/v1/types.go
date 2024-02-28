package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type EnvVar struct {
	// name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name,omitempty" expr:"template"`

	// value for the environment variable
	Value string `json:"value,omitempty" expr:"template"`

	// external value for the environment variable
	ValueFrom *corev1.EnvVarSource `json:"valueFrom,omitempty"`
}

type ContainerConfig struct {
	// override default working directory in the image (empty string to default WORKDIR for the image)
	WorkingDir *string `json:"workingDir,omitempty" expr:"template"`

	// image to be used for the container
	Image string `json:"image,omitempty" expr:"template"`

	// pulling policy for the image
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty" expr:"template"`

	// environment variables to append to the container
	Env []EnvVar `json:"env,omitempty" expr:"include"`

	// external environment variables to append to the container
	EnvFrom []corev1.EnvFromSource `json:"envFrom,omitempty"`

	// override default command in the image (empty string to default ENTRYPOINT of the image)
	Command *[]string `json:"command,omitempty" expr:"template"`

	// override default command in the image (empty string to default CMD of the image)
	Args *[]string `json:"args,omitempty" expr:"template"`

	// expected resources for the container
	Resources *Resources `json:"resources,omitempty" expr:"include"`

	// security context for the container
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`
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
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`

	// node selector to define on which node the pod should land
	NodeSelector map[string]string `json:"nodeSelector,omitempty" expr:"template,template"`

	// labels added to the scheduled pod
	Labels map[string]string `json:"labels,omitempty" expr:"template,template"`

	// annotations added to the scheduled pod
	Annotations map[string]string `json:"annotations,omitempty" expr:"template,template"`
}
