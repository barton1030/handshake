package topic

import (
	inter "handshake/Interface"
	"handshake/persistent"
	"sync"
)

type MessageQueuing struct {
	topicName string
	offset    int
	storage   inter.StorageQueueList
	lock      sync.Mutex
}

func newMessageQueuing(topicName string) MessageQueuing {
	return MessageQueuing{
		topicName: topicName,
		storage:   persistent.QueueDao,
		offset:    1,
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
	if err != nil {
		return nil, err
	}
	message3 := m.reconstruction(message2)
	if message3.Id() > 0 {
		m.offset = message3.id + 1
	}
	message = &message3
	return
}

func (m *MessageQueuing) Push(message inter.Message) (err error) {
	err = m.storage.Add(m.topicName, message)
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
	if message3 == nil {
		return message2
	}
	message2.id = message3.Id()
	message2.status = message3.Status()
	message2.data = message3.Data()
	message2.retry = message3.RetryCount()
	message2.createTime = message3.CreateTime()
	return
}
