package Interface

import "time"

type StorageUserList interface {
	MaxPrimaryKeyId() (maxPrimaryKeyId int)
	Add(user User) error
	Edit(user User) error
	Delete(user User) error
	UserById(userId int) (User, error)
	UserByPhone(phone string) (User, error)
	UserList(startId, limit int) ([]User, error)
	ClapHisLockUserById(userId int) (User, error)
	ClapHisLockUserByPhone(phone string) (User, error)
	UserCountByRoleId(roleId int) (counter int, err error)
}

type User interface {
	Id() int
	Name() string
	Status() int
	Phone() string
	Pwd() string
	RoleId() int
	CreateTime() time.Time
}
