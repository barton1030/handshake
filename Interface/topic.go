package Interface

type Topic interface {
	Name() string
	MinConcurrency() int
	MaxConcurrency() int
	FuseSalt() int
	CallbackHandler() Callback
	AlarmHandler() Alarm
	MessageQueuingHandler() MessageQueuing
}

type Callback interface {
	Do() (res map[string]interface{}, err error)
}

type Alarm interface {
	Do(information string, recipients []interface{})
}

type MessageQueuing interface {
	Pop() (message Message, err error)
	Push(message Message) (err error)
	Count() (count int)
}

type Message interface {
	Data() (data map[string]interface{}, err error)
	RetryCount() (retryCont int)
	IncrRetryCont()
	Processable() (processable bool)
	Success()
	Fail()
}
