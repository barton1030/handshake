package persistent

import (
	"encoding/json"
	inter "handshake/Interface"
	"time"
)

type logDao struct {
	transactionId int
	tableName     string
}

var LogDao = logDao{
	tableName: "hand_shake_log",
}

func (l logDao) Add(log inter.Log) (err error) {
	log2 := l.transformation(log)
	err = transactionController.dbConn(l.transactionId).Table(l.tableName).Create(&log2).Error
	return err
}

func (l logDao) transformation(log inter.Log) (log2 storageLog) {
	log2.SId = log.Id()
	log2.SBusinessType = log.BusinessType()
	log2.SBusinessId = log.BusinessId()
	data, _ := json.Marshal(log.Data())
	log2.SData = string(data)
	log2.SCreator = log.Creator()
	log2.SCreateTime = log.CreateTime()
	return log2
}

type storageLog struct {
	SId           int       `json:"id" gorm:"column:id;primary_key"`
	SData         string    `json:"data" gorm:"column:data"`
	SBusinessType int       `json:"business_type" gorm:"column:business_type"`
	SBusinessId   int       `json:"business_id" gorm:"column:business_id"`
	SCreator      int       `json:"creator" gorm:"column:creator"`
	SCreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
}

func (l storageLog) Id() int {
	return l.SId
}

func (l storageLog) Data() map[string]interface{} {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(l.SData), &data)
	return data
}

func (l storageLog) BusinessType() int {
	return l.SBusinessType
}

func (l storageLog) BusinessId() int {
	return l.SBusinessId
}

func (l storageLog) Creator() int {
	return l.SCreator
}

func (l storageLog) CreateTime() time.Time {
	return l.SCreateTime
}
