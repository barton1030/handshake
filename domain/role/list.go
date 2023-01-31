package role

type list struct {
	storage map[int]role
}

var List list

func (l *list) Add() (err error) {
	return err
}

func (l *list) Edit(cRole role) (err error) {
	return err
}

func (l *list) Delete(roleId int) (err error) {
	return err
}

func (l *list) List() (roleList map[int]role, err error) {
	return
}

func (l *list) Role(roleId int) (cRole role, err error) {
	return
}
