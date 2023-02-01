package service

import (
	"errors"
	role2 "handshake/domain/role"
)

type role struct {
}

var RoleService role

func (r role) Add(name string, creator int) (err error) {
	domainRole, err := role2.List.RoleByName(name)
	if domainRole.Id() > 0 {
		err = errors.New("不要重复添加")
	}
	if err != nil {
		return err
	}

	domainRole = role2.NewRole(name, creator)
	err = role2.List.Add(domainRole)
	return err
}

func (r role) RoleById(roleId int) (role3 map[string]interface{}, err error) {
	domainRole, err := role2.List.RoleById(roleId)
	if err != nil {
		return
	}

	role3 = make(map[string]interface{})
	role3["id"] = domainRole.Id()
	role3["name"] = domainRole.Name()
	role3["creator"] = domainRole.Creator()
	role3["create_time"] = domainRole.CreateTime()
	role3["permission"] = domainRole.PermissionMap()
	return
}

func (r role) EditName(roleId int, roleName string) (err error) {
	domainRole, err := role2.List.RoleById(roleId)
	if err != nil {
		return
	}
	domainRole.SetName(roleName)
	err = role2.List.Edit(domainRole)
	return
}

func (r role) SetPermission(roleId int, permissionKey string, permissionValue bool) (err error) {
	domainRole, err := role2.List.RoleById(roleId)
	if err != nil {
		return
	}
	domainRole.SetPermission(permissionKey, permissionValue)
	err = role2.List.Edit(domainRole)
	return
}

func (r role) List(offset, limit int) (roles []map[string]interface{}, err error) {
	domainRoles, err := role2.List.List(offset, limit)
	if err != nil {
		return
	}
	if len(domainRoles) == 0 {
		return
	}
	for _, domainRole := range domainRoles {
		role3 := make(map[string]interface{})
		role3["id"] = domainRole.Id()
		role3["name"] = domainRole.Name()
		role3["creator"] = domainRole.Creator()
		role3["create_time"] = domainRole.CreateTime()
		role3["permission"] = domainRole.PermissionMap()
		roles = append(roles, role3)
	}
	return
}

func (r role) Delete(roleId int) (err error) {
	domainRole, err := role2.List.RoleById(roleId)
	if err != nil {
		return
	}
	if domainRole.Id() <= 0 {
		err = errors.New("角色不存在")
		return
	}
	err = role2.List.Delete(domainRole)
	return
}
