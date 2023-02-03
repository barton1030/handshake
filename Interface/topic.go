package Interface

type StorageTopicList interface {
	Add(topic Topic) error
	Edit(topic Topic) error
	Delete(topic Topic) error
	TopicById(topicId int) (Topic, error)
	TopicByName(topicName string) (Topic, error)
}

type Topic interface {
	Id() (id int)
	Name() (name string)
	Status() (status int)
	MinConcurrency() (minConcurrency int)
	MaxConcurrency() (maxConcurrency int)
	FuseSalt() (fuseSalt int)
	MaxRetryCount() (maxRetryCount int)
	CallbackHandler() (callback Callback)
	AlarmHandler() (alarm Alarm)
	MessageQueuingHandler() (messageQueuing MessageQueuing)
	Recipients() (recipients []interface{})
	Creator() (creatorId int)
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
