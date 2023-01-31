package topic

type list struct {
	nextId  int
	storage map[int]topic
}

var List = list{storage: make(map[int]topic)}

func (l *list) Add() (err error) {
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

func (l *list) TopicName(topicName int) (topic topic, err error) {
	return
}
