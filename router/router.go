package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	root := gin.Default()
	user(root)
	return root
}

func user(r *gin.Engine) {

}
