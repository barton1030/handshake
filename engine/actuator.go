package engine

import (
	inter "handshake/Interface"
)

const (
	ActuatorInitStatus        = 0
	ActuatorRunStatus         = 1
	ActuatorToBeSuspendStatus = 2
	ActuatorSuspendStatus     = 3
	ActuatorExitStatus        = -1
)

type actuator struct {
	id             int
	status         int8
	topic          inter.Topic
	startSignal    chan int
	exitSignal     chan int
	suspendSignal  chan int
	errorPipe      chan int
	fusingPipe     chan int
	statisticsPipe chan int
}

func newActuator(actuatorId int, topic inter.Topic) *actuator {
	pipeName := topic.Name()
	errorPipe := conduitUnit.errorConduitByName(pipeName)
	fusingPipe := conduitUnit.fusingConduitByName(pipeName)
	statisticsPipe := conduitUnit.statisticsConduitByName(pipeName)
	return &actuator{
		id:             actuatorId,
		status:         ActuatorInitStatus,
		topic:          topic,
		startSignal:    make(chan int),
		exitSignal:     make(chan int),
		suspendSignal:  make(chan int),
		errorPipe:      errorPipe,
		fusingPipe:     fusingPipe,
		statisticsPipe: statisticsPipe,
	}
}

func (a *actuator) restart() {
	a.status = ActuatorInitStatus
	<-a.startSignal
}

func (a *actuator) start() {
	go a.implement()
	<-a.startSignal
}

func (a *actuator) stop() {
	a.status = ActuatorExitStatus
	<-a.exitSignal
	close(a.startSignal)
	close(a.suspendSignal)
	close(a.exitSignal)
}

func (a *actuator) suspend() {
	if a.status != ActuatorRunStatus {
		return
	}
	a.status = ActuatorToBeSuspendStatus
	<-a.suspendSignal
}

func (a *actuator) implement() {
	for {
		if a.status == ActuatorInitStatus {
			a.startSignal <- 1
			a.status = ActuatorRunStatus
		}
		if a.status == ActuatorToBeSuspendStatus {
			a.status = ActuatorSuspendStatus
			a.suspendSignal <- 1
		}
		if a.status == ActuatorSuspendStatus {
			continue
		}
		if a.status == ActuatorExitStatus {
			a.exitSignal <- 1
			return
		}
		// 消息队列具柄
		messageQueuing := a.topic.MessageQueuingHandler()
		message, err := messageQueuing.Pop()
		if err != nil {
			continue
		}
		data, err := message.Data()
		if err != nil {
			continue
		}
		// 消息回调处理工具具柄
		callback := a.topic.CallbackHandler()
		res, err := callback.Do(data)
		message.IncrRetryCont()
		if err != nil {
			a.alarm(err.Error(), message.Id())
			err = a.handleFail(message)
			if err != nil {
				a.alarm(err.Error(), message.Id())
			}
			a.errorPipe <- 1
			continue
		}
		if code, ok := res["code"]; ok && code != 0 {
			a.alarm(res["err"], message.Id())
			err = a.handleFail(message)
			if err != nil {
				a.alarm(err.Error(), message.Id())
			}
			a.errorPipe <- 1
			continue
		}
		message.Success()
		err = messageQueuing.Finish(message)
		if err != nil {
			a.alarm(err.Error(), message.Id())
		}
	}
}

func (a *actuator) alarm(err interface{}, messageId int) {
	alarm := a.topic.AlarmHandler()
	information := make(map[string]interface{})
	information["topic"] = a.topic.Name()
	information["messageId"] = messageId
	information["err"] = err
	recipients := a.topic.AlarmHandler().Recipients()
	alarm.Do(information, recipients)
}

func (a *actuator) handleFail(message inter.Message) (err error) {
	maxRetryCount := a.topic.MaxRetryCount()
	retryCount := message.RetryCount()
	messageQueuing := a.topic.MessageQueuingHandler()
	if retryCount < maxRetryCount {
		err = messageQueuing.Push(message)
		return
	}
	message.Fail()
	err = messageQueuing.Finish(message)
	return
}
