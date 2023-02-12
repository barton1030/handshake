package log

import (
	inter "handshake/Interface"
	"handshake/persistent"
)

type List struct {
	storage inter.StorageLogList
}

const (
	roleLog  = 1
	userLog  = 2
	topicLog = 3
)

var ListExample = List{storage: persistent.LogDao}

func (l List) SetStorage(storageInter inter.StorageLogList) List {
	return List{
		storage: storageInter,
	}
}

func (l List) AddRoleLog(log2 log) (err error) {
	log2.businessType = roleLog
	err = l.storage.Add(log2)
	return
}

func (l List) AddUserLog(log2 log) (err error) {
	log2.businessType = userLog
	err = l.storage.Add(log2)
	return
}

func (l List) AddTopicLog(log2 log) (err error) {
	log2.businessType = topicLog
	err = l.storage.Add(log2)
	return
}
