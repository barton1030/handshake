package persistent

import (
	"errors"
	"github.com/jinzhu/gorm"
	inter "handshake/Interface"
	"handshake/persistent/internal"
	"strconv"
	"sync"
	"time"
)

type manager struct {
	transactionId string
}

var Manager = manager{}

func (m manager) Begin() inter.StorageManager {
	m.transactionId = transactionController.begin()
	return m
}

func (m manager) dbConn() *gorm.DB {
	dbConn := transactionController.dbConn(m.transactionId)
	if dbConn == nil {
		return internal.DbConn()
	}
	return dbConn
}

func (m manager) Commit() error {
	dbConn := transactionController.dbConn(m.transactionId)
	if dbConn == nil {
		err := errors.New("事务不存在，请注意")
		return err
	}
	err := dbConn.Commit().Error
	if err != nil {
		return err
	}
	transactionController.delete(m.transactionId)
	return err
}

func (m manager) Rollback() error {
	dbConn := transactionController.dbConn(m.transactionId)
	if dbConn == nil {
		err := errors.New("事务不存在，请注意")
		return err
	}
	err := dbConn.Rollback().Error
	if err != nil {
		return err
	}
	transactionController.delete(m.transactionId)
	return err
}

func (m manager) UserDao() inter.StorageUserList {
	if len(m.transactionId) <= 0 {
		return UserDao
	}
	return userDao{
		transactionId: m.transactionId,
		tableName:     UserDao.tableName,
	}
}

func (m manager) TopicDao() inter.StorageTopicList {
	if len(m.transactionId) <= 0 {
		return TopicDao
	}
	return topicDao{
		transactionId: m.transactionId,
		tableName:     TopicDao.tableName,
	}
}

func (m manager) RoleDao() inter.StorageRoleList {
	if len(m.transactionId) <= 0 {
		return RoleDao
	}
	return roleDao{
		transactionId: m.transactionId,
		tableName:     RoleDao.tableName,
	}
}

func (m manager) QueueDao() inter.StorageQueueList {
	if len(m.transactionId) <= 0 {
		return QueueDao
	}
	return queueDao{
		transactionId: m.transactionId,
		tableName:     QueueDao.tableName,
	}
}

func (m manager) LogDao() inter.StorageLogList {
	if len(m.transactionId) <= 0 {
		return LogDao
	}
	return logDao{
		transactionId: m.transactionId,
		tableName:     LogDao.tableName,
	}
}

type transaction struct {
	currentDate    string
	nextId         int
	transactionMap map[string]*gorm.DB
	lock           sync.Mutex
}

var transactionController = transaction{
	currentDate:    time.Now().Format("2006-01-02"),
	nextId:         1,
	transactionMap: make(map[string]*gorm.DB),
}

func (t *transaction) begin() (transactionId string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	dbConn := internal.DbConn().Begin()
	// 事务标识生成规则
	currentData := time.Now().Format("2006-01-02")
	if t.currentDate != currentData {
		t.currentDate = currentData
		t.nextId = 1
	}
	currentDateNextId := strconv.Itoa(t.nextId)
	transactionId = t.currentDate + ":" + currentDateNextId

	t.transactionMap[transactionId] = dbConn
	t.nextId++
	return transactionId
}

func (t *transaction) dbConn(transactionId string) *gorm.DB {
	t.lock.Lock()
	defer t.lock.Unlock()
	dbConn, ok := t.transactionMap[transactionId]
	if !ok {
		return internal.DbConn()
	}
	return dbConn
}

func (t *transaction) delete(transactionId string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	delete(t.transactionMap, transactionId)
}
