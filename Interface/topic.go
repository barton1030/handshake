package Interface

type Topic interface {
	Name() (name string)
	MinConcurrency() (minConcurrency int)
	MaxConcurrency() (maxConcurrency int)
	FuseSalt() (fuseSalt int)
	MaxRetryCount() (maxRetryCount int)
	CallbackHandler() (callback Callback)
	AlarmHandler() (alarm Alarm)
	MessageQueuingHandler() (messageQueuing MessageQueuing)
	Recipients() (recipients []interface{})
}

type Callback interface {
	Do(data map[string]interface{}) (res map[string]interface{}, err error)
}

type Alarm interface {
	Do(information map[string]interface{}, recipients []interface{})
}

type MessageQueuing interface {
	Pop() (message Message, err error)
	Push(message Message) (err error)
	Finish(message Message) (err error)
	Count() (count int)
}

type Message interface {
	Id() (id int)
	Data() (data map[string]interface{}, err error)
	RetryCount() (retryCont int)
	IncrRetryCont()
	Processable() (processable bool)
	Success()
	Fail()
}
