package engine

import (
	"fmt"
	inter "handshake/Interface"
	"sync"
	"time"
)

// 控制单元负责执行器调度、错误执行数据统计、熔断保护计算触发、执行器管理等

type controller struct {
	topic                        inter.Topic
	actuator                     map[string]*actuator
	timeSliceErrorStatistics     map[string]int
	timeSliceErrorStatisticsLock sync.Mutex
}

func newController(topic inter.Topic) *controller {
	return &controller{
		topic:                    topic,
		actuator:                 make(map[string]*actuator),
		timeSliceErrorStatistics: make(map[string]int),
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

func (c *controller) monitorPipe() {
	topicName := c.topic.Name()
	errorPipe := conduitUnit.errorConduitByName(topicName)
	fusingPipe := conduitUnit.fusingConduitByName(topicName)
	statisticsPipe := conduitUnit.statisticsConduitByName(topicName)
	for {
		select {
		case err := <-errorPipe:
			analysisResult := c.fuseAnalysis(err)
			if !analysisResult {
				break
			}
			actuatorNum := len(c.actuator)
			for i := 0; i < actuatorNum; i++ {
				fusingPipe <- 1
			}
		case statistics := <-statisticsPipe:
			fmt.Println(statistics)
		}
	}
}

func (c *controller) fuseAnalysis(err int) (analysisResult bool) {
	c.timeSliceErrorStatisticsLock.Lock()
	defer c.timeSliceErrorStatisticsLock.Unlock()
	timeFormat := time.Now().Format("2006-01-02 15:04")
	if _, ok := c.timeSliceErrorStatistics[timeFormat]; !ok {
		c.timeSliceErrorStatistics[timeFormat] = 1
		return
	}
	c.timeSliceErrorStatistics[timeFormat]++
	if c.timeSliceErrorStatistics[timeFormat] < c.topic.FuseSalt() {
		return
	}
	analysisResult = true
	return
}

func (c *controller) schedule() {

}

func (c *controller) start() (startResult bool) {
	c.createPipe()

	return
}

func (c *controller) stop() (stopResult bool) {
	return
}
