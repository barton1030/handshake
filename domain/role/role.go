package role

import (
	"sync"
	"time"
)

type role struct {
	id             int
	name           string
	permissionMap  map[string]bool
	permissionLock sync.Mutex
	creator        int
	createTime     time.Time
}

func NewRole(name string, creator int) role {
	return role{
		name:       name,
		creator:    creator,
		createTime: time.Now(),
	}
}

func (r *role) SetPermission(permissionKey string, permissionValue bool) {
	r.permissionLock.Lock()
	defer r.permissionLock.Unlock()
	r.permissionMap[permissionKey] = permissionValue
}
