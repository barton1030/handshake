package user

import "time"

type user struct {
	id         int
	status     int
	name       string
	phone      string
	roleId     int
	pwd        string
	createTime time.Time
}

const (
	NormalStatus = 1
	DeleteStatus = -1
)

func NewUser(name, phone, pwd string, roleId int) user {
	return user{
		status:     NormalStatus,
		name:       name,
		phone:      phone,
		pwd:        pwd,
		roleId:     roleId,
		createTime: time.Now(),
	}
}

func (u *user) Id() int {
	return u.id
}

func (u *user) SetId(id int) {
	u.id = id
}

func (u *user) RoleId() int {
	return u.roleId
}

func (u *user) SetRole(roleId int) {
	u.roleId = roleId
}

func (u *user) Name() string {
	return u.name
}

func (u *user) SetName(name string) {
	u.name = name
}

func (u *user) Phone() string {
	return u.phone
}

func (u *user) SetPhone(phone string) {
	u.phone = phone
}

func (u *user) CreateTime() time.Time {
	return u.createTime
}

func (u *user) Pwd() string {
	return u.pwd
}

func (u *user) Status() int {
	return u.status
}

func (u *user) Delete() {
	u.status = DeleteStatus
}

func (u *user) DeleteOrNot() bool {
	if u.status == DeleteStatus {
		return true
	}
	return false
}
