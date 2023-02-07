package Interface

import "time"

type StorageUserList interface {
	Transaction
	MaxPrimaryKeyId() (maxPrimaryKeyId int)
	Add(user User) error
	Edit(user User) error
	Delete(user User) error
	UserById(userId int) (User, error)
	UserByPhone(phone string) (User, error)
	UserList(offset, limit int) ([]User, error)
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
