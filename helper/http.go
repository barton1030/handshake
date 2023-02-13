package helper

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Response(c *gin.Context, code int, data interface{}, errMsg string) {
	resp := map[string]interface{}{
		"code":  code,
		"data":  data,
		"error": errMsg,
	}
	c.JSON(0, resp)

	return
}

func ExtractRequestUri(c *gin.Context) string {
	uri := c.Request.RequestURI
	uriSlice := strings.Split(uri, "?")
	if len(uriSlice) == 0 {
		return ""
	}
	return uriSlice[0]
}

func TransformationToString(origin interface{}) string {
	target, ok := origin.(string)
	if ok {
		return target
	}
	MedianInt, ok := origin.(int)
	if ok {
		target = strconv.Itoa(MedianInt)
		return target
	}
	MedianInt64, ok := origin.(int64)
	if ok {
		target = strconv.Itoa(int(MedianInt64))
		return target
	}
	MedianInt8, ok := origin.(int8)
	if ok {
		target = strconv.Itoa(int(MedianInt8))
		return target
	}
	MedianInt32, ok := origin.(int32)
	if ok {
		target = strconv.Itoa(int(MedianInt32))
		return target
	}
	MedianBool, ok := origin.(bool)
	if ok {
		target = strconv.FormatBool(MedianBool)
		return target
	}
	return target
}
