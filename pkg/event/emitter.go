package event

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kubeshop/testkube-operator/api/events/v1"
	"github.com/kubeshop/testkube-operator/pkg/event/bus"
	"github.com/kubeshop/testkube-operator/pkg/event/common"
	ctrl "sigs.k8s.io/controller-runtime"
)

const (
	eventsBuffer      = 10000
	workersCount      = 20
	reconcileInterval = time.Second
)

// NewEmitter returns new emitter instance
func NewEmitter(eventBus bus.Bus, clusterName string) *Emitter {
	return &Emitter{
		Results:     make(chan events.EventResult, eventsBuffer),
		Loader:      NewLoader(),
		Bus:         eventBus,
		Listeners:   make(common.Listeners, 0),
		ClusterName: clusterName,
	}
}

// Emitter handles events emitting for webhooks
type Emitter struct {
	Results     chan events.EventResult
	Listeners   common.Listeners
	Loader      *Loader
	mutex       sync.Mutex
	Bus         bus.Bus
	ClusterName string
}

// Register adds new listener
func (e *Emitter) Register(listener common.Listener) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.Listeners = append(e.Listeners, listener)
}

// UpdateListeners updates listeners list
func (e *Emitter) UpdateListeners(listeners common.Listeners) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	oldMap := make(map[string]map[string]common.Listener, 0)
	newMap := make(map[string]map[string]common.Listener, 0)
	result := make([]common.Listener, 0)

	for _, l := range e.Listeners {
		if _, ok := oldMap[l.Kind()]; !ok {
			oldMap[l.Kind()] = make(map[string]common.Listener, 0)
		}

		oldMap[l.Kind()][l.Name()] = l
	}

	for _, l := range listeners {
		if _, ok := newMap[l.Kind()]; !ok {
			newMap[l.Kind()] = make(map[string]common.Listener, 0)
		}

		newMap[l.Kind()][l.Name()] = l
	}

	// check for missing listeners
	for kind, lMap := range oldMap {
		// clean missing kinds
		if _, ok := newMap[kind]; !ok {
			for _, l := range lMap {
				e.stopListener(l.Name())
			}

			continue
		}

		// stop missing listeners
		for name, l := range lMap {
			if _, ok := newMap[kind][name]; !ok {
				e.stopListener(l.Name())
			}
		}
	}

	// check for new listeners
	for kind, lMap := range newMap {
		// start all listeners for new kind
		if _, ok := oldMap[kind]; !ok {
			for _, l := range lMap {
				e.startListener(l)
				result = append(result, l)
			}

			continue
		}

		// start new listeners and restart updated ones
		for name, l := range lMap {
			if current, ok := oldMap[kind][name]; !ok {
				e.startListener(l)
			} else {
				if !common.CompareListeners(current, l) {
					e.stopListener(current.Name())
					e.startListener(l)
				}
			}

			result = append(result, l)
		}
	}

	e.Listeners = result
}

// Notify notifies emitter with webhook
func (e *Emitter) Notify(event events.Event) {
	event.ClusterName = e.ClusterName
	err := e.Bus.PublishTopic(event.Topic(), event)
	if err != nil {
		ctrl.Log.Error(err, "event could not be published", append(event.Log(), "error", err)...)
	}
}

// Listen runs emitter workers responsible for sending HTTP requests
func (e *Emitter) Listen(ctx context.Context) {
	// clean after closing Emitter
	go func() {
		<-ctx.Done()
		ctrl.Log.Info("closing event bus")

		for _, l := range e.Listeners {
			go e.Bus.Unsubscribe(l.Name())
		}

		e.Bus.Close()
	}()

	e.mutex.Lock()
	defer e.mutex.Unlock()

	for _, l := range e.Listeners {
		go e.startListener(l)
	}
}

func (e *Emitter) startListener(l common.Listener) {
	ctrl.Log.Info("starting listener", l.Name(), l.Metadata())
	err := e.Bus.SubscribeTopic("events.>", l.Name(), e.notifyHandler(l))
	if err != nil {
		ctrl.Log.Error(err, "error subscribing to event")
	}
}

func (e *Emitter) stopListener(name string) {
	ctrl.Log.Info("stoping listener", name)
	err := e.Bus.Unsubscribe(name)
	if err != nil {
		ctrl.Log.Error(err, "error unsubscribing from event")
	}
}

func (e *Emitter) notifyHandler(l common.Listener) bus.Handler {
	logOpts := map[string]string{
		"listen-on":   fmt.Sprintf("%v", l.Events()),
		"queue-group": l.Name(),
		"selector":    l.Selector(),
		"metadata":    fmt.Sprintf("%v", l.Metadata()),
	}
	return func(event events.Event) error {
		if event.Valid(l.Selector(), l.Events()) {
			res := l.Notify(event)
			ctrl.Log.Info("notification result: %v", res, logOpts)
			ctrl.Log.Info("listener notified", event.Log()...)
		} else {
			ctrl.Log.Info("dropping event not matching selector or type", event.Log()...)
		}
		return nil
	}
}

// Reconcile reloads listeners from all registered reconcilers
func (e *Emitter) Reconcile(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			ctrl.Log.Info("stopping reconciler")
			return
		default:
			listeners := e.Loader.Reconcile()
			e.UpdateListeners(listeners)
			time.Sleep(reconcileInterval)
		}
	}
}
