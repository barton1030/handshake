package engine

import inter "handshake/Interface"

const (
	ActuatorInitStatus    = 0
	ActuatorRunStatus     = 1
	ActuatorSuspendStatus = 2
	ActuatorExitStatus    = -1
)

type actuator struct {
	id     int
	status int8
	topic  inter.Topic
}

func newActuator(actuatorId int, topic inter.Topic) *actuator {
	return &actuator{
		id:     actuatorId,
		status: ActuatorInitStatus,
		topic:  topic,
	}
}

func (a actuator) start() {

}

func (a actuator) stop() {

}
