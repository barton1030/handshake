package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"handshake/helper"
	"handshake/service"
)

type user struct {
}

var User user

type userAddRequest struct {
	Operator int    `json:"operator" form:"operator" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Pwd      string `json:"pwd" form:"pwd" binding:"required"`
	RoleId   int    `json:"roleId" form:"roleId" binding:"required"`
}

type setRoleIdRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	UserId   int `json:"userId" form:"userId" binding:"required"`
	RoleId   int `json:"roleId" form:"roleId" binding:"required"`
}

type deleteUserRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	UserId   int `json:"userId" form:"userId" binding:"required"`
}

type listRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	Offset   int `json:"offset" form:"offset"`
	Limit    int `json:"limit" form:"limit" binding:"required"`
}

type userIdRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	UserId   int `json:"userId" form:"userId" binding:"required"`
}

func (u user) Add(c *gin.Context) {
	request := userAddRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app user Add: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.User.Add(request.Operator, request.RoleId, request.Name, request.Phone, request.Pwd, uri)
	if err != nil {
		err = fmt.Errorf("app user Add: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

func (u user) SetRoleId(c *gin.Context) {
	request := setRoleIdRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app user SetRoleId: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.User.SetRoleId(request.Operator, request.UserId, request.RoleId, uri)
	if err != nil {
		err = fmt.Errorf("app user SetRoleId: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

func (u user) Delete(c *gin.Context) {
	request := deleteUserRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app user Delete: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.User.Delete(request.Operator, request.UserId, uri)
	if err != nil {
		err = fmt.Errorf("app user Delete: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

func (u user) List(c *gin.Context) {
	request := listRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app user List: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	userList, err := service.User.List(request.Operator, request.Offset, request.Limit, uri)
	if err != nil {
		err = fmt.Errorf("app user List: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, userList, "")
	return
}

func (u user) UserById(c *gin.Context) {
	request := userIdRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app user UserId: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	user2, err := service.User.UserId(request.Operator, request.UserId, uri)
	if err != nil {
		err = fmt.Errorf("app user UserId: params %v error: %v",
			request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, user2, "")
	return
}
