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
	roleGroup.POST("/add", app.Role.Add)
	roleGroup.GET("/byId", app.Role.RoleById)
	roleGroup.POST("/set/name", app.Role.EditName)
	roleGroup.POST("/set/permission", app.Role.SetPermission)
	roleGroup.GET("/list", app.Role.List)
	roleGroup.GET("/delete", app.Role.Delete)
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
	topicGroup.POST("/add", app.Topic.Add)
	topicGroup.GET("/delete", app.Topic.Delete)
	topicGroup.POST("/start", app.Topic.Start)
	topicGroup.POST("/stop", app.Topic.Stop)
	topicGroup.POST("/set/callback", app.Topic.SetCallback)
	topicGroup.POST("/set/alarm", app.Topic.SetAlarm)
	topicGroup.POST("/edit", app.Topic.EditTopic)
	topicGroup.GET("/byId", app.Topic.TopicById)
	topicGroup.POST("/push/message", app.Topic.PushMessage)
}
