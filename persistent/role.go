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
	role3 := storageRole{}
	role3.Id = role2.Id()
	role3.Name = role2.Name()
	role3.Creator = role2.Creator()
	role3.CreateTime = role2.CreateTime()
	permission := role2.PermissionMap()
	permissionByteSlice, _ := json.Marshal(permission)
	role3.PermissionMap = string(permissionByteSlice)
	err = internal.DbConn().Table(r.tableName).Create(role3).Error
	return err
}

func (r roleDao) Edit(role2 inter.DomainRole) (err error) {
	return err
}

func (r roleDao) Delete(role2 inter.DomainRole) (err error) {
	return err
}

func (r roleDao) RoleById(roleId int) (role2 inter.RoleStorage, err error) {
	role2 = storageRole{}
	err = internal.DbConn().Table(r.tableName).Where("id", roleId).Scan(&role2).Error
	return
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
	json.Unmarshal([]byte(r.PermissionMap), &permissionMap)
	return permissionMap
}

func (r storageRole) RoleCreator() int {
	return r.Creator
}

func (r storageRole) RoleCreateTime() time.Time {
	return r.CreateTime
}
