package engine

import (
	"fmt"
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
	toStartSignal            chan int
	startSignal              chan int
	toExitSignal             chan int
	exitSignal               chan int
	fuseTerminationSignal    chan int
	nextActuatorId           int
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
		toStartSignal:            make(chan int, 1),
		startSignal:              make(chan int, 1),
		toExitSignal:             make(chan int, 1),
		exitSignal:               make(chan int, 1),
		fuseTerminationSignal:    make(chan int, 1),
	}
}

func (c *controller) start() (startResult bool) {
	go c.monitor()
	c.toStartSignal <- 1
	<-c.startSignal
	startResult = true
	return
}

func (c *controller) stop() (stopResult bool) {
	c.toExitSignal <- 1
	<-c.exitSignal
	stopResult = true
	return
}

// monitor 控制器监听单元
func (c *controller) monitor() {
	for {
		select {
		case <-c.toStartSignal:
			c.init()
		case <-c.toExitSignal:
			c.exit()
			return
		case <-c.errorPipe:
			c.errorPipeProcessor()
		case <-c.statisticsPipe:
		case <-time.After(5 * time.Second):
			c.queueCountProcessor()
		case <-c.fuseTerminationSignal:
			c.fuseTerminationProcessor()
		}
	}
}

// init 初始化启动逻辑
func (c *controller) init() {
	if c.status != ControllerInitStatus {
		return
	}
	c.status = ControllerRunStatus
	targetTaskNum := c.topic.MinConcurrency()
	c.actuatorSnapIn(targetTaskNum)
	c.startSignal <- 1
}

// exit 退出逻辑
func (c *controller) exit() {
	if c.status == ControllerExitStatus {
		return
	}
	c.status = ControllerExitStatus
	c.actuatorSnapIn(0)
	c.exitSignal <- 1
	c.clearPipe()
}

// errorPipeProcessor 错误信息通信管道处理逻辑
func (c *controller) errorPipeProcessor() {
	if c.status != ControllerRunStatus {
		return
	}
	analysisResult := c.fuseAnalysis()
	if !analysisResult {
		return
	}
	c.status = ControllerFusingStatus
	for _, cActuator := range c.actuatorMap {
		cActuator.suspend()
	}
	go func() {
		defer func() {
			err := recover()
			fmt.Println(err)
		}()
		time.Sleep(10 * time.Second)
		c.fuseTerminationSignal <- 1
	}()
}

// queueCountProcessor 消息队列统计处理逻辑
func (c *controller) queueCountProcessor() {
	if c.status != ControllerRunStatus {
		return
	}
	// 统计队列数量并通过分析计算所需任务数量
	taskNum := c.queueCountAnalysis()
	if taskNum <= c.topic.MinConcurrency() {
		return
	}
	c.actuatorSnapIn(taskNum)
}

// clearPipe 关闭不同执行体和自身通信管道或channel
func (c *controller) clearPipe() {
	pipeName := c.topic.Name()
	conduitUnit.closeErrorConduit(pipeName)
	conduitUnit.closeFusingConduit(pipeName)
	conduitUnit.closeStatisticsConduit(pipeName)
}

// fuseTerminationProcessor 熔断接触逻辑
func (c *controller) fuseTerminationProcessor() {
	if c.status != ControllerFusingStatus {
		return
	}
	for _, cActuator := range c.actuatorMap {
		cActuator.restart()
	}
	c.status = ControllerRunStatus
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
	c.timeSliceErrorStatistics[timeFormat] = 0
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
