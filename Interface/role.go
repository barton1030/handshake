package Interface

import "time"

type RoleListStorage interface {
	Add(role2 DomainRole) (err error)
	Edit(role2 DomainRole) (err error)
	Delete(role2 DomainRole) (err error)
	RoleById(roleId int) (RoleStorage, error)
	RoleByName(roleName string) (RoleStorage, error)
}

type RoleStorage interface {
	RoleId() int
	RoleName() string
	RolePermissionMap() map[string]bool
	RoleCreator() int
	RoleCreateTime() time.Time
}

type DomainRole interface {
	Id() int
	Name() string
	PermissionMap() map[string]bool
	Creator() int
	CreateTime() time.Time
}
