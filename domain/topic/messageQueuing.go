package topic

import (
	inter "handshake/Interface"
	"handshake/persistent"
)

type MessageQueuing struct {
	topicName string
	offset    int
	storage   inter.StorageQueueList
}

func newMessageQueuing(topicName string) MessageQueuing {
	return MessageQueuing{
		topicName: topicName,
		storage:   persistent.QueueDao,
	}
}

func (m *MessageQueuing) SetStorage(storageInter inter.StorageQueueList) *MessageQueuing {
	return &MessageQueuing{
		topicName: m.topicName,
		offset:    m.offset,
		storage:   storageInter,
	}
}

func (m *MessageQueuing) Pop() (message inter.Message, err error) {
	message2, err := m.storage.NextPendingData(m.topicName, m.offset)
	message3 := m.reconstruction(message2)
	message = &message3
	m.offset = message3.id + 1
	return
}

func (m *MessageQueuing) Push(message inter.Message) (err error) {
	messageData := NewMessage(message.Data())
	err = m.storage.Add(m.topicName, &messageData)
	return
}

func (m *MessageQueuing) Finish(message inter.Message) (err error) {
	err = m.storage.Edit(m.topicName, message)
	return
}

func (m *MessageQueuing) Count() (count int) {
	count, _ = m.storage.PendingDataCount(m.topicName)
	return
}

func (m *MessageQueuing) reconstruction(message3 inter.Message) (message2 message) {
	message2.id = message3.Id()
	message2.status = message3.Status()
	message2.data = message3.Data()
	message2.retry = message3.RetryCount()
	message2.createTime = message3.CreateTime()
	return
}
