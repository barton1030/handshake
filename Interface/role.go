package Interface

import "time"

type RoleListStorage interface {
	MaxPrimaryKeyId() (maxPrimaryKeyId int)
	Add(role2 Role) (err error)
	Edit(role2 Role) (err error)
	RoleById(roleId int) (Role, error)
	RoleByName(roleName string) (Role, error)
	List(offset, limit int) ([]Role, error)
}

type Role interface {
	Id() int
	Name() string
	Status() int
	PermissionMap() map[string]bool
	Creator() int
	CreateTime() time.Time
}
