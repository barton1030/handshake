package engine

import (
	inter "handshake/Interface"
	"time"
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
		message, err := a.topic.MessageQueuingHandler().Pop()
		if err != nil || message == nil || message.Id() <= 0 {
			time.Sleep(2 * time.Second)
			continue
		}
		toProcess := message.Processable()
		if !toProcess {
			continue
		}
		message.IncrRetryCont()

		res, err := a.topic.CallbackHandler().Do(message.Data())
		if err != nil {
			a.errorPipe <- 1
			a.alarm(err.Error(), message.Id())
			maxRetryCount := a.topic.MaxRetryCount()
			retryCount := message.RetryCount()
			if retryCount >= maxRetryCount {
				message.Fail()
				err = a.topic.MessageQueuingHandler().Finish(message)
			} else {
				err = a.topic.MessageQueuingHandler().Push(message)
			}
			if err != nil {
				a.alarm(err.Error(), message.Id())
			}
			continue
		}

		if code, ok := res["code"].(int); ok && code != 200 {
			a.errorPipe <- 1
			a.alarm(err.Error(), message.Id())
			maxRetryCount := a.topic.MaxRetryCount()
			retryCount := message.RetryCount()
			if retryCount >= maxRetryCount {
				message.Fail()
				err = a.topic.MessageQueuingHandler().Finish(message)
			} else {
				err = a.topic.MessageQueuingHandler().Push(message)
			}
			if err != nil {
				a.alarm(err.Error(), message.Id())
			}
			continue
		}
		message.Success()
		err = a.topic.MessageQueuingHandler().Finish(message)
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
	alarm.Do(information)
}
