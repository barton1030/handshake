package topic

import (
	"encoding/json"
	"handshake/helper"
)

type alarm struct {
	url                string
	method             string
	cookies            map[string]interface{}
	headers            map[string]interface{}
	templateParameters map[string]interface{} // 模版参数，用于接收相关的可配置参数
	recipients         map[int]int
}

func NewAlarm(url, method string, recipients map[int]int, headers, cookies, templateParameters map[string]interface{}) alarm {
	return alarm{
		url:                url,
		method:             method,
		recipients:         recipients,
		cookies:            cookies,
		headers:            headers,
		templateParameters: templateParameters,
	}
}

func (a alarm) Do(alarmInformation map[string]interface{}) (res map[string]interface{}, err error) {
	params := a.buildParameterSet(alarmInformation)
	res, err = helper.R.HttpRequest(a.url, a.method, a.headers, a.cookies, params)
	return
}

func (a alarm) Url() string {
	return a.url
}

func (a alarm) Method() string {
	return a.method
}

func (a alarm) Cookies() map[string]interface{} {
	return a.cookies
}

func (a alarm) Headers() map[string]interface{} {
	return a.headers
}

func (a alarm) Recipients() map[int]int {
	return a.recipients
}

func (a alarm) TemplateParameters() map[string]interface{} {
	return a.templateParameters
}

func (a alarm) buildParameterSet(alarmInformation map[string]interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	params = alarmInformation
	recipientLen := len(a.recipients)
	recipientMap := make([]int, 0, recipientLen)
	for _, recipient := range a.recipients {
		recipientMap = append(recipientMap, recipient)
	}
	recipientByte, _ := json.Marshal(recipientMap)
	params["recipients"] = string(recipientByte)
	for templateKey, templateValue := range a.templateParameters {
		params[templateKey] = templateValue
	}
	return params
}
