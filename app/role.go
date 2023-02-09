package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"handshake/helper"
	"handshake/service"
)

type role struct {
}

var Role role

type AddRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Operator int    `json:"operator" form:"operator" binding:"required"`
}

type RoleByIdRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	RoleId   int `json:"roleId" form:"roleId" binding:"required"`
}

type EditNameRequest struct {
	Operator int    `json:"operator" form:"operator" binding:"required"`
	RoleId   int    `json:"roleId" form:"roleId" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
}

type SetPermissionRequest struct {
	Operator        int    `json:"operator" form:"operator" binding:"required"`
	RoleId          int    `json:"roleId" form:"roleId" binding:"required"`
	PermissionKey   string `json:"permissionKey" form:"permissionKey" binding:"required"`
	PermissionValue bool   `json:"permissionValue" form:"permissionValue"`
}

type ListRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	Offset   int `json:"offset" form:"offset"`
	Limit    int `json:"limit" form:"limit" binding:"required"`
}

type DeleteRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	RoleId   int `json:"roleId" form:"roleId" binding:"required"`
}

// Add 角色添加入口
func (r role) Add(c *gin.Context) {
	request := AddRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role Add: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.Role.Add(request.Operator, request.Name, uri)
	if err != nil {
		err = fmt.Errorf("app role Add: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

// RoleById 通过角色id获取角色信息
func (r role) RoleById(c *gin.Context) {
	request := RoleByIdRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role RoleById: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	role2, err := service.Role.RoleById(request.Operator, request.RoleId, uri)
	if err != nil {
		err = fmt.Errorf("app role RoleById: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, role2, "")
	return
}

// EditName 编辑角色名称
func (r role) EditName(c *gin.Context) {
	request := EditNameRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role EditName: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.Role.EditName(request.Operator, request.RoleId, request.Name, uri)
	if err != nil {
		err = fmt.Errorf("app role EditName: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

// SetPermission 设置角色相关权限
func (r role) SetPermission(c *gin.Context) {
	request := SetPermissionRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role SetPermission: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.Role.SetPermission(request.Operator, request.RoleId, request.PermissionKey, request.PermissionValue, uri)
	if err != nil {
		err = fmt.Errorf("app role SetPermission: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

// List 角色列表
func (r role) List(c *gin.Context) {
	request := ListRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role List: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	roles, err := service.Role.List(request.Operator, request.Offset, request.Limit, uri)
	if err != nil {
		err = fmt.Errorf("app role List: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, roles, "")
	return
}

// Delete 删除角色
func (r role) Delete(c *gin.Context) {
	request := DeleteRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role Delete: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.Role.Delete(request.Operator, request.RoleId, uri)
	if err != nil {
		err = fmt.Errorf("app role Delete: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}
