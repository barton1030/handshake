package persistent

import (
	"encoding/json"
	inter "handshake/Interface"
	"strconv"
	"time"
)

type roleDao struct {
	transactionId int
	tableName     string
}

var RoleDao = roleDao{
	tableName: "hand_shake_role",
}

func (r roleDao) MaxPrimaryKeyId() (maxPrimaryKeyId int) {
	role3 := storageRole{}
	err := transactionController.dbConn(r.transactionId).Table(r.tableName).Last(&role3).Error
	if err != nil {
		return
	}
	if role3.Id() <= 0 {
		return
	}
	maxPrimaryKeyId = role3.Id()
	return
}

func (r roleDao) Add(role2 inter.Role) (err error) {
	role3 := r.transformation(role2)
	err = transactionController.dbConn(r.transactionId).Table(r.tableName).Create(role3).Error
	return err
}

func (r roleDao) Edit(role2 inter.Role) (err error) {
	role3 := r.transformation(role2)
	whereRoleId := strconv.Itoa(role2.Id())
	err = transactionController.dbConn(r.transactionId).Table(r.tableName).Where("id = ?", whereRoleId).Updates(role3).Error
	return err
}

func (r roleDao) RoleById(roleId int) (inter.Role, error) {
	role := storageRole{}
	whereRoleId := strconv.Itoa(roleId)
	err := transactionController.dbConn(r.transactionId).Table(r.tableName).Where("id = ?", whereRoleId).First(&role).Error
	if err != nil && err.Error() == "record not found" {
		err = nil
	}
	return role, err
}

func (r roleDao) RoleByName(roleName string) (inter.Role, error) {
	role := storageRole{}
	err := transactionController.dbConn(r.transactionId).Table(r.tableName).Where("name = ?", roleName).First(&role).Error

	if err != nil && err.Error() == "record not found" {
		err = nil
	}
	return role, err
}

func (r roleDao) List(offset, limit int) ([]inter.Role, error) {
	var roles []storageRole
	err := transactionController.dbConn(r.transactionId).Table(r.tableName).Offset(offset).Limit(limit).Find(&roles).Error
	rolesLen := len(roles)
	interRoleStorage := make([]inter.Role, rolesLen, rolesLen)
	for key, value := range roles {
		interRoleStorage[key] = value
	}
	return interRoleStorage, err
}

func (r roleDao) transformation(role2 inter.Role) (role3 storageRole) {
	role3.SId = role2.Id()
	role3.SStatus = role2.Status()
	role3.SName = role2.Name()
	role3.SCreator = role2.Creator()
	role3.SCreateTime = role2.CreateTime()
	permission := role2.PermissionMap()
	permissionByteSlice, _ := json.Marshal(permission)
	role3.SPermissionMap = string(permissionByteSlice)
	return role3
}

type storageRole struct {
	SId            int       `json:"id" gorm:"column:id;primary_key"`
	SStatus        int       `json:"status" gorm:"column:status"`
	SName          string    `json:"name" gorm:"column:name"`
	SPermissionMap string    `json:"permission_map" gorm:"column:permission_map"`
	SCreator       int       `json:"creator" gorm:"column:creator"`
	SCreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
}

func (r storageRole) Id() int {
	return r.SId
}

func (r storageRole) Name() string {
	return r.SName
}

func (r storageRole) Status() int {
	return r.SStatus
}

func (r storageRole) PermissionMap() map[string]bool {
	permissionMap := make(map[string]bool)
	if len(r.SPermissionMap) < 0 {
		return permissionMap
	}
	json.Unmarshal([]byte(r.SPermissionMap), &permissionMap)
	return permissionMap
}

func (r storageRole) Creator() int {
	return r.SCreator
}

func (r storageRole) CreateTime() time.Time {
	return r.SCreateTime
}
