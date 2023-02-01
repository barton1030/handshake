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

func (r *role) Id() int {
	return r.id
}

func (r *role) Name() string {
	return r.name
}

func (r *role) PermissionMap() map[string]bool {
	return r.permissionMap
}

func (r *role) Creator() int {
	return r.creator
}

func (r *role) CreateTime() time.Time {
	return r.createTime
}

func (r *role) SetPermission(permissionKey string, permissionValue bool) {
	r.permissionLock.Lock()
	defer r.permissionLock.Unlock()
	r.permissionMap[permissionKey] = permissionValue
}

func (r *role) Permission(permissionKey string) (permissionValue bool) {
	value, ok := r.permissionMap[permissionKey]
	if ok {
		permissionValue = value
	}
	return
}
