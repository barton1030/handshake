package service

import (
	"errors"
	inter "handshake/Interface"
	"handshake/domain"
	"handshake/domain/log"
	user2 "handshake/domain/user"
	"time"
)

type user struct {
}

var User user

func (u user) Add(operator, roleId int, name, phone, pwd string) (err error) {
	begin := domain.Manager.Begin()
	role3, err := begin.RoleList().ClapHisLockRoleById(roleId)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	if role3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("角色不存在，请确认")
		return err
	}
	user3, err := begin.UserList().ClapHisLockUserByPhone(phone)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if user3.Id() > 0 {
		_ = begin.Rollback()
		err = errors.New("当前手机号已注册，请确认")
		return
	}
	domainUser := user2.NewUser(name, phone, pwd, roleId)
	err = begin.UserList().Add(domainUser)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	userLogData := u.reconstruction(&domainUser)
	userLog := log.NewLog(userLogData, domainUser.Id(), operator, time.Now())
	err = begin.LogList().AddRoleLog(userLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return
}

func (u user) SetRoleId(operator, userId, roleId int) (err error) {
	begin := domain.Manager.Begin()
	role3, err := begin.RoleList().ClapHisLockRoleById(roleId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if role3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("角色不存在，请确认")
		return
	}
	if role3.DeleteOrNot() {
		_ = begin.Rollback()
		err = errors.New("角色已废弃，请确认")
		return
	}
	user3, err := begin.UserList().ClapHisLockUserById(userId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if user3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("用户不存在，请确认")
		return
	}
	user3.SetRole(roleId)
	err = begin.UserList().Edit(user3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	userLogData := u.reconstruction(&user3)
	userLog := log.NewLog(userLogData, user3.Id(), operator, time.Now())
	err = begin.LogList().AddRoleLog(userLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return
}

func (u user) Delete(operator, userId int) (err error) {
	begin := domain.Manager.Begin()
	user3, err := begin.UserList().ClapHisLockUserById(userId)
	if err != nil {
		_ = begin.Rollback()
		return
	}
	if user3.Id() <= 0 {
		_ = begin.Rollback()
		err = errors.New("用户不存在，请确认")
		return
	}
	user3.Delete()
	err = begin.UserList().Edit(user3)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	userLogData := u.reconstruction(&user3)
	userLog := log.NewLog(userLogData, user3.Id(), operator, time.Now())
	err = begin.LogList().AddRoleLog(userLog)
	if err != nil {
		_ = begin.Rollback()
		return err
	}
	err = begin.Commit()
	return
}

func (u user) List(operator, offset, limit int) (list []map[string]interface{}, err error) {
	users, err := domain.Manager.UserList().List(offset, limit)
	if err != nil {
		return
	}
	userNum := len(users)
	list = make([]map[string]interface{}, userNum, userNum)
	for index, user3 := range users {
		user4 := u.reconstruction(&user3)
		list[index] = user4
	}
	return
}

func (u user) UserId(operator, userId int) (user4 map[string]interface{}, err error) {
	user3, err := domain.Manager.UserList().UserById(userId)
	if err != nil {
		return
	}
	user4 = u.reconstruction(&user3)
	return
}

func (u user) reconstruction(user3 inter.User) (user4 map[string]interface{}) {
	user4 = make(map[string]interface{})
	user4["id"] = user3.Id()
	user4["name"] = user3.Name()
	user4["phone"] = user3.Phone()
	user4["role_id"] = user3.RoleId()
	user4["create_time"] = user3.CreateTime()
	return user4
}
