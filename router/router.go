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
}

func user(r *gin.Engine) {

}

func topic(r *gin.Engine) {
	topicGroup := r.Group("/topic")
	topicGroup.POST("/add", app.TopicController.Add)
}
