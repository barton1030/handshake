package service

import (
	"errors"
	"handshake/domain"
)

func permissionVerification(operator int, uri string) (err error) {
	user3, err := domain.Manager.UserList().UserById(operator)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("操作者不存在，请确认")
		return
	}
	role3, err := domain.Manager.RoleList().RoleById(user3.RoleId())
	if err != nil {
		return
	}
	permission := role3.Permission(uri)
	if !permission {
		err = errors.New("操作者无权限，请确认")
		return
	}
	return
}
