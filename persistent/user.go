package persistent

import (
	inter "handshake/Interface"
	"handshake/persistent/internal"
	"strconv"
	"time"
)

type userDao struct {
	tableName string
}

var UserDao = userDao{
	tableName: "hand_shake_user",
}

func (u userDao) Add(user inter.User) error {
	user2 := u.transformation(user)
	err := internal.DbConn().Table(u.tableName).Create(&user2).Error
	return err
}

func (u userDao) Edit(user inter.User) error {
	user2 := u.transformation(user)
	whereUserId := strconv.Itoa(user2.Id())
	err := internal.DbConn().Table(u.tableName).Where("user_id = ?", whereUserId).Updates(user2).Limit(1).Error
	return err
}

func (u userDao) Delete(user inter.User) error {
	err := internal.DbConn().Table(u.tableName).Delete(&struct {
		UserId int
	}{UserId: user.Id()}).Limit(1).Error
	return err
}

func (u userDao) UserById(userId int) (inter.User, error) {
	user := storageUser{}
	whereUserId := strconv.Itoa(userId)
	err := internal.DbConn().Table(u.tableName).Where("user_id = ?", whereUserId).First(&user).Error
	if err != nil && err.Error() == "record not found" {
		err = nil
	}
	return user, err
}

func (u userDao) UserList(offset, limit int) ([]inter.User, error) {
	var users []storageUser
	err := internal.DbConn().Table(u.tableName).Offset(offset).Limit(limit).Find(&users).Error
	rolesLen := len(users)
	interUsers := make([]inter.User, rolesLen, rolesLen)
	for key, value := range users {
		interUsers[key] = value
	}
	return interUsers, err
}

func (u userDao) transformation(user inter.User) (user2 storageUser) {
	user2.UserId = user.Id()
	user2.UserStatus = user.Status()
	user2.UserName = user.Name()
	user2.UserPhone = user.Phone()
	user2.UserPwd = user.Pwd()
	user2.UserRoleId = user.RoleId()
	user2.UserCreateTime = user.CreateTime()
	return user2
}

type storageUser struct {
	UserId         int       `json:"user_id" gorm:"column:user_id;primary_key"`
	UserStatus     int       `json:"user_status" gorm:"column:user_status"`
	UserName       string    `json:"user_name" gorm:"column:user_name"`
	UserPhone      string    `json:"user_phone" gorm:"column:user_phone"`
	UserPwd        string    `json:"user_pwd" gorm:"column:user_pwd"`
	UserRoleId     int       `json:"user_role_id" gorm:"column:user_role_id"`
	UserCreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

func (u storageUser) Id() int {
	return u.UserId
}

func (u storageUser) Name() string {
	return u.UserName
}

func (u storageUser) Phone() string {
	return u.UserPhone
}

func (u storageUser) Pwd() string {
	return u.UserPwd
}

func (u storageUser) CreateTime() time.Time {
	return u.UserCreateTime
}

func (u storageUser) RoleId() int {
	return u.UserRoleId
}

func (u storageUser) Status() int {
	return u.UserStatus
}
