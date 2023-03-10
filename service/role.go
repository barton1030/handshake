package service

import (
	"errors"
	inter "handshake/Interface"
	"handshake/domain"
	"handshake/domain/log"
	role2 "handshake/domain/role"
	"time"
)

type role struct {
}

var Role role

func (r role) Add(operator int, name string) (err error) {
	begin := domain.Manager.Begin()
	role3, err := begin.RoleList().ClapHisLockRoleByName(name)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	if role3.Id() > 0 {
		_ = begin.Rollback()
		err = errors.New("不要重复添加")
		return err
	}
	role4 := role2.NewRole(name, operator)
	err = begin.RoleList().Add(role4)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	roleLogData := r.reconstruction(&role4)
	roleLog := log.NewLog(roleLogData, role4.Id(), operator, time.Now())
	err = begin.LogList().AddRoleLog(roleLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return err
}

func (r role) RoleById(operator, roleId int) (role4 map[string]interface{}, err error) {
	role3, err := domain.Manager.RoleList().RoleById(roleId)
	if err != nil {
		return
	}
	if role3.Id() <= 0 {
		err = errors.New("角色不存在，请注意")
		return
	}
	role4 = r.reconstruction(&role3)
	return
}

func (r role) EditName(operator, roleId int, roleName string) (err error) {
	begin := domain.Manager.Begin()
	role3, err := begin.RoleList().ClapHisLockRoleByName(roleName)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	if role3.Id() > 0 {
		_ = begin.Rollback()
		err = errors.New("不要重复添加")
		return err
	}
	role3, err = begin.RoleList().ClapHisLockRoleById(roleId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if role3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("角色不存在，请注意！")
		return err
	}
	role3.SetName(roleName)
	err = begin.RoleList().Edit(role3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	roleLogData := r.reconstruction(&role3)
	roleLog := log.NewLog(roleLogData, role3.Id(), operator, time.Now())
	err = begin.LogList().AddRoleLog(roleLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return
}

func (r role) SetPermission(operator, roleId int, permissionKey string, permissionValue bool) (err error) {
	begin := domain.Manager.Begin()
	role3, err := begin.RoleList().ClapHisLockRoleById(roleId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if role3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("角色不存在，请注意！")
		return err
	}
	role3.SetPermission(permissionKey, permissionValue)
	err = begin.RoleList().Edit(role3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	roleLogData := r.reconstruction(&role3)
	roleLog := log.NewLog(roleLogData, role3.Id(), operator, time.Now())
	err = begin.LogList().AddRoleLog(roleLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return
}

func (r role) List(operator, offset, limit int) (list []map[string]interface{}, err error) {
	domainRoles, err := domain.Manager.RoleList().List(offset, limit)
	if err != nil {
		return
	}
	roleNum := len(domainRoles)
	list = make([]map[string]interface{}, roleNum, roleNum)
	for index, role2 := range domainRoles {
		role3 := r.reconstruction(&role2)
		list[index] = role3
	}
	return
}

func (r role) Delete(operator, roleId int) (err error) {
	begin := domain.Manager.Begin()
	role3, err := begin.RoleList().ClapHisLockRoleById(roleId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if role3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("角色不存在")
		return
	}
	counter, err := begin.UserList().UserCountByRoleId(roleId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if counter > 0 {
		_ = begin.Rollback()
		err = errors.New("该角色已与具体用户绑定不可删除, 请注意")
		return
	}
	role3.Delete()
	err = begin.RoleList().Edit(role3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	roleLogData := r.reconstruction(&role3)
	roleLog := log.NewLog(roleLogData, role3.Id(), operator, time.Now())
	err = begin.LogList().AddRoleLog(roleLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return
}

func (r role) reconstruction(role3 inter.Role) (role4 map[string]interface{}) {
	role4 = make(map[string]interface{})
	role4["id"] = role3.Id()
	role4["name"] = role3.Name()
	role4["creator"] = role3.Creator()
	role4["create_time"] = role3.CreateTime()
	role4["permission"] = role3.PermissionMap()
	return role4
}
