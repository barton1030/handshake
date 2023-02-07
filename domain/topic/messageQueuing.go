package topic

import (
	"fmt"
	inter "handshake/Interface"
	"handshake/persistent"
)

type messageQueuing struct {
	topicName string
	nextId    int
	offset    int
	storage   inter.StorageQueueList
}

func newMessageQueuing(topicName string) messageQueuing {
	return messageQueuing{
		topicName: topicName,
		storage:   persistent.QueueDao,
		nextId:    1,
	}
}

func (m *messageQueuing) init() {
	maxPrimaryKeyId := m.storage.MaxPrimaryKeyId(m.topicName)
	if maxPrimaryKeyId < 0 {
		return
	}
	m.nextId = maxPrimaryKeyId + 1
}

func (m *messageQueuing) Pop() (message inter.Message, err error) {
	message2, err := m.storage.NextPendingData(m.topicName, m.offset)
	message3 := m.reconstruction(message2)
	message = &message3
	m.offset = message3.id + 1
	fmt.Println(m.offset)
	return
}

func (m *messageQueuing) Push(message inter.Message) (err error) {
	messageData := NewMessage(message.Data())
	messageData.id = m.nextId
	err = m.storage.Add(m.topicName, &messageData)
	if err != nil {
		return
	}
	m.nextId++
	return
}

func (m *messageQueuing) Finish(message inter.Message) (err error) {
	err = m.storage.Edit(m.topicName, message)
	return
}

func (m *messageQueuing) Count() (count int) {
	count, _ = m.storage.PendingDataCount(m.topicName)
	return
}

func (m *messageQueuing) reconstruction(message3 inter.Message) (message2 message) {
	message2.id = message3.Id()
	message2.status = message3.Status()
	message2.data = message3.Data()
	message2.retry = message3.RetryCount()
	message2.createTime = message3.CreateTime()
	return
}
