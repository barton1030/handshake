package helper

import (
	"github.com/gin-gonic/gin"
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
