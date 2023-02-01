package persistent

import (
	"encoding/json"
	inter "handshake/Interface"
	"handshake/persistent/internal"
	"time"
)

type roleDao struct {
	tableName string
}

var RoleDao = roleDao{
	tableName: "hand_shake_role",
}

func (r roleDao) Add(role2 inter.DomainRole) (err error) {
	role3 := r.transformation(role2)
	err = internal.DbConn().Table(r.tableName).Create(role3).Error
	return err
}

func (r roleDao) Edit(role2 inter.DomainRole) (err error) {
	role3 := r.transformation(role2)
	err = internal.DbConn().Table(r.tableName).Model(&struct {
		Id int
	}{Id: role2.Id()}).Updates(role3).Error
	return err
}

func (r roleDao) Delete(role2 inter.DomainRole) (err error) {
	err = internal.DbConn().Table(r.tableName).Delete(&struct {
		Id int
	}{Id: role2.Id()}).Limit(1).Error
	return err
}

func (r roleDao) RoleById(roleId int) (inter.RoleStorage, error) {
	role := storageRole{}
	err := internal.DbConn().Table(r.tableName).First(&role, roleId).Error
	return role, err
}

func (r roleDao) RoleByName(roleName string) (inter.RoleStorage, error) {
	role := storageRole{}
	err := internal.DbConn().Table(r.tableName).Where("name = ?", roleName).First(&role).Error

	if err != nil && err.Error() == "record not found" {
		err = nil
	}
	return role, err
}

func (r roleDao) List(offset, limit int) ([]inter.RoleStorage, error) {
	var roles []storageRole
	err := internal.DbConn().Table(r.tableName).Offset(offset).Limit(limit).Find(&roles).Error
	rolesLen := len(roles)
	interRoleStorage := make([]inter.RoleStorage, rolesLen, rolesLen)
	for key, value := range roles {
		interRoleStorage[key] = value
	}
	return interRoleStorage, err
}

func (r roleDao) transformation(role2 inter.DomainRole) (role3 storageRole) {
	role3.Id = role2.Id()
	role3.Name = role2.Name()
	role3.Creator = role2.Creator()
	role3.CreateTime = role2.CreateTime()
	permission := role2.PermissionMap()
	permissionByteSlice, _ := json.Marshal(permission)
	role3.PermissionMap = string(permissionByteSlice)
	return role3
}

type storageRole struct {
	Id            int       `json:"id" gorm:"id"`
	Name          string    `json:"name" gorm:"name"`
	PermissionMap string    `json:"permission_map" gorm:"permission_map"`
	Creator       int       `json:"creator" gorm:"creator"`
	CreateTime    time.Time `json:"create_time" gorm:"create_time"`
}

func (r storageRole) RoleId() int {
	return r.Id
}

func (r storageRole) RoleName() string {
	return r.Name
}

func (r storageRole) RolePermissionMap() map[string]bool {
	permissionMap := make(map[string]bool)
	if len(r.PermissionMap) < 0 {
		return permissionMap
	}
	json.Unmarshal([]byte(r.PermissionMap), &permissionMap)
	return permissionMap
}

func (r storageRole) RoleCreator() int {
	return r.Creator
}

func (r storageRole) RoleCreateTime() time.Time {
	return r.CreateTime
}
