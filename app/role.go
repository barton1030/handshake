package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"handshake/helper"
	"handshake/service"
)

type roleController struct {
}

var RoleController roleController

type AddRoleRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Creator int    `json:"creator" form:"creator" binding:"required"`
}

type RoleByIdRequest struct {
	RoleId int `json:"role_id" form:"role_id" binding:"required"`
}

// Add 角色添加入口
func (r roleController) Add(c *gin.Context) {
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

// RoleById 通过角色id获取角色信息
func (r roleController) RoleById(c *gin.Context) {
	request := RoleByIdRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role RoleById: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}

	role, err := service.RoleService.RoleById(request.RoleId)
	if err != nil {
		err = fmt.Errorf("app role RoleById: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, role, "")
	return
}

