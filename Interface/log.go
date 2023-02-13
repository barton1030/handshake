package Interface

import "time"

type StorageLogList interface {
	Add(log Log) error
}

type Log interface {
	Id() int
	Data() map[string]interface{}
	BusinessType() int
	BusinessId() int
	Creator() int
	CreateTime() time.Time
}
