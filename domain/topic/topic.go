package topic

import (
	inter "handshake/Interface"
	"handshake/engine"
	"time"
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
	queue          MessageQueuing
	creator        int
	createTime     time.Time
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
		queue:          newMessageQueuing(name),
		creator:        creator,
		createTime:     time.Now(),
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

func (t *topic) SetMinConcurrency(minConcurrency int) {
	t.minConcurrency = minConcurrency
}

func (t *topic) MaxConcurrency() int {
	return t.maxConcurrency
}

func (t *topic) SetMaxConcurrency(maxConcurrency int) {
	t.maxConcurrency = maxConcurrency
}

func (t *topic) FuseSalt() int {
	return t.fuseSalt
}

func (t *topic) SetFuseSalt(fuseSalt int) {
	t.fuseSalt = fuseSalt
}

func (t *topic) MaxRetryCount() int {
	return t.maxRetryCount
}

func (t *topic) SetMaxRetryCount(maxRetryCount int) {
	t.maxRetryCount = maxRetryCount
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

func (t *topic) CreateTime() time.Time {
	return t.createTime
}

func (t *topic) DiscardOrNot() bool {
	if t.status != DeleteStatus {
		return false
	}
	return true
}

func (t *topic) Start() bool {
	if t.status == DeleteStatus {
		return false
	}
	if t.status == StartStatus {
		return true
	}
	t.status = StartStatus
	return true
}

func (t *topic) StartUp() bool {
	if t.status != StartStatus {
		return false
	}
	startResult := engine.ManagerUnit.RegisterTopic(t)
	return startResult
}

func (t *topic) Stop() bool {
	if t.status == DeleteStatus {
		return false
	}
	if t.status == StopStatus {
		return true
	}
	t.status = StopStatus
	return true
}

func (t *topic) StopUp() bool {
	if t.status != StopStatus {
		return false
	}
	stopResult := engine.ManagerUnit.CancelTopic(t)
	return stopResult
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
