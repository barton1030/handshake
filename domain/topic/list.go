package topic

import (
	inter "handshake/Interface"
	"handshake/persistent"
)

type List struct {
	storage inter.StorageTopicList
}

var ListExample = List{storage: persistent.TopicDao}

func (l *List) SetStorage(storageInter inter.StorageTopicList) *List {
	return &List{
		storage: storageInter,
	}
}

func (l *List) Add(topic2 topic) (err error) {
	err = l.storage.Add(&topic2)
	return
}

func (l *List) Edit(topic2 topic) (err error) {
	err = l.storage.Edit(&topic2)
	return
}

func (l *List) TopicId(topicId int) (topic2 topic, err error) {
	storageTopic, err := l.storage.TopicById(topicId)
	if err != nil {
		return
	}
	topic2 = l.reconstruction(storageTopic)
	return
}

func (l *List) TopicName(topicName string) (topic3 topic, err error) {
	topic2, err := l.storage.TopicByName(topicName)
	if err != nil {
		return
	}
	topic3 = l.reconstruction(topic2)
	return
}

func (l *List) ClapHisLockTopicByName(topicName string) (topic2 topic, err error) {
	storageTopic, err := l.storage.ClapHisLockTopicByName(topicName)
	if err != nil {
		return
	}
	topic2 = l.reconstruction(storageTopic)
	return
}

func (l *List) ClapHisLockTopicById(topicId int) (topic2 topic, err error) {
	storageTopic, err := l.storage.ClapHisLockTopicById(topicId)
	if err != nil {
		return
	}
	topic2 = l.reconstruction(storageTopic)
	return
}

func (l *List) List(startId, limit int) (topicList []topic, err error) {
	storageTopicList, err := l.storage.TopicList(startId, limit)
	if err != nil {
		return
	}
	if len(storageTopicList) <= 0 {
		return
	}

	for _, storageTopic := range storageTopicList {
		topic2 := l.reconstruction(storageTopic)
		topicList = append(topicList, topic2)
	}

	return
}

func (l *List) reconstruction(topic2 inter.Topic) (topic3 topic) {
	topic3.id = topic2.Id()
	topic3.name = topic2.Name()
	topic3.status = topic2.Status()
	topic3.maxRetryCount = topic2.MaxRetryCount()
	topic3.minConcurrency = topic2.MinConcurrency()
	topic3.maxConcurrency = topic2.MaxConcurrency()
	topic3.fuseSalt = topic2.FuseSalt()
	topic3.creator = topic2.Creator()
	topic3.queue = newMessageQueuing(topic2.Name())
	topic2AlamHandler := topic2.AlarmHandler()
	topic3.alarm = alarm{
		url:                topic2AlamHandler.Url(),
		method:             topic2AlamHandler.Method(),
		recipients:         topic2AlamHandler.Recipients(),
		cookies:            topic2AlamHandler.Cookies(),
		headers:            topic2AlamHandler.Headers(),
		templateParameters: topic2AlamHandler.TemplateParameters(),
	}
	topic2CallbackHandler := topic2.CallbackHandler()
	topic3.callback = callback{
		url:     topic2CallbackHandler.Url(),
		method:  topic2CallbackHandler.Method(),
		headers: topic2CallbackHandler.Headers(),
		cookies: topic2CallbackHandler.Cookies(),
	}
	topic3.createTime = topic2.CreateTime()
	return
}
