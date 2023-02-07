package persistent

import (
	"errors"
	"github.com/jinzhu/gorm"
	"handshake/persistent/internal"
	"sync"
)

func Init() {
	internal.DbInit()
	internal.RedisInit()
}

func Close() {
	internal.CloseDb()
	internal.CloseRedis()
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

func (t *transaction) BeginTransaction() (transactionId int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	dbConn := internal.DbConn().Begin()
	t.transactionMap[t.nextId] = dbConn
	transactionId = t.nextId
	t.nextId++
	return transactionId
}

func (t *transaction) Transaction(transactionId int) *gorm.DB {
	t.lock.Lock()
	defer t.lock.Unlock()
	dbConn, ok := t.transactionMap[transactionId]
	if !ok {
		return internal.DbConn()
	}
	return dbConn
}

func (t *transaction) Commit(transactionId int) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	dbConn, ok := t.transactionMap[transactionId]
	if !ok {
		err := errors.New("事务不存在，请注意")
		return err
	}
	err := dbConn.Commit().Error
	if err != nil {
		return err
	}
	delete(t.transactionMap, transactionId)
	return err
}

func (t *transaction) Rollback(transactionId int) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	dbConn, ok := t.transactionMap[transactionId]
	if !ok {
		err := errors.New("事务不存在，请注意")
		return err
	}
	err := dbConn.Rollback().Error
	if err != nil {
		return err
	}
	delete(t.transactionMap, transactionId)
	return err
}

type base struct {
}

func (b base) BeginTransaction() (transactionId int) {
	transactionId = transactionController.BeginTransaction()
	return transactionId
}

func (b base) Transaction(transactionId int) *gorm.DB {
	dbConn := transactionController.Transaction(transactionId)
	return dbConn
}

func (b base) Commit(transactionId int) error {
	err := transactionController.Commit(transactionId)
	return err
}

func (b base) Rollback(transactionId int) error {
	err := transactionController.Rollback(transactionId)
	return err
}

func (b base) DbConn() *gorm.DB {
	dbConn := internal.DbConn()
	return dbConn
}
