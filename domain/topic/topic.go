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
	StartStatus = 1
	StopStatus  = 0
)

func NewTopic(name string, maxRetryCount, minConcurrency, maxConcurrency, fuseSalt, creator int) topic {
	return topic{
		status:         0,
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

func (t *topic) SetAlarm(alarm inter.Alarm) (err error) {
	return
}

func (t *topic) SetCallback(callback inter.Callback) (err error) {
	return
}

func (t *topic) Creator() int {
	return t.creator
}

func (t *topic) Start() (err error) {
	if t.status == StartStatus {
		return err
	}
	t.status = StartStatus
	engine.ManagerUnit.RegisterTopic(t)
	return
}

func (t *topic) Stop() (err error) {
	if t.status == StopStatus {
		return err
	}
	t.status = StopStatus
	engine.ManagerUnit.CancelTopic(t)
	return
}
