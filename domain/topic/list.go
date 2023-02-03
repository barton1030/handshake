package topic

import (
	inter "handshake/Interface"
	"handshake/persistent"
)

type list struct {
	nextId  int
	storage inter.StorageTopicList
}

var List = list{nextId: 1, storage: persistent.TopicDao}

func (l *list) Add(topic2 topic) (err error) {
	topic2.id = l.nextId
	err = l.storage.Add(&topic2)
	return
}

func (l *list) Edit(topic2 topic) (err error) {
	err = l.storage.Edit(&topic2)
	return
}

func (l *list) Delete() (err error) {
	return
}

func (l *list) TopicId(topicId int) (topic2 topic, err error) {
	storageTopic, err := l.storage.TopicById(topicId)
	if err != nil {
		return
	}
	topic2 = l.reconstruction(storageTopic)
	return
}

func (l *list) TopicName(topicName string) (topic topic, err error) {
	return
}

func (l *list) reconstruction(topic2 inter.Topic) (topic3 topic) {
	topic3.id = topic2.Id()
	topic3.name = topic2.Name()
	topic3.status = topic2.Status()
	topic3.maxRetryCount = topic2.MaxRetryCount()
	topic3.minConcurrency = topic2.MinConcurrency()
	topic3.maxConcurrency = topic2.MaxConcurrency()
	topic3.fuseSalt = topic2.FuseSalt()
	topic3.queue = messageQueuing{
		topicName: topic2.Name(),
		storage:   make(map[int]interface{}),
	}
	topic2AlamHandler := topic2.AlarmHandler()
	topic3.alarm = alarm{
		url:        topic2AlamHandler.Url(),
		method:     topic2AlamHandler.Method(),
		recipients: topic2AlamHandler.Recipients(),
	}
	topic2CallbackHandler := topic2.CallbackHandler()
	topic3.callback = callback{
		url:     topic2CallbackHandler.Url(),
		method:  topic2CallbackHandler.Method(),
		headers: topic2CallbackHandler.Headers(),
		cookies: topic2CallbackHandler.Cookies(),
	}
	return
}
