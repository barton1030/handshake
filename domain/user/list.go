package user

type list struct {
	nextId  int
	storage map[int]user
}

var List list

func (l *list) Add(user2 user) (err error) {
	return err
}

func (l *list) Edit(user2 user) (err error) {
	return err
}

func (l *list) Delete(user2 user) (err error) {
	return err
}

func (l *list) UserId(userId int) (user2 user, err error) {
	return
}

func (l *list) UserName(userName string) (user2 user, err error) {
	return
}
