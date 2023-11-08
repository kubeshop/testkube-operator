package bus

import (
	"github.com/kubeshop/testkube-operator/api/events/v1"
)

type Handler func(event events.Event) error

type Bus interface {
	Publish(event events.Event) error
	Subscribe(queue string, handler Handler) error
	Unsubscribe(queue string) error

	PublishTopic(topic string, event events.Event) error
	SubscribeTopic(topic string, queue string, handler Handler) error

	Close() error
}
