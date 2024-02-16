package v1

import (
	corev1 "k8s.io/api/core/v1"

	testsv3 "github.com/kubeshop/testkube-operator/api/tests/v3"
)

type ContentGit struct {
	// uri for the Git repository
	Uri string `json:"uri,omitempty"`
	// branch, commit or a tag name to fetch
	Revision string `json:"revision,omitempty"`
	// plain text username to fetch with
	Username *string `json:"username,omitempty"`
	// external username to fetch with
	UsernameFrom *corev1.EnvVarSource `json:"usernameFrom,omitempty"`
	// plain text token to fetch with
	Token *string `json:"token,omitempty"`
	// external token to fetch with
	TokenFrom *corev1.EnvVarSource `json:"tokenFrom,omitempty"`
	// authorization type for the credentials
	AuthType testsv3.GitAuthType `json:"authType,omitempty"`
	// paths to fetch for the sparse checkout
	Paths []string `json:"paths,omitempty"`
	// where to mount the fetched repository contents (defaults to "repo" directory in the data volume)
	MountPath string `json:"mountPath,omitempty"`
}

type ContentFile struct {
	// path where the file should be accessible at
	// +kubebuilder:validation:MinLength=1
	Path string `json:"path"`
	// plain-text content to put inside
	Content string `json:"content,omitempty"`
	// external source to use
	ContentFrom *corev1.EnvVarSource `json:"contentFrom,omitempty"`
	// mode to use for the file
	Mode *int32 `json:"mode,omitempty"`
}

type Content struct {
	// git repository details
	Git *ContentGit `json:"git,omitempty"`
	// files to load
	Files []ContentFile `json:"files,omitempty"`
}
