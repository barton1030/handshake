package service

import (
	"errors"
	role2 "handshake/domain/role"
	user2 "handshake/domain/user"
)

type user struct {
}

var UserService user

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
