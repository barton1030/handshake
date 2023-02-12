package middlerware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"handshake/helper"
	"handshake/service"
)

type middleware struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
}

func PermissionVerification(c *gin.Context) {
	request := middleware{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("middleware PermissionVerification: params %v error: %v", request, err)
		helper.Response(c, 908, nil, err.Error())
		c.Abort()
	}
	uri := helper.ExtractRequestUri(c)
	err := service.PermissionVerification(request.Operator, uri)
	if err != nil {
		err = fmt.Errorf("middleware PermissionVerification: params %v error: %v", request, err)
		helper.Response(c, 909, nil, err.Error())
		c.Abort()
	}
}
