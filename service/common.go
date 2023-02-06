package service

import (
	"errors"
	role2 "handshake/domain/role"
	user2 "handshake/domain/user"
)

func permissionVerification(userId int, uri string) (err error) {
	user3, err := user2.List.UserId(userId)
	if err != nil {
		return
	}
	if user3.Id() <= 0 {
		err = errors.New("用户不存在，请确认")
		return
	}
	role3, err := role2.List.RoleById(user3.RoleId())
	if err != nil {
		return
	}
	permission := role3.Permission(uri)
	if !permission {
		err = errors.New("无权限，请确认")
		return
	}
	return
}
