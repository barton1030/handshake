package app

import (
	"github.com/gin-gonic/gin"
	"handshake/helper"
)

type topic struct {
}

var TopicController topic

func (t topic) Add(c *gin.Context) {
	helper.Response(c, 0, nil, "")
}

func (t topic) Edit(c *gin.Context) {

}

func (t topic) Delete(c *gin.Context) {

}
