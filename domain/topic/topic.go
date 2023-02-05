package topic

import (
	inter "handshake/Interface"
	"handshake/engine"
)

type topic struct {
	id             int
	status         int
	name           string
	maxRetryCount  int
	minConcurrency int
	maxConcurrency int
	fuseSalt       int
	alarm          alarm
	callback       callback
	queue          messageQueuing
	creator        int
}

const (
	StopStatus   = 1
	StartStatus  = 2
	DeleteStatus = -1
)

func NewTopic(name string, maxRetryCount, minConcurrency, maxConcurrency, fuseSalt, creator int) topic {
	return topic{
		status:         StopStatus,
		name:           name,
		maxRetryCount:  maxRetryCount,
		minConcurrency: minConcurrency,
		maxConcurrency: maxConcurrency,
		fuseSalt:       fuseSalt,
		queue: messageQueuing{
			topicName: name,
			storage:   make(map[int]interface{}),
		},
		creator: creator,
	}
}

func (t *topic) Id() int {
	return t.id
}

func (t *topic) Name() string {
	return t.name
}

func (t *topic) Status() int {
	return t.status
}

func (t *topic) MinConcurrency() int {
	return t.minConcurrency
}

func (t *topic) MaxConcurrency() int {
	return t.maxConcurrency
}

func (t *topic) FuseSalt() int {
	return t.fuseSalt
}

func (t *topic) MaxRetryCount() int {
	return t.maxRetryCount
}

func (t *topic) CallbackHandler() inter.Callback {
	return &t.callback
}

func (t *topic) AlarmHandler() inter.Alarm {
	return &t.alarm
}

func (t *topic) MessageQueuingHandler() (queue inter.MessageQueuing) {
	return &t.queue
}

func (t *topic) SetAlarm(topicAlarm alarm) {
	t.alarm = topicAlarm
	return
}

func (t *topic) SetCallback(topicCallback callback) {
	t.callback = topicCallback
	return
}

func (t *topic) Creator() int {
	return t.creator
}

func (t *topic) Start() bool {
	if t.status == StartStatus {
		return true
	}
	t.status = StartStatus
	engine.ManagerUnit.RegisterTopic(t)
	return true
}

func (t *topic) Stop() bool {
	if t.status == StopStatus {
		return true
	}
	t.status = StopStatus
	engine.ManagerUnit.CancelTopic(t)
	return true
}

func (t *topic) InOperation() bool {
	if t.status == StopStatus {
		return false
	}
	return true
}

func (t *topic) Abandonment() bool {
	if t.status != StopStatus {
		return false
	}
	t.status = DeleteStatus
	return true
}
