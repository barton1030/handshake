package role

type list struct {
	nextId  int
	storage map[int]role
}

var List = list{storage: make(map[int]role)}

func (l *list) Add(role2 role) (err error) {
	l.storage[l.nextId] = role2
	l.nextId++
	return err
}

func (l *list) Edit(roles role) (err error) {
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
