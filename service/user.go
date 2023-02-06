package service

import (
	"errors"
	role2 "handshake/domain/role"
	user2 "handshake/domain/user"
)

type user struct {
}

var User user

func (u user) Add(name, phone, pwd string, roleId int) error {
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
	return err
}

func (u user) SetRoleId(userId, roleId int) error {
	role3, err := role2.List.RoleById(roleId)
	if err != nil {
		return err
	}
	if role3.Id() <= 0 {
		err = errors.New("角色不存在，请确认")
		return err
	}
	user3, err := user2.List.UserId(userId)
	if err != nil {
		return err
	}
	if user3.Id() <= 0 {
		err = errors.New("用户不存在，请确认")
		return err
	}
	user3.SetRole(roleId)
	err = user2.List.Edit(user3)
	return err
}

func (u user) Delete(userId int) error {
	user3, err := user2.List.UserId(userId)
	if err != nil {
		return err
	}
	if user3.Id() <= 0 {
		err = errors.New("用户不存在，请确认")
		return err
	}
	user3.Delete()
	err = user2.List.Edit(user3)
	return err
}

func (u user) List(offset, limit int) (userList []map[string]interface{}, err error) {
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
	return userList, err
}

func (u user) UserId(userId int) (user4 map[string]interface{}, err error) {
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
	return user4, err
}
