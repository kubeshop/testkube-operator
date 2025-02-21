package manager

import (
	"context"
	"strconv"

	configmapclient "github.com/kubeshop/testkube-operator/pkg/configmap"
	cronjobclient "github.com/kubeshop/testkube-operator/pkg/cronjob/client"
	namespaceclient "github.com/kubeshop/testkube-operator/pkg/namespace"
)

const (
	testkubeConfigMapName  = "tesrkube-api-server"
	enableCronJobsFLagName = "enable-cron-jobs"
)

//go:generate mockgen -destination=./mock_client.go -package=managerr "github.com/kubeshop/testkube-operator/pkg/cronjob/manager" Interface
type Interface interface {
	CleanForNewArchitecture(ctx context.Context) error
}

// Manager provide methods to manage cronjobs
type Manager struct {
	namespaceClient namespaceclient.Interface
	configMapClient configmapclient.Interface
	cronJobClient   cronjobclient.Interface
}

// New is a method to create new cronjob manager
func New(namespaceClient namespaceclient.Interface, configMapClient configmapclient.Interface, cronJobClient cronjobclient.Interface) *Manager {
	return &Manager{
		namespaceClient: namespaceClient,
		configMapClient: configMapClient,
		cronJobClient:   cronJobClient,
	}
}

// CleanForNewArchitecture is a method to clean cronjobs for new architecture
func (m *Manager) CleanForNewArchitecture(ctx context.Context) error {
	list, err := m.namespaceClient.ListAll(ctx, "")
	if err != nil {
		return err
	}

	namespaces := make([]string, 0)
	for _, namespace := range list.Items {
		data, err := m.configMapClient.Get(ctx, testkubeConfigMapName, namespace.Name)
		if err != nil {
			return err
		}

		if flag, ok := data[enableCronJobsFLagName]; ok && flag != "" {
			value, err := strconv.ParseBool(flag)
			if err != nil {
				return err
			}

			if value {
				namespaces = append(namespaces, namespace.Name)
			}
		}
	}

	for _, namespace := range namespaces {
		resources := []string{cronjobclient.TestResourceURI, cronjobclient.TestSuiteResourceURI, cronjobclient.TestWorkflowResourceURI}
		for _, resource := range resources {
			if err = m.cronJobClient.DeleteAll(ctx, resource, namespace); err != nil {
				return err
			}
		}
	}

	return nil
}
