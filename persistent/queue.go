package persistent

import (
	"encoding/json"
	inter "handshake/Interface"
	"handshake/persistent/internal"
	"strconv"
	"time"
)

type queueDao struct {
	transactionId int
	tableName     string
}

var QueueDao = queueDao{
	tableName: "hand_shake_queue_",
}

func (q queueDao) MaxPrimaryKeyId(topicName string) (maxPrimaryKeyId int) {
	tableName := q.tableName + topicName
	message2 := storageMessage{}
	err := transactionController.dbConn(q.transactionId).Table(tableName).Last(&message2).Error
	if err != nil {
		return
	}
	if message2.Id() <= 0 {
		return
	}
	maxPrimaryKeyId = message2.Id()
	return
}

func (q queueDao) Add(topicName string, message inter.Message) (err error) {
	tableName := q.tableName + topicName
	message2 := q.transformation(message)
	// 只推redis
	if message2.Id() > 0 {
		err = q.OnlyAdd(topicName, message2)
		return err
	}

	// create方法返回了持久后自增主键
	err = transactionController.dbConn(q.transactionId).Table(tableName).Create(&message2).Error
	if err != nil {
		return err
	}
	messageJson, err := json.Marshal(message2)
	if err != nil {
		return err
	}
	redisConn := internal.RedisConn().Get()
	defer redisConn.Close()
	_, err = redisConn.Do("LPUSH", tableName, string(messageJson))
	return err
}

func (q queueDao) OnlyAdd(topicName string, message storageMessage) (err error) {
	tableName := q.tableName + topicName
	err = q.Edit(topicName, message)
	if err != nil {
		return
	}
	messageJson, err := json.Marshal(message)
	if err != nil {
		return err
	}
	redisConn := internal.RedisConn().Get()
	defer redisConn.Close()
	_, err = redisConn.Do("LPUSH", tableName, string(messageJson))
	return err
}

func (q queueDao) NextPendingData(topicName string, offset int) (message inter.Message, err error) {
	redisConn := internal.RedisConn().Get()
	defer redisConn.Close()
	tableName := q.tableName + topicName
	resp, err := redisConn.Do("BRPOP", tableName, 1)
	if err != nil {
		return
	}
	message2 := storageMessage{}
	if res, ok := resp.([]interface{}); ok {
		if value, ok := res[1].([]byte); ok {
			json.Unmarshal(value, &message2)
		}
	}
	//if message2.Id() > 0 {
	//	message = message2
	//	return
	//}
	//whereId := strconv.Itoa(offset)
	//err = transactionController.dbConn(q.transactionId).Table(tableName).Where("id = ?", whereId).First(&message2).Error
	//if err != nil && err.Error() == "record not found" {
	//	err = nil
	//	return
	//}
	message = message2
	return
}

func (q queueDao) PendingDataCount(topicName string) (count int, err error) {
	redisConn := internal.RedisConn().Get()
	defer redisConn.Close()
	tableName := q.tableName + topicName
	pendingDataCount, err := redisConn.Do("LLEN", tableName)
	if err != nil {
		return
	}
	if dataCount, ok := pendingDataCount.(int64); ok {
		count = int(dataCount)
	}
	return
}

func (q queueDao) Edit(topicName string, message inter.Message) error {
	tableName := q.tableName + topicName
	message2 := q.transformation(message)
	whereId := strconv.Itoa(message2.Id())
	err := transactionController.dbConn(q.transactionId).Table(tableName).Where("id = ?", whereId).Updates(message2).Limit(1).Error
	return err
}

func (q queueDao) transformation(message inter.Message) (message2 storageMessage) {
	message2.SId = message.Id()
	message2.SStatus = message.Status()
	storageData, _ := json.Marshal(message.Data())
	message2.SData = string(storageData)
	message2.SRetry = message.RetryCount()
	message2.SCreateTime = message.CreateTime()
	return
}

type storageMessage struct {
	SId         int       `json:"id" gorm:"column:id;primary_key"`
	SData       string    `json:"data" gorm:"column:data"`
	SStatus     int       `json:"status" gorm:"column:status"`
	SRetry      int       `json:"retry" gorm:"column:retry"`
	SCreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

func (m storageMessage) Id() int {
	return m.SId
}

func (m storageMessage) Status() int {
	return m.SStatus
}

func (m storageMessage) RetryCount() int {
	return m.SRetry
}

func (m storageMessage) IncrRetryCont() {

}

func (m storageMessage) Data() map[string]interface{} {
	data := make(map[string]interface{})
	json.Unmarshal([]byte(m.SData), &data)
	return data
}

func (m storageMessage) Processable() (processable bool) {
	return
}

func (m storageMessage) Success() {

}

func (m storageMessage) Fail() {

}

func (m storageMessage) CreateTime() time.Time {
	return m.SCreateTime
}
