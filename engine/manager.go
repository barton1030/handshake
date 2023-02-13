package engine

import (
	inter "handshake/Interface"
	"sync"
)

// 管理单元用于控制器管理、监听topic领域对象变化并根据变化情况作出反应

type manager struct {
	controllerCollection map[string]*controller
	lock                 sync.Mutex
}

var ManagerUnit = manager{controllerCollection: make(map[string]*controller)}

func (m *manager) RegisterTopic(topic inter.Topic) bool {
	topicName := topic.Name()
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.controllerCollection[topicName]; ok {
		return true
	}
	topicController := newController(topic)
	startResult := topicController.start()
	if !startResult {
		return false
	}
	m.controllerCollection[topicName] = topicController
	return true
}

func (m *manager) CancelTopic(topic inter.Topic) bool {
	topicName := topic.Name()
	m.lock.Lock()
	defer m.lock.Unlock()
	topicController, ok := m.controllerCollection[topicName]
	if !ok {
		return true
	}
	stopResult := topicController.stop()
	if !stopResult {
		return false
	}
	delete(m.controllerCollection, topicName)
	return true
}
