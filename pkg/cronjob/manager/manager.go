package manager

import (
	"context"
	"strconv"

	configmapclient "github.com/kubeshop/testkube-operator/pkg/configmap"
	cronjobclient "github.com/kubeshop/testkube-operator/pkg/cronjob/client"
	namespaceclient "github.com/kubeshop/testkube-operator/pkg/namespace"
)

const (
	enableCronJobsFLagName = "enable-cron-jobs"
)

//go:generate mockgen -destination=./mock_client.go -package=manager "github.com/kubeshop/testkube-operator/pkg/cronjob/manager" Interface
type Interface interface {
	IsNamespaceForNewArchitecture(ctx context.Context, namespace string) (bool, error)
	CleanForNewArchitecture(ctx context.Context) error
}

// Manager provide methods to manage cronjobs
type Manager struct {
	namespaceClient namespaceclient.Interface
	configMapClient configmapclient.Interface
	cronJobClient   cronjobclient.Interface
	configMapName   string
}

// New is a method to create new cronjob manager
func New(namespaceClient namespaceclient.Interface, configMapClient configmapclient.Interface,
	cronJobClient cronjobclient.Interface, configMapName string) *Manager {
	return &Manager{
		namespaceClient: namespaceClient,
		configMapClient: configMapClient,
		cronJobClient:   cronJobClient,
		configMapName:   configMapName,
	}
}

func (m *Manager) IsNamespaceForNewArchitecture(ctx context.Context, namespace string) (bool, error) {
	if m.configMapName == "" {
		return false, nil
	}

	data, err := m.configMapClient.Get(ctx, m.configMapName, namespace)
	if err != nil {
		return false, err
	}

	if data == nil {
		return false, nil
	}

	var result bool
	if flag, ok := data[enableCronJobsFLagName]; ok && flag != "" {
		result, err = strconv.ParseBool(flag)
		if err != nil {
			return false, err
		}
	}

	return result, nil
}

// CleanForNewArchitecture is a method to clean cronjobs for new architecture
func (m *Manager) CleanForNewArchitecture(ctx context.Context) error {
	list, err := m.namespaceClient.ListAll(ctx, "")
	if err != nil {
		return err
	}

	namespaces := make([]string, 0)
	for _, namespace := range list.Items {
		result, err := m.IsNamespaceForNewArchitecture(ctx, namespace.Name)
		if err != nil {
			return err
		}

		if result {
			namespaces = append(namespaces, namespace.Name)
		}
	}

	for _, namespace := range namespaces {
		if err = m.cronJobClient.DeleteAll(ctx, cronjobclient.TestWorkflowResourceURI, namespace); err != nil {
			return err
		}
	}

	return nil
}
