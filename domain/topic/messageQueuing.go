package topic

import inter "handshake/Interface"

type messageQueuing struct {
	topicName string
	storage   map[int]interface{}
}

func (m *messageQueuing) Pop() (message inter.Message, err error) {
	data := make(map[string]interface{})
	data["name"] = "barton"
	messageData := NewMessage(data)
	message = &messageData
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
