package engine

import inter "handshake/Interface"

// 控制单元负责执行器调度、错误执行数据统计、熔断保护计算触发、执行器管理等

type controller struct {
	topic    inter.Topic
	actuator map[string]*actuator
}

func newController(topic inter.Topic) *controller {
	return &controller{
		topic:    topic,
		actuator: make(map[string]*actuator),
	}
}

// 创建控制器、执行器通信管道
func (c *controller) createPipe() {
	pipeName := c.topic.Name()
	pipeCap := c.topic.MaxConcurrency()
	conduitUnit.setUpErrorConduit(pipeName, pipeCap)
	conduitUnit.setUpFusingConduit(pipeName, pipeCap)
	conduitUnit.setUpStatisticsConduit(pipeName, pipeCap)
}

func (c *controller) schedule() {

}
