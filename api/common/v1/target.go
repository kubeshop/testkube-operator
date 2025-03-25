package v1

type Target struct {
	Match     map[string][]string `json:"match,omitempty"`
	Not       map[string][]string `json:"not,omitempty"`
	Replicate []string            `json:"replicate,omitempty"`
}
