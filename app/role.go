package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"handshake/helper"
	"handshake/service"
)

type role struct {
}

var RoleController role

type AddRoleRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Creator int    `json:"creator" form:"creator" binding:"required"`
}

func (r role) Add(c *gin.Context) {
	request := AddRoleRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role Add: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.RoleService.Add(request.Name, request.Creator)
	if err != nil {
		err = fmt.Errorf("app role Add: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}
