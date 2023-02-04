package router

import (
	"github.com/gin-gonic/gin"
	"handshake/app"
)

func Router() *gin.Engine {
	root := gin.Default()
	role(root)
	user(root)
	topic(root)
	return root
}

func role(r *gin.Engine) {
	roleGroup := r.Group("/role")
	roleGroup.POST("/add", app.RoleController.Add)
	roleGroup.GET("/id", app.RoleController.RoleById)
	roleGroup.GET("/name/edit", app.RoleController.EditName)
	roleGroup.GET("/permission/edit", app.RoleController.SetPermission)
	roleGroup.GET("/list", app.RoleController.List)
	roleGroup.GET("/delete", app.RoleController.Delete)
}

func user(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("/add", app.User.Add)
	userGroup.POST("/edit", app.User.SetRoleId)
	userGroup.GET("/delete", app.User.Delete)
	userGroup.GET("/list", app.User.List)
	userGroup.GET("/userId", app.User.UserId)
}

func topic(r *gin.Engine) {
	topicGroup := r.Group("/topic")
	topicGroup.POST("/add", app.TopicController.Add)
	topicGroup.POST("/start", app.TopicController.Start)
	topicGroup.POST("/stop", app.TopicController.Stop)
}
