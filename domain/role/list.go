package role

import (
	inter "handshake/Interface"
	"handshake/persistent"
)

type List struct {
	storage inter.StorageRoleList
}

var ListExample = List{storage: persistent.RoleDao}

func (l *List) SetStorage(storageInter inter.StorageRoleList) *List {
	return &List{
		storage: storageInter,
	}
}

func (l *List) Add(role2 role) (err error) {
	err = l.storage.Add(&role2)
	return err
}

func (l *List) Edit(role2 role) (err error) {
	err = l.storage.Edit(&role2)
	return err
}

func (l *List) List(startId, limit int) (roleList []role, err error) {
	storageRoles, err := l.storage.List(startId, limit)
	if err != nil {
		return
	}
	for _, storageRole := range storageRoles {
		role2 := l.reconstruction(storageRole)
		roleList = append(roleList, role2)
	}
	return
}

func (l *List) RoleById(roleId int) (role2 role, err error) {
	storageRole, err := l.storage.RoleById(roleId)
	if err != nil {
		return
	}
	role2 = l.reconstruction(storageRole)
	return
}

func (l *List) RoleByName(roleName string) (role2 role, err error) {
	storageRole, err := l.storage.RoleByName(roleName)
	if err != nil {
		return
	}
	role2 = l.reconstruction(storageRole)
	return
}

func (l *List) ClapHisLockRoleByName(roleName string) (role2 role, err error) {
	storageRole, err := l.storage.ClapHisLockRoleByName(roleName)
	if err != nil {
		return
	}
	role2 = l.reconstruction(storageRole)
	return
}

func (l *List) ClapHisLockRoleById(roleId int) (role2 role, err error) {
	storageRole, err := l.storage.ClapHisLockRoleById(roleId)
	if err != nil {
		return
	}
	role2 = l.reconstruction(storageRole)
	return
}

func (l *List) reconstruction(originRole inter.Role) (role2 role) {
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
