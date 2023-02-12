package persistent

import (
	"errors"
	"github.com/jinzhu/gorm"
	inter "handshake/Interface"
	"handshake/persistent/internal"
	"sync"
)

type manager struct {
	transactionId int
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
	if m.transactionId <= 0 {
		return UserDao
	}
	return userDao{
		transactionId: m.transactionId,
		tableName:     UserDao.tableName,
	}
}

func (m manager) TopicDao() inter.StorageTopicList {
	if m.transactionId <= 0 {
		return TopicDao
	}
	return topicDao{
		transactionId: m.transactionId,
		tableName:     TopicDao.tableName,
	}
}

func (m manager) RoleDao() inter.StorageRoleList {
	if m.transactionId <= 0 {
		return RoleDao
	}
	return roleDao{
		transactionId: m.transactionId,
		tableName:     RoleDao.tableName,
	}
}

func (m manager) QueueDao() inter.StorageQueueList {
	if m.transactionId <= 0 {
		return QueueDao
	}
	return queueDao{
		transactionId: m.transactionId,
		tableName:     QueueDao.tableName,
	}
}

func (m manager) LogDao() inter.StorageLogList {
	if m.transactionId <= 0 {
		return LogDao
	}
	return logDao{
		transactionId: m.transactionId,
		tableName:     LogDao.tableName,
	}
}

type transaction struct {
	nextId         int
	transactionMap map[int]*gorm.DB
	lock           sync.Mutex
}

var transactionController = transaction{
	nextId:         1,
	transactionMap: make(map[int]*gorm.DB),
}

func (t *transaction) begin() (transactionId int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	dbConn := internal.DbConn().Begin()
	t.transactionMap[t.nextId] = dbConn
	transactionId = t.nextId
	t.nextId++
	return transactionId
}

func (t *transaction) dbConn(transactionId int) *gorm.DB {
	t.lock.Lock()
	defer t.lock.Unlock()
	dbConn, ok := t.transactionMap[transactionId]
	if !ok {
		return internal.DbConn()
	}
	return dbConn
}

func (t *transaction) delete(transactionId int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	delete(t.transactionMap, transactionId)
}
