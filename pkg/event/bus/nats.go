package bus

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/nats-io/nats.go"
	errs "github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/kubeshop/testkube-operator/api/events/v1"
)

var (
	_ Bus = (*NATSBus)(nil)
)

const (
	SubscribeBuffer  = 1
	SubscriptionName = "events"
)

func NewNATSConnection(uri string) (*nats.EncodedConn, error) {
	nc, err := nats.Connect(uri)
	if err != nil {
		ctrl.Log.Error(err, "error connecting to nats")
		return nil, err
	}

	// automatic NATS JSON CODEC
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		ctrl.Log.Error(err, "error adding encoder to nats connection")
		return nil, err
	}

	return ec, nil
}

func NewNATSBus(nc *nats.EncodedConn) *NATSBus {
	return &NATSBus{
		nc: nc,
	}
}

type NATSBus struct {
	nc            *nats.EncodedConn
	subscriptions sync.Map
}

// Publish publishes event to NATS on events topic
func (n *NATSBus) Publish(event events.Event) error {
	return n.PublishTopic(SubscriptionName, event)
}

// Subscribe subscribes to NATS events topic
func (n *NATSBus) Subscribe(queueName string, handler Handler) error {
	return n.SubscribeTopic(SubscriptionName, queueName, handler)
}

// PublishTopic publishes event to NATS on given topic
func (n *NATSBus) PublishTopic(topic string, event events.Event) error {
	return n.nc.Publish(topic, event)
}

// SubscribeTopic subscribes to NATS topic
func (n *NATSBus) SubscribeTopic(topic, queueName string, handler Handler) error {
	// sanitize names for NATS
	queue, err := ListenerName(queueName)
	if err != nil {
		return errs.Wrap(err, "could not sanitize queue name")
	}

	// async subscribe on queue
	s, err := n.nc.QueueSubscribe(topic, queue, handler)

	if err == nil {
		// store subscription for later unsubscribe
		key := n.queueName(SubscriptionName, queue)
		n.subscriptions.Store(key, s)
	}

	return err
}

func (n *NATSBus) Unsubscribe(queueName string) error {
	// sanitize names for NATS
	queue, err := ListenerName(queueName)
	if err != nil {
		return errs.Wrap(err, "could not sanitize queue name")
	}

	key := n.queueName(SubscriptionName, queue)
	if s, ok := n.subscriptions.Load(key); ok {
		return s.(*nats.Subscription).Drain()
	}
	return nil
}

func (n *NATSBus) Close() error {
	n.nc.Close()
	return nil
}

func (n *NATSBus) queueName(subscription, queue string) string {
	return fmt.Sprintf("%s.%s", SubscriptionName, queue)
}

// ListenerName returns name of listener which can be used by event bus to identify listener
func ListenerName(in string) (string, error) {
	reg, err := regexp.Compile("[^A-Za-z0-9.]+")
	if err != nil {
		ctrl.Log.Error(err, "error sanitizing listener name")
		return "", err
	}
	return reg.ReplaceAllString(in, ""), nil
}
