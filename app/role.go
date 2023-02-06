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

type AddRequest struct {
	Name   string `json:"name" form:"name" binding:"required"`
	UserId int    `json:"UserId" form:"UserId" binding:"required"`
}

type RoleByIdRequest struct {
	UserId int `json:"UserId" form:"UserId" binding:"required"`
	RoleId int `json:"roleId" form:"roleId" binding:"required"`
}

type EditNameRequest struct {
	UserId int    `json:"UserId" form:"UserId" binding:"required"`
	RoleId int    `json:"roleId" form:"roleId" binding:"required"`
	Name   string `json:"name" form:"name" binding:"required"`
}

type SetPermissionRequest struct {
	UserId          int    `json:"UserId" form:"UserId" binding:"required"`
	RoleId          int    `json:"roleId" form:"roleId" binding:"required"`
	PermissionKey   string `json:"permissionKey" form:"permissionKey" binding:"required"`
	PermissionValue bool   `json:"permissionValue" form:"permissionValue"`
}

type ListRequest struct {
	UserId int `json:"UserId" form:"UserId" binding:"required"`
	Offset int `json:"offset" form:"offset"`
	Limit  int `json:"limit" form:"limit" binding:"required"`
}

type DeleteRequest struct {
	UserId int `json:"UserId" form:"UserId" binding:"required"`
	RoleId int `json:"roleId" form:"roleId" binding:"required"`
}

// Add 角色添加入口
func (r roleController) Add(c *gin.Context) {
	request := AddRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role Add: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.RoleService.Add(uri, request.Name, request.UserId)
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
	uri := helper.ExtractRequestUri(c)
	role, err := service.RoleService.RoleById(request.UserId, request.RoleId, uri)
	if err != nil {
		err = fmt.Errorf("app role RoleById: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, role, "")
	return
}

// EditName 编辑角色名称
func (r roleController) EditName(c *gin.Context) {
	request := EditNameRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role EditName: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.RoleService.EditName(request.UserId, request.RoleId, request.Name, uri)
	if err != nil {
		err = fmt.Errorf("app role EditName: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

// SetPermission 设置角色相关权限
func (r roleController) SetPermission(c *gin.Context) {
	request := SetPermissionRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role SetPermission: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.RoleService.SetPermission(request.UserId, request.RoleId, request.PermissionKey, request.PermissionValue, uri)
	if err != nil {
		err = fmt.Errorf("app role SetPermission: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

// List 角色列表
func (r roleController) List(c *gin.Context) {
	request := ListRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role List: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	roles, err := service.RoleService.List(request.UserId, request.Offset, request.Limit, uri)
	if err != nil {
		err = fmt.Errorf("app role List: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, roles, "")
	return
}

// Delete 删除角色
func (r roleController) Delete(c *gin.Context) {
	request := DeleteRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app role Delete: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	uri := helper.ExtractRequestUri(c)
	err := service.RoleService.Delete(request.UserId, request.RoleId, uri)
	if err != nil {
		err = fmt.Errorf("app role Delete: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}
