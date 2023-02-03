package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"handshake/helper"
	"handshake/service"
)

type topic struct {
}

var TopicController topic

type addTopicRequest struct {
	Name           string `json:"name" form:"name" binding:"required"`
	MaxRetryCount  int    `json:"maxRetryCount" form:"maxRetryCount" binding:"required"`
	MinConcurrency int    `json:"minConcurrency" form:"minConcurrency" binding:"required"`
	MaxConcurrency int    `json:"maxConcurrency" form:"maxConcurrency" binding:"required"`
	FuseSalt       int    `json:"fuseSalt" form:"fuseSalt" binding:"required"`
	Creator        int    `json:"creator" form:"creator" binding:"required"`
}

type startTopicRequest struct {
	TopicId int `json:"topicId" form:"topicId" binding:"required"`
}

func (t topic) Add(c *gin.Context) {
	request := addTopicRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic Add: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.TopicService.Add(request.Name, request.MaxRetryCount, request.MaxConcurrency, request.MinConcurrency, request.FuseSalt, request.Creator)
	if err != nil {
		err = fmt.Errorf("app topic Add: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

func (t topic) Edit(c *gin.Context) {

}

func (t topic) Delete(c *gin.Context) {

}

func (t topic) Start(c *gin.Context) {
	request := startTopicRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic Start: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.TopicService.Start(request.TopicId)
	if err != nil {
		err = fmt.Errorf("app topic Start: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}
