package helper

import (
	"github.com/gin-gonic/gin"
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
