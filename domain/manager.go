package domain

import (
	inter "handshake/Interface"
	"handshake/domain/log"
	"handshake/domain/role"
	"handshake/domain/topic"
	"handshake/domain/user"
	"handshake/persistent"
)

type manager struct {
	storageManager inter.StorageManager
}

var Manager = manager{
	storageManager: persistent.Manager,
}

func (m manager) Begin() manager {
	return manager{
		storageManager: m.storageManager.Begin(),
	}
}

func (m manager) Commit() error {
	err := m.storageManager.Commit()
	return err
}

func (m manager) Rollback() error {
	err := m.storageManager.Rollback()
	return err
}

func (m manager) RoleList() *role.List {
	roleStorageInter := m.storageManager.RoleDao()
	list := role.ListExample.SetStorage(roleStorageInter)
	return list
}

func (m manager) UserList() *user.List {
	roleStorageInter := m.storageManager.UserDao()
	list := user.ListExample.SetStorage(roleStorageInter)
	return list
}

func (m manager) TopicList() *topic.List {
	roleStorageInter := m.storageManager.TopicDao()
	list := topic.ListExample.SetStorage(roleStorageInter)
	return list
}

func (m manager) LogList() log.List {
	roleStorageInter := m.storageManager.LogDao()
	list := log.ListExample.SetStorage(roleStorageInter)
	return list
}
