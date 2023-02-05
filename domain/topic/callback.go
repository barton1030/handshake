package topic

type callback struct {
	url     string
	method  string
	headers map[string]interface{}
	cookies map[string]interface{}
}

func NewCallBack(url, method string, cookies, headers map[string]interface{}) callback {
	return callback{
		url:     url,
		method:  method,
		headers: headers,
		cookies: cookies,
	}
}

func (c callback) Do(data map[string]interface{}) (res map[string]interface{}, err error) {
	return
}

func (c callback) Headers() map[string]interface{} {
	return c.headers
}

func (c callback) Cookies() map[string]interface{} {
	return c.cookies
}

func (c callback) Url() string {
	return c.url
}

func (c callback) Method() string {
	return c.method
}
