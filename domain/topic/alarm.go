package topic

type alarm struct {
	url        string
	method     string
	recipients []interface{}
}

func (a alarm) Do(information map[string]interface{}, recipients []interface{}) {

}
