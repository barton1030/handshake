package persistent

import (
	inter "handshake/Interface"
	"strconv"
	"time"
)

type userDao struct {
	transactionId int
	tableName     string
}

var UserDao = userDao{
	tableName: "hand_shake_user",
}

func (u userDao) MaxPrimaryKeyId() (maxPrimaryKeyId int) {
	user := storageUser{}
	err := transactionController.dbConn(u.transactionId).Table(u.tableName).Last(&user).Error
	if err != nil {
		return
	}
	if user.Id() <= 0 {
		return
	}
	maxPrimaryKeyId = user.Id()
	return
}

func (u userDao) Add(user inter.User) error {
	user2 := u.transformation(user)
	err := transactionController.dbConn(u.transactionId).Table(u.tableName).Create(&user2).Error
	return err
}

func (u userDao) Edit(user inter.User) error {
	user2 := u.transformation(user)
	whereUserId := strconv.Itoa(user2.Id())
	err := transactionController.dbConn(u.transactionId).Table(u.tableName).Where("id = ?", whereUserId).Updates(user2).Limit(1).Error
	return err
}

func (u userDao) Delete(user inter.User) error {
	err := transactionController.dbConn(u.transactionId).Table(u.tableName).Delete(&struct {
		UserId int
	}{UserId: user.Id()}).Limit(1).Error
	return err
}

func (u userDao) UserById(userId int) (inter.User, error) {
	user := storageUser{}
	whereUserId := strconv.Itoa(userId)
	err := transactionController.dbConn(u.transactionId).Table(u.tableName).Where("id = ?", whereUserId).First(&user).Error
	if err != nil && err.Error() == "record not found" {
		err = nil
	}
	return user, err
}

func (u userDao) UserByPhone(phone string) (inter.User, error) {
	user := storageUser{}
	err := transactionController.dbConn(u.transactionId).Table(u.tableName).Where("phone = ?", phone).First(&user).Error
	if err != nil && err.Error() == "record not found" {
		err = nil
	}
	return user, err
}

func (u userDao) UserList(startId, limit int) ([]inter.User, error) {
	var users []storageUser
	whereStartId := strconv.Itoa(startId)
	err := transactionController.dbConn(u.transactionId).Table(u.tableName).Where("id > ?", whereStartId).Limit(limit).Find(&users).Error
	rolesLen := len(users)
	interUsers := make([]inter.User, rolesLen, rolesLen)
	for key, value := range users {
		interUsers[key] = value
	}
	return interUsers, err
}

func (u userDao) transformation(user inter.User) (user2 storageUser) {
	user2.SId = user.Id()
	user2.SStatus = user.Status()
	user2.SName = user.Name()
	user2.SPhone = user.Phone()
	user2.SPwd = user.Pwd()
	user2.SRoleId = user.RoleId()
	user2.SCreateTime = user.CreateTime()
	return user2
}

type storageUser struct {
	SId         int       `json:"user_id" gorm:"column:id;primary_key"`
	SStatus     int       `json:"user_status" gorm:"column:status"`
	SName       string    `json:"user_name" gorm:"column:name"`
	SPhone      string    `json:"user_phone" gorm:"column:phone"`
	SPwd        string    `json:"user_pwd" gorm:"column:pwd"`
	SRoleId     int       `json:"user_role_id" gorm:"column:role_id"`
	SCreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

func (u storageUser) Id() int {
	return u.SId
}

func (u storageUser) Name() string {
	return u.SName
}

func (u storageUser) Phone() string {
	return u.SPhone
}

func (u storageUser) Pwd() string {
	return u.SPwd
}

func (u storageUser) CreateTime() time.Time {
	return u.SCreateTime
}

func (u storageUser) RoleId() int {
	return u.SRoleId
}

func (u storageUser) Status() int {
	return u.SStatus
}
