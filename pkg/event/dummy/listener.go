package dummy

import (
	"sync/atomic"

	"github.com/kubeshop/testkube-operator/api/events/v1"
	"github.com/kubeshop/testkube-operator/pkg/event/common"
)

var _ common.Listener = (*DummyListener)(nil)

type DummyListener struct {
	Id                string
	NotificationCount int32
	SelectorString    string
}

func (l *DummyListener) GetNotificationCount() int {
	cnt := atomic.LoadInt32(&l.NotificationCount)
	return int(cnt)
}

func (l *DummyListener) Notify(event events.Event) events.EventResult {
	atomic.AddInt32(&l.NotificationCount, 1)
	return events.EventResult{Id: event.Id}
}

func (l *DummyListener) Name() string {
	if l.Id != "" {
		return l.Id
	}
	return "dummy"
}

func (l *DummyListener) Events() []events.EventType {
	return events.AllEventTypes
}

func (l *DummyListener) Selector() string {
	return l.SelectorString
}

func (l *DummyListener) Kind() string {
	return "dummy"
}

func (l *DummyListener) Metadata() map[string]string {
	return map[string]string{
		"id":       l.Name(),
		"selector": l.Selector(),
	}
}
