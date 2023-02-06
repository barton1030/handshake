package topic

import "time"

type message struct {
	id         int
	data       map[string]interface{}
	status     int
	retry      int
	createTime time.Time
}

const (
	initStatus    = 1
	successStatus = 2
	failStatus    = -1
)

func NewMessage(data map[string]interface{}) message {
	return message{
		status:     initStatus,
		data:       data,
		createTime: time.Now(),
	}
}

func (m *message) Id() int {
	return m.id
}

func (m *message) Status() int {
	return m.status
}

func (m *message) RetryCount() int {
	return m.retry
}

func (m *message) IncrRetryCont() {
	m.retry++
}

func (m *message) Data() map[string]interface{} {
	return m.data
}

func (m *message) Processable() (processable bool) {
	return
}

func (m *message) Success() {
	if m.status == successStatus {
		return
	}
	m.status = successStatus
}

func (m *message) Fail() {
	if m.status == failStatus {
		return
	}
	m.status = failStatus
}

func (m *message) CreateTime() time.Time {
	return m.createTime
}
