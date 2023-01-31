package topic

type callback struct {
	url     string
	method  string
	headers map[string]interface{}
	cookies map[string]interface{}
}

func (c callback) Do(data map[string]interface{}) (res map[string]interface{}, err error) {
	return
}
