package Interface

import "time"

type StorageTopicList interface {
	MaxPrimaryKeyId() (maxPrimaryKeyId int)
	Add(topic Topic) error
	Edit(topic Topic) error
	TopicById(topicId int) (Topic, error)
	TopicByName(topicName string) (Topic, error)
	TopicList(startId, limit int) ([]Topic, error)
	ClapHisLockTopicById(topicId int) (Topic, error)
	ClapHisLockTopicByName(topicName string) (Topic, error)
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
	CreateTime() time.Time
}

type Callback interface {
	Do(data map[string]interface{}) (res map[string]interface{}, err error)
	Headers() map[string]interface{}
	Cookies() map[string]interface{}
	Url() string
	Method() string
}

type Alarm interface {
	Do(information map[string]interface{}) (res map[string]interface{}, err error)
	Url() string
	Method() string
	Cookies() map[string]interface{}
	Headers() map[string]interface{}
	Recipients() map[int]int
	TemplateParameters() map[string]interface{}
}

type StorageQueueList interface {
	MaxPrimaryKeyId(topicName string) (maxPrimaryKeyId int)
	Add(topicName string, message Message) error
	Edit(topicName string, message Message) error
	NextPendingData(topicName string) (message Message, err error)
	PendingDataCount(topicName string) (count int, err error)
}

type MessageQueuing interface {
	Pop() (message Message, err error)
	Push(message Message) (err error)
	Finish(message Message) (err error)
	Count() (count int)
}

type Message interface {
	Id() int
	Status() int
	Data() map[string]interface{}
	RetryCount() int
	IncrRetryCont()
	Processable() (processable bool)
	Success()
	Fail()
	CreateTime() time.Time
}
