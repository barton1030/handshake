package engine

import (
	inter "handshake/Interface"
	"sync"
	"time"
)

// 控制单元负责执行器调度、错误执行数据统计、熔断保护计算触发、执行器管理等

const (
	SingleActuatorTaskVolume = 2000
	ControllerInitStatus     = 0
	ControllerRunStatus      = 1
	ControllerFusingStatus   = 2
	ControllerExitStatus     = -2
)

type controller struct {
	topic                    inter.Topic
	status                   int
	actuatorMap              map[int]*actuator
	timeSliceErrorStatistics map[string]int
	errorPipe                chan int
	fusingPipe               chan int
	statisticsPipe           chan int
	toBeExitSignal           chan int
	exitSignal               chan int
	nextActuatorId           int
	snapInLock               sync.Mutex
	lock                     sync.Mutex
}

func newController(topic inter.Topic) *controller {
	pipeName := topic.Name()
	pipeCap := topic.MaxConcurrency()
	errorPipe := conduitUnit.setUpErrorConduit(pipeName, pipeCap)
	fusingPipe := conduitUnit.setUpFusingConduit(pipeName, pipeCap)
	statisticsPipe := conduitUnit.setUpStatisticsConduit(pipeName, pipeCap)
	return &controller{
		topic:                    topic,
		status:                   ControllerInitStatus,
		actuatorMap:              make(map[int]*actuator),
		timeSliceErrorStatistics: make(map[string]int),
		errorPipe:                errorPipe,
		fusingPipe:               fusingPipe,
		statisticsPipe:           statisticsPipe,
		toBeExitSignal:           make(chan int),
		exitSignal:               make(chan int),
	}
}

func (c *controller) start() (startResult bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.status = ControllerRunStatus
	targetTaskNum := c.topic.MinConcurrency()
	c.actuatorSnapIn(targetTaskNum)
	go c.monitor()
	startResult = true
	return
}

func (c *controller) stop() (stopResult bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.status = ControllerExitStatus
	c.actuatorSnapIn(0)
	c.toBeExitSignal <- 1
	<-c.exitSignal
	stopResult = true
	return
}

func (c *controller) monitor() {
	for {
		select {
		case <-c.errorPipe:
			if c.status != ControllerRunStatus {
				break
			}
			analysisResult := c.fuseAnalysis()
			if !analysisResult {
				break
			}
			c.status = ControllerFusingStatus
			go func() {
				for _, cActuator := range c.actuatorMap {
					cActuator.suspend()
				}
			}()
		case <-c.statisticsPipe:
		case <-time.After(5 * time.Second):
			if c.status != ControllerRunStatus {
				break
			}
			// 统计队列数量并通过分析计算所需任务数量
			taskNum := c.queueCountAnalysis()
			c.actuatorSnapIn(taskNum)
		case <-c.toBeExitSignal:
			c.exitSignal <- 1
			return
		}
	}
}

// fuseAnalysis 计算规定时间片段内错误数据量，进行熔断计算保护
func (c *controller) fuseAnalysis() (analysisResult bool) {
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

// queueCountAnalysis 分析队列数据量，计算并发执行量
func (c *controller) queueCountAnalysis() (taskNum int) {
	topicMessageQueuing := c.topic.MessageQueuingHandler()
	messageDataCount := topicMessageQueuing.Count()
	taskNum = messageDataCount / SingleActuatorTaskVolume
	return
}

// actuatorSnapIn 执行器管理单元，统一管理执行器的新增和收缩操作
func (c *controller) actuatorSnapIn(targetActuatorNum int) {
	c.snapInLock.Lock()
	defer c.snapInLock.Unlock()
	if targetActuatorNum < 0 {
		return
	}
	// 最大并发数不能大于设置上限
	if targetActuatorNum > c.topic.MaxConcurrency() {
		targetActuatorNum = c.topic.MaxConcurrency()
	}
	currentActuatorNum := len(c.actuatorMap)
	if currentActuatorNum == targetActuatorNum {
		return
	}
	toIncr := true
	if currentActuatorNum > targetActuatorNum {
		toIncr = false
	}
	for {
		cActuatorNum := len(c.actuatorMap)
		if cActuatorNum == targetActuatorNum {
			return
		}
		if toIncr {
			cActuator := newActuator(c.nextActuatorId, c.topic)
			cActuator.start()
			c.actuatorMap[c.nextActuatorId] = cActuator
			c.nextActuatorId++
		} else {
			preActuatorId := c.nextActuatorId - 1
			cActuator := c.actuatorMap[preActuatorId]
			cActuator.stop()
			delete(c.actuatorMap, preActuatorId)
			c.nextActuatorId--
		}
	}
}
