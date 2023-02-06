package role

import (
	inter "handshake/Interface"
	"handshake/persistent"
)

type list struct {
	nextId  int
	storage inter.RoleListStorage
}

var List = list{nextId: 1, storage: persistent.RoleDao}

func (l *list) Init() {
	maxPrimaryKeyId := l.storage.MaxPrimaryKeyId()
	l.nextId = maxPrimaryKeyId + 1
}

func (l *list) Add(role2 role) (err error) {
	role2.id = l.nextId
	err = l.storage.Add(&role2)
	if err == nil {
		l.nextId++
	}
	return err
}

func (l *list) Edit(role2 role) (err error) {
	err = l.storage.Edit(&role2)
	return err
}

func (l *list) List(offset, limit int) (roleList []role, err error) {
	storageRoles, err := l.storage.List(offset, limit)
	if err != nil {
		return
	}
	for _, storageRole := range storageRoles {
		role2 := l.reconstruction(storageRole)
		roleList = append(roleList, role2)
	}
	return
}

func (l *list) RoleById(roleId int) (role2 role, err error) {
	storageRole, err := l.storage.RoleById(roleId)
	if err != nil {
		return
	}
	role2 = l.reconstruction(storageRole)
	return
}

func (l *list) RoleByName(roleName string) (role2 role, err error) {
	storageRole, err := l.storage.RoleByName(roleName)
	if err != nil {
		return
	}
	role2 = l.reconstruction(storageRole)
	return
}

func (l *list) reconstruction(originRole inter.Role) (role2 role) {
	role2.id = originRole.Id()
	role2.name = originRole.Name()
	role2.creator = originRole.Creator()
	role2.permissionMap = originRole.PermissionMap()
	if role2.permissionMap == nil {
		role2.permissionMap = make(map[string]bool)
	}
	role2.createTime = originRole.CreateTime()
	return
}
