package topic

import inter "handshake/Interface"

type messageQueuing struct {
	topicName string
	storage   map[int]interface{}
}

func newMessageQueuing(topicName string) messageQueuing {
	return messageQueuing{
		topicName: topicName,
		storage:   make(map[int]interface{}),
	}
}

func (m *messageQueuing) Pop() (message inter.Message, err error) {
	data := make(map[string]interface{})
	data["name"] = "barton"
	domainMessage := NewMessage(data)
	message = &domainMessage
	return
}

func (m *messageQueuing) Push(message inter.Message) (err error) {
	return
}

func (m *messageQueuing) Finish(message inter.Message) (err error) {
	return
}

func (m *messageQueuing) Count() (count int) {
	return
}
