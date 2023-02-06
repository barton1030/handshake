package user

import (
	inter "handshake/Interface"
	"handshake/persistent"
)

type list struct {
	nextId  int
	storage inter.StorageUserList
}

var List = list{nextId: 1, storage: persistent.UserDao}

func (l *list) Init() {
	maxPrimaryKeyId := l.storage.MaxPrimaryKeyId()
	l.nextId = maxPrimaryKeyId + 1
}

func (l *list) Add(user2 user) (err error) {
	user2.id = l.nextId
	err = l.storage.Add(&user2)
	if err != nil {
		return err
	}
	l.nextId++
	return err
}

func (l *list) Edit(user2 user) (err error) {
	err = l.storage.Edit(&user2)
	return err
}

func (l *list) Delete(user2 user) (err error) {
	err = l.storage.Delete(&user2)
	return err
}

func (l *list) UserId(userId int) (user2 user, err error) {
	storageUser, err := l.storage.UserById(userId)
	if err != nil {
		return
	}
	user2 = l.reconstruction(storageUser)
	return
}

func (l *list) List(offset, limit int) (userList []user, err error) {
	storageUserList, err := l.storage.UserList(offset, limit)
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

func (l *list) reconstruction(user inter.User) (user2 user) {
	user2.id = user.Id()
	user2.name = user.Name()
	user2.phone = user.Phone()
	user2.roleId = user.RoleId()
	user2.pwd = user.Pwd()
	user2.createTime = user.CreateTime()
	return
}
