package user

import (
	inter "handshake/Interface"
	"handshake/persistent"
)

type List struct {
	storage inter.StorageUserList
}

var ListExample = List{storage: persistent.UserDao}

func (l *List) SetStorage(storageInter inter.StorageUserList) *List {
	return &List{
		storage: storageInter,
	}
}

func (l *List) Add(user2 user) (err error) {
	err = l.storage.Add(&user2)
	return err
}

func (l *List) Edit(user2 user) (err error) {
	err = l.storage.Edit(&user2)
	return err
}

func (l *List) Delete(user2 user) (err error) {
	err = l.storage.Delete(&user2)
	return err
}

func (l *List) UserById(userId int) (user2 user, err error) {
	storageUser, err := l.storage.UserById(userId)
	if err != nil {
		return
	}
	user2 = l.reconstruction(storageUser)
	return
}

func (l *List) ClapHisLockUserById(userId int) (user2 user, err error) {
	storageUser, err := l.storage.ClapHisLockUserById(userId)
	if err != nil {
		return
	}
	user2 = l.reconstruction(storageUser)
	return
}

func (l *List) UserByPhone(phone string) (user2 user, err error) {
	storageUser, err := l.storage.UserByPhone(phone)
	if err != nil {
		return
	}
	user2 = l.reconstruction(storageUser)
	return
}

func (l *List) ClapHisLockUserByPhone(phone string) (user2 user, err error) {
	storageUser, err := l.storage.ClapHisLockUserByPhone(phone)
	if err != nil {
		return
	}
	user2 = l.reconstruction(storageUser)
	return
}

func (l *List) List(startId, limit int) (userList []user, err error) {
	storageUserList, err := l.storage.UserList(startId, limit)
	if err != nil {
		return
	}
	if len(storageUserList) <= 0 {
		return
	}

	for _, storageUser := range storageUserList {
		user2 := l.reconstruction(storageUser)
		userList = append(userList, user2)
	}

	return
}

func (l *List) UserCountByRoleId(roleId int) (counter int, err error) {
	counter, err = l.storage.UserCountByRoleId(roleId)
	return
}

func (l *List) reconstruction(user inter.User) (user2 user) {
	user2.id = user.Id()
	user2.name = user.Name()
	user2.phone = user.Phone()
	user2.roleId = user.RoleId()
	user2.pwd = user.Pwd()
	user2.createTime = user.CreateTime()
	user2.status = user.Status()
	return
}
