package log

import "time"

type log struct {
	id           int
	data         map[string]interface{}
	businessType int
	businessId   int
	creator      int
	createTime   time.Time
}

func NewLog(data map[string]interface{}, businessId, creator int, createTime time.Time) log {
	return log{
		data:       data,
		businessId: businessId,
		creator:    creator,
		createTime: createTime,
	}
}

func (l log) Id() int {
	return l.id
}

func (l log) Data() map[string]interface{} {
	return l.data
}

func (l log) BusinessType() int {
	return l.businessType
}

func (l log) BusinessId() int {
	return l.businessId
}

func (l log) Creator() int {
	return l.creator
}

func (l log) CreateTime() time.Time {
	return l.createTime
}
