package topic

import "time"

type message struct {
	id         int
	data       map[string]interface{}
	status     int
	retry      int
	createTime time.Time
}

func NewMessage(data map[string]interface{}) message {
	return message{
		data:       data,
		createTime: time.Now(),
	}
}

func (m *message) Id() (id int) {
	return
}

func (m *message) RetryCount() (retryCont int) {
	return
}

func (m *message) IncrRetryCont() {

}

func (m *message) Data() (data map[string]interface{}, err error) {
	return
}

func (m *message) Processable() (processable bool) {
	return
}

func (m *message) Success() {

}

func (m *message) Fail() {

}
