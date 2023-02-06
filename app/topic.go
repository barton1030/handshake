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

type stopTopicRequest struct {
	TopicId int `json:"topicId" form:"topicId" binding:"required"`
}

type deleteTopicRequest struct {
	TopicId int `json:"topicId" form:"topicId" binding:"required"`
}

type setTopicCallbackRequest struct {
	TopicId int                    `json:"topicId" form:"topicId" binding:"required"`
	Url     string                 `json:"url" form:"url" binding:"required"`
	Method  string                 `json:"method" form:"method" binding:"required"`
	Headers map[string]interface{} `json:"headers" form:"headers" binding:"required"`
	Cookies map[string]interface{} `json:"cookies" form:"cookies" binding:"required"`
}

type setTopicAlarmRequest struct {
	TopicId    int           `json:"topicId" form:"topicId" binding:"required"`
	Url        string        `json:"url" form:"url" binding:"required"`
	Method     string        `json:"method" form:"method" binding:"required"`
	Recipients []interface{} `json:"recipients" form:"recipients" binding:"required"`
}

type editTopicRequest struct {
	TopicId        int `json:"topicId" form:"topicId" binding:"required"`
	MaxRetryCount  int `json:"maxRetryCount" form:"maxRetryCount" binding:"required"`
	MinConcurrency int `json:"minConcurrency" form:"minConcurrency" binding:"required"`
	MaxConcurrency int `json:"maxConcurrency" form:"maxConcurrency" binding:"required"`
	FuseSalt       int `json:"fuseSalt" form:"fuseSalt" binding:"required"`
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

func (t topic) SetCallback(c *gin.Context) {
	request := setTopicCallbackRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic SetCallback: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.TopicService.SetCallback(request.TopicId, request.Url, request.Method, request.Headers, request.Cookies)
	if err != nil {
		err = fmt.Errorf("app topic SetCallback: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

func (t topic) Delete(c *gin.Context) {
	request := deleteTopicRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic Delete: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.TopicService.Delete(request.TopicId)
	if err != nil {
		err = fmt.Errorf("app topic Delete: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
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

func (t topic) Stop(c *gin.Context) {
	request := stopTopicRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic Stop: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.TopicService.Stop(request.TopicId)
	if err != nil {
		err = fmt.Errorf("app topic Stop: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

func (t topic) SetAlarm(c *gin.Context) {
	request := setTopicAlarmRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic SetAlarm: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.TopicService.SetAlarm(request.TopicId, request.Url, request.Method, request.Recipients)
	if err != nil {
		err = fmt.Errorf("app topic SetAlarm: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

// EditTopic 编辑主题信息
func (t topic) EditTopic(c *gin.Context) {
	request := editTopicRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic SetTopic: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.TopicService.Edit(request.TopicId, request.MaxRetryCount, request.MinConcurrency, request.MaxConcurrency, request.FuseSalt)
	if err != nil {
		err = fmt.Errorf("app topic SetTopic: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}
