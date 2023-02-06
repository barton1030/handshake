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
	startResult := topic3.Start()
	if !startResult {
		err = errors.New("主题启动失败")
		return err
	}
	err = topic2.List.Edit(topic3)
	return err
}

func (t topic) Stop(topicId int) error {
	topic3, err := topic2.List.TopicId(topicId)
	if err != nil {
		return err
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return err
	}
	stopResult := topic3.Stop()
	if !stopResult {
		err = errors.New("终止失败")
		return err
	}
	err = topic2.List.Edit(topic3)
	return err
}

func (t topic) Delete(topicId int) error {
	topic3, err := topic2.List.TopicId(topicId)
	if err != nil {
		return err
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return err
	}
	inOperation := topic3.InOperation()
	if inOperation {
		err = errors.New("主题运行中，请保证主题为停止态时执行删除操作！")
		return err
	}
	abandonmentResult := topic3.Abandonment()
	if !abandonmentResult {
		err = errors.New("废弃失败")
		return err
	}
	err = topic2.List.Edit(topic3)
	return err
}

func (t topic) SetCallback(topicId int, url, method string, headers, cookies map[string]interface{}) error {
	topic3, err := topic2.List.TopicId(topicId)
	if err != nil {
		return err
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return err
	}
	callback := topic2.NewCallBack(url, method, cookies, headers)
	topic3.SetCallback(callback)
	err = topic2.List.Edit(topic3)
	return err
}

func (t topic) SetAlarm(topicId int, url, method string, recipients []interface{}) error {
	topic3, err := topic2.List.TopicId(topicId)
	if err != nil {
		return err
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return err
	}
	alarm := topic2.NewAlarm(url, method, recipients)
	topic3.SetAlarm(alarm)
	err = topic2.List.Edit(topic3)
	return err
}

func (t topic) Edit(topicId, maxRetryCount, minConcurrency, maxConcurrency, fuseSalt int) error {
	topic3, err := topic2.List.TopicId(topicId)
	if err != nil {
		return err
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return err
	}
	topic3.SetFuseSalt(fuseSalt)
	topic3.SetMaxRetryCount(maxRetryCount)
	topic3.SetMinConcurrency(minConcurrency)
	topic3.SetMaxConcurrency(maxConcurrency)
	err = topic2.List.Edit(topic3)
	return err
}

func (t topic) TopicById(topicId int) (map[string]interface{}, error) {
	topic4 := make(map[string]interface{})
	topic3, err := topic2.List.TopicId(topicId)
	if err != nil {
		return topic4, err
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return topic4, err
	}

	topic4["id"] = topic3.Id()
	topic4["name"] = topic3.Name()
	topic4["status"] = topic3.Status()
	topic4["maxRetryCount"] = topic3.MaxRetryCount()
	topic4["minConcurrency"] = topic3.MinConcurrency()
	topic4["maxConcurrency"] = topic3.MaxConcurrency()
	topic4["fuseSalt"] = topic3.FuseSalt()
	alamHandler := topic3.AlarmHandler()
	alarm := make(map[string]interface{})
	alarm["url"] = alamHandler.Url()
	alarm["method"] = alamHandler.Method()
	alarm["recipients"] = alamHandler.Recipients()
	topic4["alarm"] = alarm
	callbackHandler := topic3.CallbackHandler()
	callback := make(map[string]interface{})
	callback["url"] = callbackHandler.Url()
	callback["method"] = callbackHandler.Method()
	callback["headers"] = callbackHandler.Headers()
	callback["cookies"] = callbackHandler.Cookies()
	topic4["callback"] = callback

	return topic4, err
}
