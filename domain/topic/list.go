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

func (l *list) Edit() (err error) {
	return
}

func (l *list) Delete() (err error) {
	return
}

func (l *list) TopicId(topicId int) (topic topic, err error) {
	return
}

func (l *list) TopicName(topicName string) (topic topic, err error) {
	return
}
