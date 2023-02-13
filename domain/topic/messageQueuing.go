package topic

import (
	inter "handshake/Interface"
	"handshake/persistent"
	"sync"
	"time"
)

type MessageQueuing struct {
	topicName string
	storage   inter.StorageQueueList
	lock      sync.Mutex
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
		storage:   storageInter,
	}
}

func (m *MessageQueuing) Pop() (message inter.Message, err error) {
	message2, err := m.storage.NextPendingData(m.topicName)
	if err != nil {
		return nil, err
	}
	message3 := m.reconstruction(message2)
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

// 触发器暂时不再嵌入到队列对象中
type trigger struct {
	lastTriggerTime time.Time
	triggerTime     time.Time
	topicName       string
	lock            sync.Mutex
}

func (t *trigger) Send(triggerTime time.Time) {
	t.lock.Lock()
	defer t.lock.Unlock()
	nextExecutionTime := t.lastTriggerTime.Add(5 * time.Minute)
	if !triggerTime.After(nextExecutionTime) {
		return
	}
}

func (t *trigger) LastTriggerTime() time.Time {
	return t.lastTriggerTime
}

func (t *trigger) TopicName() string {
	return t.topicName
}

func (t *trigger) TriggerTime() time.Time {
	return t.triggerTime
}
