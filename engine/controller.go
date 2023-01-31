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
	errorPipe                    chan int
	fusingPipe                   chan int
	statisticsPipe               chan int
}

func newController(topic inter.Topic) *controller {
	pipeName := topic.Name()
	pipeCap := topic.MaxConcurrency()
	errorPipe := conduitUnit.setUpErrorConduit(pipeName, pipeCap)
	fusingPipe := conduitUnit.setUpFusingConduit(pipeName, pipeCap)
	statisticsPipe := conduitUnit.setUpStatisticsConduit(pipeName, pipeCap)
	return &controller{
		topic:                    topic,
		actuator:                 make(map[string]*actuator),
		timeSliceErrorStatistics: make(map[string]int),
		errorPipe:                errorPipe,
		fusingPipe:               fusingPipe,
		statisticsPipe:           statisticsPipe,
	}
}

func (c *controller) monitorPipe() {
	for {
		select {
		case <-c.errorPipe:
			analysisResult := c.fuseAnalysis()
			if !analysisResult {
				break
			}
			actuatorNum := len(c.actuator)
			for i := 0; i < actuatorNum; i++ {
				c.fusingPipe <- 1
			}
		case statistics := <-c.statisticsPipe:
			fmt.Println(statistics)
		}
	}
}

func (c *controller) fuseAnalysis() (analysisResult bool) {
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

	return
}

func (c *controller) stop() (stopResult bool) {
	return
}
