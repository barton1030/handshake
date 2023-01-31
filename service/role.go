package service

import role2 "handshake/domain/role"

type role struct {
}

var RoleService role

func (r role) Add(name string, creator int) (err error) {
	// 新增
	domainRole := role2.NewRole(name, creator)
	err = role2.List.Add(domainRole)
	return err
}
