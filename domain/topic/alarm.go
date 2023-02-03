package topic

type alarm struct {
	url        string
	method     string
	recipients []interface{}
}

func (a alarm) Do(information map[string]interface{}, recipients []interface{}) (res map[string]interface{}, err error) {
	return
}

func (a alarm) Url() string {
	return a.url
}

func (a alarm) Method() string {
	return a.method
}

func (a alarm) Recipients() []interface{} {
	return a.recipients
}
