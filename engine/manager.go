package engine

import "sync"

// 管理单元用于控制器管理、监听topic领域对象变化并根据变化情况作出反应

type manager struct {
	controllerCollection map[string]*controller
	lock                 sync.Mutex
}

var managerUnit manager

func managerInit() {
	managerUnit.controllerCollection = make(map[string]*controller)
}

func (m *manager) MonitorTopic() {

}
