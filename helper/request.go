package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	httpUrl "net/url"
	"strings"
)

type request struct {
}

var R request

func (r request) HttpRequest(url, method string, headers, cookies, params map[string]interface{}) (resp map[string]interface{}, err error) {
	if method != "GET" {
		resp, err = r.httpPost(url, method, headers, cookies, params)
		return
	}
	resp, err = r.httpGet(url, method, headers, cookies, params)
	return
}

func (r request) httpGet(url, method string, headers, cookies, params map[string]interface{}) (resp map[string]interface{}, err error) {
	client := http.Client{}
	paramsSpliceSalt := r.queryParmaSplice(params)
	url += "?" + paramsSpliceSalt
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	r.setHeaders(req, headers)
	r.setCookies(req, cookies)
	// 执行
	response, err := client.Do(req)
	defer func() {
		if response == nil {
			return
		}
		response.Body.Close()
	}()
	if response == nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	// 解析
	resp = make(map[string]interface{})
	data, err := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &resp)

	return
}

func (r request) httpPost(url, method string, headers, cookies, params map[string]interface{}) (resp map[string]interface{}, err error) {
	client := http.Client{}
	ioBody := r.setBody(params)
	req, err := http.NewRequest(method, url, ioBody)
	if err != nil {
		return
	}
	r.setHeaders(req, headers)
	// 默认使用
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.setCookies(req, cookies)
	// 执行
	response, err := client.Do(req)
	defer func() {
		if response == nil {
			return
		}
		response.Body.Close()
	}()
	if response == nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	// 解析
	resp = make(map[string]interface{})
	data, err := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &resp)
	return
	return
}

func (r request) setHeaders(req *http.Request, headers map[string]interface{}) {
	for headerKey, headerValue := range headers {
		targetValue := TransformationToString(headerValue)
		req.Header.Add(headerKey, targetValue)
	}
}

func (r request) setCookies(req *http.Request, cookies map[string]interface{}) {
	for cookieKey, cookieValue := range cookies {
		targetValue := TransformationToString(cookieValue)
		cookie := &http.Cookie{Name: cookieKey, Value: targetValue}
		req.AddCookie(cookie)
	}
}

func (r request) setBody(params map[string]interface{}) (ioBody *strings.Reader) {
	values := httpUrl.Values{}
	for paramKey, paramValue := range params {
		targetValue := TransformationToString(paramValue)
		values.Set(paramKey, targetValue)
	}
	ioBody = strings.NewReader(values.Encode())
	return
}

func (r request) queryParmaSplice(params map[string]interface{}) string {
	paramSpliceSalt := ""
	paramsLength := len(params)
	counter := 0
	for key, value := range params {
		targetValue := TransformationToString(value)
		paramsUnit := key + "=" + targetValue
		paramSpliceSalt += paramsUnit
		counter++
		if counter == paramsLength {
			break
		}
		paramSpliceSalt += "&"
	}
	return paramSpliceSalt
}
