package role

import (
	inter "handshake/Interface"
	"handshake/persistent"
)

type list struct {
	nextId  int
	storage inter.RoleListStorage
}

var List = list{nextId: 1, storage: persistent.RoleDao}

func (l *list) Add(role2 role) (err error) {
	role2.id = l.nextId
	err = l.storage.Add(&role2)
	if err == nil {
		l.nextId++
	}
	return err
}

func (l *list) Edit(role2 role) (err error) {
	return err
}

func (l *list) Delete(roleId int) (err error) {
	return err
}

func (l *list) List() (roleList map[int]role, err error) {
	return
}

func (l *list) Role(roleId int) (role2 role, err error) {
	return
}
