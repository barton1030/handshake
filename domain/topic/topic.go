package topic

import inter "handshake/Interface"

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

func NewTopic(name string, maxRetryCount, minConcurrency, maxConcurrency, fuseSalt, creator int) topic {
	return topic{
		status:         0,
		name:           name,
		maxRetryCount:  maxRetryCount,
		minConcurrency: minConcurrency,
		maxConcurrency: maxConcurrency,
		fuseSalt:       fuseSalt,
		queue:          newMessageQueuing(name),
		creator:        creator,
	}
}

func (t *topic) Id() (id int) {
	id = t.id
	return
}

func (t *topic) Name() (name string) {
	name = t.name
	return
}

func (t *topic) Status() (status int) {
	status = t.status
	return
}

func (t *topic) MinConcurrency() (minConcurrency int) {
	minConcurrency = t.minConcurrency
	return
}

func (t *topic) MaxConcurrency() (maxConcurrency int) {
	maxConcurrency = t.maxConcurrency
	return
}

func (t *topic) FuseSalt() (fuseSalt int) {
	fuseSalt = t.fuseSalt
	return
}

func (t *topic) MaxRetryCount() (maxRetryCount int) {
	maxRetryCount = t.maxRetryCount
	return
}

func (t *topic) CallbackHandler() (callback inter.Callback) {
	return
}

func (t *topic) AlarmHandler() (alarm inter.Alarm) {
	return
}

func (t *topic) MessageQueuingHandler() (messageQueuing inter.MessageQueuing) {
	return
}

func (t *topic) Recipients() (recipients []interface{}) {
	return
}

func (t *topic) SetAlarm(alarm inter.Alarm) (err error) {
	return
}

func (t *topic) SetCallback(callback inter.Callback) (err error) {
	return
}

func (t *topic) Creator() (creatorId int) {
	creatorId = t.creator
	return
}
