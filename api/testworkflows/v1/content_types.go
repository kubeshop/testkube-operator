package v1

import (
	corev1 "k8s.io/api/core/v1"

	testsv3 "github.com/kubeshop/testkube-operator/api/tests/v3"
)

type ContentGit struct {
	// uri for the Git repository
	Uri string `json:"uri,omitempty" expr:"template"`
	// branch, commit or a tag name to fetch
	Revision string `json:"revision,omitempty" expr:"template"`
	// plain text username to fetch with
	Username string `json:"username,omitempty" expr:"template"`
	// external username to fetch with
	UsernameFrom *corev1.EnvVarSource `json:"usernameFrom,omitempty"`
	// plain text token to fetch with
	Token string `json:"token,omitempty" expr:"template"`
	// external token to fetch with
	TokenFrom *corev1.EnvVarSource `json:"tokenFrom,omitempty"`
	// authorization type for the credentials
	AuthType testsv3.GitAuthType `json:"authType,omitempty" expr:"template"`
	// where to mount the fetched repository contents (defaults to "repo" directory in the data volume)
	MountPath string `json:"mountPath,omitempty" expr:"template"`
	// paths to fetch for the sparse checkout
	Paths []string `json:"paths,omitempty" expr:"template"`
}

type ContentFile struct {
	// path where the file should be accessible at
	// +kubebuilder:validation:MinLength=1
	Path string `json:"path" expr:"template"`
	// plain-text content to put inside
	Content string `json:"content,omitempty" expr:"template"`
	// external source to use
	ContentFrom *corev1.EnvVarSource `json:"contentFrom,omitempty"`
	// mode to use for the file
	Mode *int32 `json:"mode,omitempty"`
}

type Content struct {
	// git repository details
	Git *ContentGit `json:"git,omitempty" expr:"include"`
	// files to load
	Files []ContentFile `json:"files,omitempty" expr:"include"`
}
