package v3

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testsuitesv3 "github.com/kubeshop/testkube-operator/api/testsuite/v3"
)

type EventType string

const (
	EventTypeCreate EventType = "create"
	EventTypeUpdate EventType = "update"
	EventTypeDelete EventType = "delete"
)

type Update struct {
	Type      EventType
	Timestamp time.Time
	Resource  *testsuitesv3.TestSuite
}

type WatcherUpdate Watcher[Update]

//go:generate mockgen -source=./rest.go -destination=./mock_rest.go -package=v3 "github.com/kubeshop/testkube-operator/pkg/client/testsuites/v3" RESTInterface
type RESTInterface interface {
	WatchUpdates(ctx context.Context, environmentId string, includeInitialData bool) WatcherUpdate
}

// NewRESTClient creates new Test Suite client
func NewRESTClient(client client.Client, restConfig *rest.Config, namespace string) (*TestSuitesRESTClient, error) {
	// Build the scheme
	scheme := runtime.NewScheme()
	if err := metav1.AddMetaToScheme(scheme); err != nil {
		return nil, err
	}

	if err := testsuitesv3.SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}

	codecs := serializer.NewCodecFactory(scheme)
	parameterCodec := runtime.NewParameterCodec(scheme)

	// Build the REST client
	cfg := *restConfig
	gv := testsuitesv3.GroupVersion
	cfg.GroupVersion = &gv
	cfg.APIPath = "/apis"
	cfg.NegotiatedSerializer = codecs.WithoutConversion()
	httpClient, err := rest.HTTPClientFor(&cfg)
	if err != nil {
		return nil, err
	}

	restClient, err := rest.RESTClientForConfigAndClient(&cfg, httpClient)
	if err != nil {
		return nil, err
	}

	return &TestSuitesRESTClient{
		k8sClient:      client,
		restClient:     restClient,
		parameterCodec: parameterCodec,
		namespace:      namespace}, nil
}

// TestSuitesRESTClient implements REST methods to work with Test Suite
type TestSuitesRESTClient struct {
	k8sClient      client.Client
	restClient     *rest.RESTClient
	parameterCodec runtime.ParameterCodec
	namespace      string
}

func (s TestSuitesRESTClient) WatchUpdates(ctx context.Context, environmentId string, includeInitialData bool) WatcherUpdate {
	// Load initial data
	list := &testsuitesv3.TestSuiteList{}
	if includeInitialData {
		opts := &client.ListOptions{Namespace: s.namespace}
		if err := s.k8sClient.List(ctx, list, opts); err != nil {
			return NewError[Update](err)
		}
	}

	// Start watching
	opts := metav1.ListOptions{Watch: true, ResourceVersion: list.ResourceVersion}
	watcher, err := s.restClient.Get().
		Namespace(s.namespace).
		Resource("testsuites").
		VersionedParams(&opts, s.parameterCodec).
		Watch(ctx)
	if err != nil {
		return NewError[Update](err)
	}

	result := NewWatcher[Update]()
	go func() {
		// Send initial data
		for _, k8sObject := range list.Items {
			updateType := EventTypeCreate
			updateTime := getUpdateTime(k8sObject)
			if !updateTime.Equal(k8sObject.CreationTimestamp.Time) {
				updateType = EventTypeUpdate
			}

			result.Send(Update{
				Type:      updateType,
				Timestamp: updateTime,
				Resource:  &k8sObject,
			})
		}

		// Watch
		for event := range watcher.ResultChan() {
			// Continue watching if that's just a bookmark
			if event.Type == watch.Bookmark {
				continue
			}

			// Load the current Kubernetes object
			k8SObject, ok := event.Object.(*testsuitesv3.TestSuite)
			if !ok || k8SObject == nil {
				continue
			}

			// Handle Kubernetes flavours that do not have Deleted event
			if k8SObject.DeletionTimestamp != nil {
				event.Type = watch.Deleted
			}

			updateTime := getUpdateTime(*k8SObject)
			switch event.Type {
			case watch.Added:
				result.Send(Update{
					Type:      EventTypeCreate,
					Timestamp: updateTime,
					Resource:  k8SObject,
				})
			case watch.Modified:
				result.Send(Update{
					Type:      EventTypeUpdate,
					Timestamp: updateTime,
					Resource:  k8SObject,
				})
			case watch.Deleted:
				result.Send(Update{
					Type:      EventTypeDelete,
					Timestamp: updateTime,
					Resource:  k8SObject,
				})
			}
		}

		result.Close(context.Cause(ctx))
	}()

	return result
}

func getUpdateTime(t testsuitesv3.TestSuite) time.Time {
	updateTime := t.CreationTimestamp.Time
	if t.DeletionTimestamp != nil {
		updateTime = t.DeletionTimestamp.Time
	} else {
		for _, field := range t.ManagedFields {
			if field.Time.After(updateTime) {
				updateTime = field.Time.Time
			}
		}
	}

	return updateTime
}
