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

func (r role) RoleById(roleId int) (role map[string]interface{}, err error) {
	domainRole, err := role2.List.RoleById(roleId)
	if err != nil {
		return
	}
	role = make(map[string]interface{})
	role["id"] = domainRole.Id()
	role["name"] = domainRole.Name()
	role["creator"] = domainRole.Creator()
	role["create_time"] = domainRole.CreateTime()
	role["permission"] = domainRole.PermissionMap()
	return
}
