package cronjob

import (
	"fmt"
)

// GetMetadataName returns cron job metadata name
func GetMetadataName(name, resource string) string {
	return fmt.Sprintf("%s-%s", name, resource)
}
