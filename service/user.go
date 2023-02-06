package service

import (
	"errors"
	role2 "handshake/domain/role"
	user2 "handshake/domain/user"
)

type user struct {
}

var User user

func (u user) Add(operator, roleId int, name, phone, pwd, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return err
	}
	role3, err := role2.List.RoleById(roleId)
	if err != nil {
		return err
	}
	if role3.Id() <= 0 {
		err = errors.New("角色不存在，请确认")
		return err
	}
	domainUser := user2.NewUser(name, phone, pwd, roleId)
	err = user2.List.Add(domainUser)
	return
}

func (u user) SetRoleId(operator, userId, roleId int, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	role3, err := role2.List.RoleById(roleId)
	if err != nil {
		return
	}
	if role3.Id() <= 0 {
		err = errors.New("角色不存在，请确认")
		return
	}
	user3, err := user2.List.UserId(userId)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("用户不存在，请确认")
		return
	}
	user3.SetRole(roleId)
	err = user2.List.Edit(user3)
	return
}

func (u user) Delete(operator, userId int, uri string) (err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	user3, err := user2.List.UserId(userId)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("用户不存在，请确认")
		return
	}
	user3.Delete()
	err = user2.List.Edit(user3)
	return
}

func (u user) List(operator, offset, limit int, uri string) (userList []map[string]interface{}, err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	users, err := user2.List.List(offset, limit)
	if err != nil {
		return
	}
	userNum := len(users)
	userList = make([]map[string]interface{}, userNum, userNum)
	for index, user3 := range users {
		user4 := make(map[string]interface{})
		user4["id"] = user3.Id()
		user4["name"] = user3.Name()
		user4["phone"] = user3.Phone()
		user4["role_id"] = user3.RoleId()
		user4["create_time"] = user3.CreateTime()
		userList[index] = user4
	}
	return
}

func (u user) UserId(operator, userId int, uri string) (user4 map[string]interface{}, err error) {
	err = permissionVerification(operator, uri)
	if err != nil {
		return
	}
	user3, err := user2.List.UserId(userId)
	if err != nil {
		return
	}
	user4 = make(map[string]interface{})
	user4["id"] = user3.Id()
	user4["name"] = user3.Name()
	user4["phone"] = user3.Phone()
	user4["role_id"] = user3.RoleId()
	user4["create_time"] = user3.CreateTime()
	return
}
