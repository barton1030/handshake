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
	roleGroup.GET("/byId", app.RoleController.RoleById)
	roleGroup.POST("/set/name", app.RoleController.EditName)
	roleGroup.POST("/set/permission", app.RoleController.SetPermission)
	roleGroup.GET("/list", app.RoleController.List)
	roleGroup.GET("/delete", app.RoleController.Delete)
}

func user(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("/add", app.User.Add)
	userGroup.POST("/set/roleId", app.User.SetRoleId)
	userGroup.GET("/delete", app.User.Delete)
	userGroup.GET("/list", app.User.List)
	userGroup.GET("/byId", app.User.UserById)
}

func topic(r *gin.Engine) {
	topicGroup := r.Group("/topic")
	topicGroup.POST("/add", app.TopicController.Add)
	topicGroup.GET("/delete", app.TopicController.Delete)
	topicGroup.POST("/start", app.TopicController.Start)
	topicGroup.POST("/stop", app.TopicController.Stop)
	topicGroup.POST("/set/callback", app.TopicController.SetCallback)
	topicGroup.POST("/set/alarm", app.TopicController.SetAlarm)
	topicGroup.POST("/edit", app.TopicController.EditTopic)
	topicGroup.GET("/byId", app.TopicController.TopicById)
	topicGroup.POST("/push/message", app.TopicController.PushMessage)
}
