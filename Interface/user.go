package Interface

import "time"

type StorageUserList interface {
	Add(user User) error
	Edit(user User) error
	Delete(user User) error
	UserById(userId int) (User, error)
	UserList(offset, limit int) ([]User, error)
}

type User interface {
	Id() int
	Name() string
	Phone() string
	Pwd() string
	RoleId() int
	CreateTime() time.Time
}
