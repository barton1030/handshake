package service

import (
	"errors"
	"handshake/domain"
	role2 "handshake/domain/role"
)

type role struct {
}

var Role role

func (r role) Add(operator int, name, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	role3, err := domain.Manager.RoleList().RoleByName(name)
	if role3.Id() > 0 {
		err = errors.New("不要重复添加")
	}
	if err != nil {
		return err
	}
	role4 := role2.NewRole(name, operator)
	err = domain.Manager.RoleList().Add(role4)
	return err
}

func (r role) RoleById(operator, roleId int, uri string) (role4 map[string]interface{}, err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	role3, err := domain.Manager.RoleList().RoleById(roleId)
	if err != nil {
		return
	}
	role4 = make(map[string]interface{})
	role4["id"] = role3.Id()
	role4["name"] = role3.Name()
	role4["creator"] = role3.Creator()
	role4["create_time"] = role3.CreateTime()
	role4["permission"] = role3.PermissionMap()
	return
}

func (r role) EditName(operator, roleId int, roleName, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	role3, err := domain.Manager.RoleList().RoleById(roleId)
	if err != nil {
		return
	}
	role3.SetName(roleName)
	err = domain.Manager.RoleList().Edit(role3)
	return
}

func (r role) SetPermission(operator, roleId int, permissionKey string, permissionValue bool, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	role3, err := domain.Manager.RoleList().RoleById(roleId)
	if err != nil {
		return
	}
	role3.SetPermission(permissionKey, permissionValue)
	err = domain.Manager.RoleList().Edit(role3)
	return
}

func (r role) List(operator, offset, limit int, uri string) (list []map[string]interface{}, err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	domainRoles, err := domain.Manager.RoleList().List(offset, limit)
	if err != nil {
		return
	}
	roleNum := len(domainRoles)
	list = make([]map[string]interface{}, roleNum, roleNum)
	for index, role2 := range domainRoles {
		role3 := make(map[string]interface{})
		role3["id"] = role2.Id()
		role3["name"] = role2.Name()
		role3["creator"] = role2.Creator()
		role3["create_time"] = role2.CreateTime()
		role3["permission"] = role2.PermissionMap()
		list[index] = role3
	}
	return
}

func (r role) Delete(operator, roleId int, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	role3, err := domain.Manager.RoleList().RoleById(roleId)
	if err != nil {
		return
	}
	if role3.Id() <= 0 {
		err = errors.New("角色不存在")
		return
	}
	role3.Delete()
	err = domain.Manager.RoleList().Edit(role3)
	return
}
