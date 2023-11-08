package event

import (
	"github.com/kubeshop/testkube-operator/api/events/v1"
	"github.com/kubeshop/testkube-operator/pkg/event/common"

	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	ListenerKindWebsocket string = "websocket"
	ListenerKindSlack     string = "slack"
	ListenerKindWebhook   string = "webhook"
)

type Listener interface {
	// Name uniquely identifies listener
	Name() string
	// Notify sends event to listener
	Notify(event events.Event) events.EventResult
	// Kind of listener
	Kind() string
	// Selector is used to filter events
	Selector() string
	// Event is used to filter events
	Events() []events.EventType
	// Metadata with additional information about listener
	Metadata() map[string]string
}

type ListenerLoader interface {
	// Load listeners from configuration
	Load() (listeners common.Listeners, err error)
	// Kind of listener
}

func NewLoader() *Loader {
	return &Loader{
		Loaders: make([]common.ListenerLoader, 0),
	}
}

// Loader updates list of available listeners in the background as we don't want to load them on each event
type Loader struct {
	Loaders []common.ListenerLoader
}

// Register registers new listener reconciler
func (s *Loader) Register(loader common.ListenerLoader) {
	s.Loaders = append(s.Loaders, loader)
}

// Reconcile loop for reconciling listeners from different sources
func (s *Loader) Reconcile() (listeners common.Listeners) {
	listeners = make(common.Listeners, 0)
	for _, loader := range s.Loaders {
		l, err := loader.Load()
		if err != nil {
			ctrl.Log.Error(err, "error loading listeners")
			continue
		}
		listeners = append(listeners, l...)
	}

	return listeners
}
