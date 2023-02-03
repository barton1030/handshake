package service

import (
	"errors"
	topic2 "handshake/domain/topic"
)

type topic struct {
}

var TopicService topic

func (t topic) Add(name string, maxRetryCount, minConcurrency, maxConcurrency, fuseSalt, creator int) error {
	topic3, err := topic2.List.TopicName(name)
	if err != nil {
		return err
	}
	if topic3.Id() > 0 {
		err = errors.New("主题名重复，请注意！")
		return err
	}
	topic4 := topic2.NewTopic(name, maxRetryCount, minConcurrency, maxConcurrency, fuseSalt, creator)
	err = topic2.List.Add(topic4)
	return err
}

func (t topic) Start(topicId int) error {
	topic3, err := topic2.List.TopicId(topicId)
	if err != nil {
		return err
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return err
	}
	err = topic3.Start()
	if err != nil {
		return err
	}
	err = topic2.List.Edit(topic3)
	return err
}
