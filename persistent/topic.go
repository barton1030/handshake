package persistent

import (
	inter "handshake/Interface"
	"handshake/persistent/internal"
)

type topicDao struct {
	tableName string
}

var TopicDao = topicDao{
	tableName: "hand_shake_topic",
}

func (t topicDao) Add(topic inter.Topic) error {
	topic2 := t.transformation(topic)
	err := internal.DbConn().Table(t.tableName).Create(&topic2).Error
	return err
}

func (t topicDao) Edit(topic inter.Topic) error {
	topic2 := t.transformation(topic)
	var err = internal.DbConn().Table(t.tableName).Model(&struct {
		SId int
	}{SId: topic2.Id()}).Updates(topic2).Error
	return err
}

func (t topicDao) Delete(topic inter.Topic) error {
	topic2 := t.transformation(topic)
	err := internal.DbConn().Table(t.tableName).Delete(&struct {
		SId int
	}{SId: topic2.Id()}).Limit(1).Error
	return err
}

func (t topicDao) TopicById(topicId int) (inter.Topic, error) {
	topic := storageTopic{}
	err := internal.DbConn().Table(t.tableName).First(&topic, struct {
		SId int
	}{SId: topicId}).Error
	if err != nil && err.Error() == "record not found" {
		err = nil
	}
	return topic, err
}

func (t topicDao) TopicByName(topicName string) (inter.Topic, error) {
	topic := storageTopic{}
	err := internal.DbConn().Table(t.tableName).Where("name = ?", topicName).First(&topic).Error
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
	return topic2
}

type storageTopic struct {
	SId             int    `json:"s_id" gorm:"s_id"`
	SName           string `json:"s_name" gorm:"s_name"`
	SStatus         int    `json:"s_status" gorm:"s_status"`
	SMaxRetryCount  int    `json:"s_max_retry_count" gorm:"s_max_retry_count"`
	SMinConcurrency int    `json:"s_min_concurrency" gorm:"s_min_concurrency"`
	SMaxConcurrency int    `json:"s_max_concurrency" gorm:"s_max_concurrency"`
	SFuseSalt       int    `json:"s_fuse_salt" gorm:"s_fuse_salt"`
	SAlarm          string `json:"s_alarm" gorm:"s_alarm"`
	SCallback       string `json:"s_callback" gorm:"s_callback"`
	SCreator        int    `json:"s_creator" gorm:"s_creator"`
}

func (t storageTopic) Id() (id int) {
	id = t.SId
	return
}

func (t storageTopic) Name() (name string) {
	name = t.SName
	return
}

func (t storageTopic) Status() (status int) {
	status = t.SStatus
	return
}

func (t storageTopic) MinConcurrency() (minConcurrency int) {
	minConcurrency = t.SMinConcurrency
	return
}

func (t storageTopic) MaxConcurrency() (maxConcurrency int) {
	maxConcurrency = t.SMaxConcurrency
	return
}

func (t storageTopic) FuseSalt() (fuseSalt int) {
	return
}

func (t storageTopic) MaxRetryCount() (maxRetryCount int) {
	return
}

func (t storageTopic) CallbackHandler() (callback inter.Callback) {
	return
}

func (t storageTopic) AlarmHandler() (alarm inter.Alarm) {
	return
}

func (t storageTopic) MessageQueuingHandler() (messageQueuing inter.MessageQueuing) {
	return
}

func (t storageTopic) Recipients() (recipients []interface{}) {
	return
}

func (t storageTopic) Creator() (creatorId int) {
	return
}
