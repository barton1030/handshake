package service

import (
	"errors"
	inter "handshake/Interface"
	"handshake/domain"
	"handshake/domain/log"
	topic2 "handshake/domain/topic"
	"time"
)

type topic struct {
}

var Topic topic

func (t topic) Add(operator int, name string, maxRetryCount, minConcurrency, maxConcurrency, fuseSalt int) (err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	begin := domain.Manager.Begin()
	topic3, err := begin.TopicList().ClapHisLockTopicByName(name)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if topic3.Id() > 0 {
		_ = begin.Rollback()
		err = errors.New("主题名重复，请注意！")
		return
	}
	topic4 := topic2.NewTopic(name, maxRetryCount, minConcurrency, maxConcurrency, fuseSalt, operator)
	err = begin.TopicList().Add(topic4)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	topicLogData := t.reconstruction(&topic4)
	topicLog := log.NewLog(topicLogData, topic4.Id(), operator, topic4.CreateTime())
	err = begin.LogList().AddTopicLog(topicLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()

	return
}

func (t topic) Start(operator, topicId int) (err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	begin := domain.Manager.Begin()
	topic3, err := begin.TopicList().ClapHisLockTopicById(topicId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if topic3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("主题不存在，请确认！")
		return
	}
	if topic3.DiscardOrNot() {
		_ = begin.Rollback()
		err = errors.New("当前主题已废弃, 请注意！")
		return
	}
	if topic3.Creator() != operator {
		_ = begin.Rollback()
		err = errors.New("操作人与主题创建者不一致，请确认！")
		return
	}
	startResult := topic3.Start()
	if !startResult {
		_ = begin.Rollback()
		err = errors.New("主题启动失败")
		return
	}
	err = begin.TopicList().Edit(topic3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	topicLogData := t.reconstruction(&topic3)
	topicLog := log.NewLog(topicLogData, topic3.Id(), operator, time.Now())
	err = begin.LogList().AddTopicLog(topicLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	startUpResult := topic3.StartUp()
	if !startUpResult {
		_ = begin.Rollback()
		return
	}
	err = begin.Commit()
	return
}

func (t topic) Stop(operator, topicId int) (err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	begin := domain.Manager.Begin()
	topic3, err := begin.TopicList().ClapHisLockTopicById(topicId)
	if err != nil {
		return err
	}
	if topic3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("主题不存在，请确认！")
		return err
	}
	if topic3.DiscardOrNot() {
		_ = begin.Rollback()
		err = errors.New("当前主题已废弃, 请注意！")
		return
	}
	if topic3.Creator() != operator {
		_ = begin.Rollback()
		err = errors.New("操作人与主题创建者不一致，请确认！")
		return
	}
	stopResult := topic3.Stop()
	if !stopResult {
		_ = begin.Rollback()
		err = errors.New("终止失败")
		return err
	}
	err = begin.TopicList().Edit(topic3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	topicLogData := t.reconstruction(&topic3)
	topicLog := log.NewLog(topicLogData, topic3.Id(), operator, time.Now())
	err = begin.LogList().AddTopicLog(topicLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	stopUpResult := topic3.StopUp()
	if !stopUpResult {
		_ = begin.Rollback()
		return
	}
	err = begin.Commit()
	return err
}

func (t topic) Delete(operator, topicId int) (err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	begin := domain.Manager.Begin()
	topic3, err := begin.TopicList().ClapHisLockTopicById(topicId)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	if topic3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("主题不存在，请确认！")
		return err
	}
	if topic3.Creator() != operator {
		_ = begin.Rollback()
		err = errors.New("操作人与主题创建者不一致，请确认！")
		return
	}
	inOperation := topic3.InOperation()
	if inOperation {
		_ = begin.Rollback()
		err = errors.New("主题运行中，请保证主题为停止态时执行删除操作！")
		return err
	}
	abandonmentResult := topic3.Abandonment()
	if !abandonmentResult {
		_ = begin.Rollback()
		err = errors.New("废弃失败, 请确认主题是否已停止！")
		return err
	}
	err = begin.TopicList().Edit(topic3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	topicLogData := t.reconstruction(&topic3)
	topicLog := log.NewLog(topicLogData, topic3.Id(), operator, time.Now())
	err = begin.LogList().AddTopicLog(topicLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return err
}

func (t topic) SetCallback(operator, topicId int, url, method string, headers, cookies map[string]interface{}) (err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	begin := domain.Manager.Begin()
	topic3, err := begin.TopicList().ClapHisLockTopicById(topicId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if topic3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("主题不存在，请确认！")
		return
	}
	if topic3.Creator() != operator {
		_ = begin.Rollback()
		err = errors.New("操作人与主题创建者不一致，请确认！")
		return
	}
	callback := topic2.NewCallBack(url, method, cookies, headers)
	topic3.SetCallback(callback)
	err = begin.TopicList().Edit(topic3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	topicLogData := t.reconstruction(&topic3)
	topicLog := log.NewLog(topicLogData, topic3.Id(), operator, time.Now())
	err = begin.LogList().AddTopicLog(topicLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return
}

func (t topic) SetAlarm(operator, topicId int, url, method string, recipients map[int]int, headers, cookies, templateParameters map[string]interface{}) (err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	begin := domain.Manager.Begin()
	topic3, err := begin.TopicList().ClapHisLockTopicById(topicId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if topic3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("主题不存在，请确认！")
		return
	}
	if topic3.Creator() != operator {
		_ = begin.Rollback()
		err = errors.New("操作人与主题创建者不一致，请确认！")
		return
	}
	alarm := topic2.NewAlarm(url, method, recipients, headers, cookies, templateParameters)
	topic3.SetAlarm(alarm)
	err = begin.TopicList().Edit(topic3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	topicLogData := t.reconstruction(&topic3)
	topicLog := log.NewLog(topicLogData, topic3.Id(), operator, time.Now())
	err = begin.LogList().AddTopicLog(topicLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return
}

func (t topic) Edit(operator, topicId, maxRetryCount, minConcurrency, maxConcurrency, fuseSalt int) (err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	begin := domain.Manager.Begin()
	topic3, err := begin.TopicList().ClapHisLockTopicById(topicId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if topic3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("主题不存在，请确认！")
		return
	}
	if topic3.Creator() != operator {
		_ = begin.Rollback()
		err = errors.New("操作人与主题创建者不一致，请确认！")
		return
	}
	topic3.SetFuseSalt(fuseSalt)
	topic3.SetMaxRetryCount(maxRetryCount)
	topic3.SetMinConcurrency(minConcurrency)
	topic3.SetMaxConcurrency(maxConcurrency)
	err = begin.TopicList().Edit(topic3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	topicLogData := t.reconstruction(&topic3)
	topicLog := log.NewLog(topicLogData, topic3.Id(), operator, time.Now())
	err = begin.LogList().AddTopicLog(topicLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return
}

func (t topic) TopicById(operator, topicId int) (topic4 map[string]interface{}, err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	topic3, err := domain.Manager.TopicList().TopicId(topicId)
	if err != nil {
		return topic4, err
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return topic4, err
	}
	topic4 = t.reconstruction(&topic3)
	return topic4, err
}

func (t topic) TopicByName(operator int, topicName string) (topic4 map[string]interface{}, err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	topic3, err := domain.Manager.TopicList().TopicName(topicName)
	if err != nil {
		return topic4, err
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return topic4, err
	}
	topic4 = t.reconstruction(&topic3)
	return topic4, err
}

func (t topic) TopicList(operator, startId, limit int) (list []map[string]interface{}, err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	topics, err := domain.Manager.TopicList().List(startId, limit)
	if err != nil {
		return list, err
	}

	topicNum := len(topics)
	list = make([]map[string]interface{}, topicNum, topicNum)
	for index, topic3 := range topics {
		topic4 := t.reconstruction(&topic3)
		list[index] = topic4
	}
	return list, err
}

func (t topic) PushMessage(operator, topicId int, message map[string]interface{}) (err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者用户不存在，请注意！")
		return
	}
	topic3, err := domain.Manager.TopicList().TopicId(topicId)
	if err != nil {
		return
	}
	if topic3.Id() <= 0 {
		err = errors.New("主题不存在，请确认！")
		return
	}
	if topic3.Creator() != operator {
		err = errors.New("操作人与主题创建者不一致，请确认！")
		return
	}
	data := topic2.NewMessage(message)
	queueHandler := topic3.MessageQueuingHandler()
	err = queueHandler.Push(&data)
	return
}

func (t topic) reconstruction(topic3 inter.Topic) (topic4 map[string]interface{}) {
	topic4 = make(map[string]interface{})
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
	alarm["headers"] = alamHandler.Headers()
	alarm["cookies"] = alamHandler.Cookies()
	alarm["templateParameters"] = alamHandler.TemplateParameters()
	topic4["alarm"] = alarm
	callbackHandler := topic3.CallbackHandler()
	callback := make(map[string]interface{})
	callback["url"] = callbackHandler.Url()
	callback["method"] = callbackHandler.Method()
	callback["headers"] = callbackHandler.Headers()
	callback["cookies"] = callbackHandler.Cookies()
	topic4["callback"] = callback
	topic4["createTime"] = topic3.CreateTime()
	return topic4
}
