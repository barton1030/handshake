package user

import "time"

type user struct {
	id         int
	name       string
	phone      string
	role       int
	createTime time.Time
}

func NewUser(name, phone string) user {
	return user{
		name:       name,
		phone:      phone,
		createTime: time.Now(),
	}
}

func (u *user) SetId(id int) (err error) {
	u.id = id
	return err
}

func (u *user) SetRole(roleId int) (err error) {
	u.role = roleId
	return err
}

func (u *user) SetName(name string) (err error) {
	u.name = name
	return err
}

func (u *user) SetPhone(phone string) (err error) {
	u.phone = phone
	return err
}
