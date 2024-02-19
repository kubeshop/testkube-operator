package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type Expression string

type ContainerConfig struct {
	// override default working directory in the image (empty string to default WORKDIR for the image)
	WorkingDir *string `json:"workingDir"`

	// image to be used for the container
	Image string `json:"image,omitempty"`

	// pulling policy for the image
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// environment variables to append to the container
	Env []corev1.EnvVar `json:"env,omitempty"`

	// external environment variables to append to the container
	EnvFrom []corev1.EnvFromSource `json:"envFrom,omitempty"`

	// override default command in the image (empty string to default ENTRYPOINT of the image)
	Command *[]string `json:"command,omitempty"`

	// override default command in the image (empty string to default CMD of the image)
	Args *[]string `json:"args,omitempty"`

	// expected resources for the container
	Resources Resources `json:"resources,omitempty"`

	// security context for the container
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`
}

type Resources struct {
	// +kubebuilder:validation:XIntOrString
	// resource limits for the container
	Limits map[corev1.ResourceName]intstr.IntOrString `json:"limits,omitempty"`

	// +kubebuilder:validation:XIntOrString
	// resource requests for the container
	Requests map[corev1.ResourceName]intstr.IntOrString `json:"requests,omitempty"`
}

type JobConfig struct {
	// labels added to the scheduled job
	Labels map[string]string `json:"labels,omitempty"`

	// annotations added to the scheduled job
	Annotations map[string]string `json:"annotations,omitempty"`
}

type PodConfig struct {
	// default service account name for the scheduled pod
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// references to secrets with credentials for pulling the images from registry
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`

	// node selector to define on which node the pod should land
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// labels added to the scheduled pod
	Labels map[string]string `json:"labels,omitempty"`

	// annotations added to the scheduled pod
	Annotations map[string]string `json:"annotations,omitempty"`
}
