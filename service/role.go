package service

import (
	"errors"
	role2 "handshake/domain/role"
)

type role struct {
}

var RoleService role

func (r role) Add(operator int, name, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	domainRole, err := role2.List.RoleByName(name)
	if domainRole.Id() > 0 {
		err = errors.New("不要重复添加")
	}
	if err != nil {
		return err
	}
	domainRole = role2.NewRole(name, operator)
	err = role2.List.Add(domainRole)
	return err
}

func (r role) RoleById(operator, roleId int, uri string) (role3 map[string]interface{}, err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
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

func (r role) EditName(operator, roleId int, roleName, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	domainRole, err := role2.List.RoleById(roleId)
	if err != nil {
		return
	}
	domainRole.SetName(roleName)
	err = role2.List.Edit(domainRole)
	return
}

func (r role) SetPermission(operator, roleId int, permissionKey string, permissionValue bool, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	domainRole, err := role2.List.RoleById(roleId)
	if err != nil {
		return
	}
	domainRole.SetPermission(permissionKey, permissionValue)
	err = role2.List.Edit(domainRole)
	return
}

func (r role) List(operator, offset, limit int, uri string) (roles []map[string]interface{}, err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
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

func (r role) Delete(operator, roleId int, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	role3, err := role2.List.RoleById(roleId)
	if err != nil {
		return
	}
	if role3.Id() <= 0 {
		err = errors.New("角色不存在")
		return
	}
	role3.Delete()
	err = role2.List.Edit(role3)
	return
}
