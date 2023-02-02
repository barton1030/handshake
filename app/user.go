package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"handshake/helper"
	"handshake/service"
)

type userController struct {
}

var User userController

type userAddRequest struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Phone  string `json:"phone" form:"name" binding:"required"`
	Pwd    string `json:"pwd" form:"pwd" binding:"required"`
	RoleId int    `json:"role_id" form:"role_id" binding:"required"`
}

type setRoleIdRequest struct {
	UserId int `json:"user_id" form:"user_id" binding:"required"`
	RoleId int `json:"role_id" form:"role_id" binding:"required"`
}

func (u userController) Add(c *gin.Context) {
	request := userAddRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app user Add: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}

	err := service.UserService.Add(request.Name, request.Phone, request.Pwd, request.RoleId)
	if err != nil {
		err = fmt.Errorf("app user Add: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

func (u userController) SetRoleId(c *gin.Context) {
	request := setRoleIdRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app user SetRoleId: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}

	err := service.UserService.SetRoleId(request.UserId, request.RoleId)
	if err != nil {
		err = fmt.Errorf("app user SetRoleId: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}
