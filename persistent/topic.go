package persistent

import (
	"encoding/json"
	inter "handshake/Interface"
	"strconv"
)

type topicDao struct {
	base
	tableName string
}

var TopicDao = topicDao{
	tableName: "hand_shake_topic",
}

func (t topicDao) MaxPrimaryKeyId() (maxPrimaryKeyId int) {
	topic2 := storageTopic{}
	err := t.DbConn().Table(t.tableName).Last(&topic2).Error
	if err != nil {
		return
	}
	if topic2.Id() <= 0 {
		return
	}
	maxPrimaryKeyId = topic2.Id()
	return
}

func (t topicDao) Add(topic inter.Topic) error {
	topic2 := t.transformation(topic)
	err := t.DbConn().Table(t.tableName).Create(&topic2).Error
	return err
}

func (t topicDao) Edit(topic inter.Topic) error {
	topic2 := t.transformation(topic)
	err := t.DbConn().Table(t.tableName).Model(&storageTopic{SId: topic2.SId}).Updates(topic2).Limit(1).Error
	return err
}

func (t topicDao) TopicById(topicId int) (inter.Topic, error) {
	topic := storageTopic{}
	whereTopicId := strconv.Itoa(topicId)
	err := t.DbConn().Table(t.tableName).Where("id = ?", whereTopicId).First(&topic).Error
	if err != nil && err.Error() == "record not found" {
		err = nil
	}
	return topic, err
}

func (t topicDao) TopicByName(topicName string) (inter.Topic, error) {
	topic := storageTopic{}
	err := t.DbConn().Table(t.tableName).Where("name = ?", topicName).First(&topic).Error
	if err != nil && err.Error() == "record not found" {
		err = nil
	}
	return topic, err
}

func (t topicDao) transformation(topic inter.Topic) (topic2 storageTopic) {
	topic2.SId = topic.Id()
	topic2.SName = topic.Name()
	topic2.SStatus = topic.Status()
	topic2.SMinConcurrency = topic.MinConcurrency()
	topic2.SMaxConcurrency = topic.MaxConcurrency()
	topic2.SMaxRetryCount = topic.MaxRetryCount()
	topic2.SFuseSalt = topic.FuseSalt()
	topic2.SCreator = topic.Creator()
	domainCallback := topic.CallbackHandler()
	callback := storageCallback{
		SUrl:     domainCallback.Url(),
		SMethod:  domainCallback.Method(),
		SHeaders: domainCallback.Headers(),
		SCookies: domainCallback.Cookies(),
	}
	callbackJson, _ := json.Marshal(callback)
	topic2.SCallback = string(callbackJson)
	domainAlarm := topic.AlarmHandler()
	alarm := storageAlarm{
		SUrl:        domainAlarm.Url(),
		SMethod:     domainAlarm.Method(),
		SRecipients: domainAlarm.Recipients(),
	}
	alarmJson, _ := json.Marshal(alarm)
	topic2.SAlarm = string(alarmJson)
	return topic2
}

type storageTopic struct {
	SId             int    `json:"id" gorm:"column:id;primary_key"`
	SName           string `json:"name" gorm:"column:name"`
	SStatus         int    `json:"status" gorm:"column:status"`
	SMaxRetryCount  int    `json:"max_retry_count" gorm:"column:max_retry_count"`
	SMinConcurrency int    `json:"min_concurrency" gorm:"column:min_concurrency"`
	SMaxConcurrency int    `json:"max_concurrency" gorm:"column:max_concurrency"`
	SFuseSalt       int    `json:"fuse_salt" gorm:"column:fuse_salt"`
	SAlarm          string `json:"alarm" gorm:"column:alarm"`
	SCallback       string `json:"callback" gorm:"column:callback"`
	SCreator        int    `json:"creator" gorm:"column:creator"`
}

func (t storageTopic) Id() int {
	return t.SId
}

func (t storageTopic) Name() string {
	return t.SName
}

func (t storageTopic) Status() int {
	return t.SStatus
}

func (t storageTopic) MinConcurrency() int {
	return t.SMinConcurrency
}

func (t storageTopic) MaxConcurrency() int {
	return t.SMaxConcurrency
}

func (t storageTopic) FuseSalt() int {
	return t.SFuseSalt
}

func (t storageTopic) MaxRetryCount() int {
	return t.SMaxRetryCount
}

func (t storageTopic) CallbackHandler() inter.Callback {
	callback := storageCallback{}
	json.Unmarshal([]byte(t.SCallback), &callback)
	return callback
}

func (t storageTopic) AlarmHandler() inter.Alarm {
	alarm := storageAlarm{}
	json.Unmarshal([]byte(t.SAlarm), &alarm)
	return alarm
}

func (t storageTopic) MessageQueuingHandler() (queue inter.MessageQueuing) {
	return
}

func (t storageTopic) Creator() int {
	return t.SCreator
}

type storageCallback struct {
	SUrl     string                 `json:"s_url" gorm:"s_url"`
	SMethod  string                 `json:"s_method" gorm:"s_method"`
	SHeaders map[string]interface{} `json:"s_headers" gorm:"s_headers"`
	SCookies map[string]interface{} `json:"s_cookies" gorm:"s_cookies"`
}

func (c storageCallback) Do(data map[string]interface{}) (res map[string]interface{}, err error) {
	return
}

func (c storageCallback) Headers() map[string]interface{} {
	return c.SHeaders
}

func (c storageCallback) Cookies() map[string]interface{} {
	return c.SCookies
}

func (c storageCallback) Url() string {
	return c.SUrl
}

func (c storageCallback) Method() string {
	return c.SMethod
}

type storageAlarm struct {
	SUrl        string        `json:"s_url" gorm:"s_url"`
	SMethod     string        `json:"s_method" gorm:"s_method"`
	SRecipients []interface{} `json:"s_recipients" gorm:"s_recipients"`
}

func (a storageAlarm) Do(information map[string]interface{}, recipients []interface{}) (res map[string]interface{}, err error) {
	return
}

func (a storageAlarm) Url() string {
	return a.SUrl
}

func (a storageAlarm) Method() string {
	return a.SMethod
}

func (a storageAlarm) Recipients() []interface{} {
	return a.SRecipients
}
