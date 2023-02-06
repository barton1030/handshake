package Interface

type StorageTopicList interface {
	MaxPrimaryKeyId() (maxPrimaryKeyId int)
	Add(topic Topic) error
	Edit(topic Topic) error
	TopicById(topicId int) (Topic, error)
	TopicByName(topicName string) (Topic, error)
}

type Topic interface {
	Id() int
	Name() string
	Status() int
	MinConcurrency() int
	MaxConcurrency() int
	FuseSalt() int
	MaxRetryCount() int
	CallbackHandler() Callback
	AlarmHandler() Alarm
	MessageQueuingHandler() (queue MessageQueuing)
	Creator() int
}

type Callback interface {
	Do(data map[string]interface{}) (res map[string]interface{}, err error)
	Headers() map[string]interface{}
	Cookies() map[string]interface{}
	Url() string
	Method() string
}

type Alarm interface {
	Do(information map[string]interface{}, recipients []interface{}) (res map[string]interface{}, err error)
	Url() string
	Method() string
	Recipients() []interface{}
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
