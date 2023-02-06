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
	Operator       int    `json:"operator" form:"operator" binding:"required"`
	Name           string `json:"name" form:"name" binding:"required"`
	MaxRetryCount  int    `json:"maxRetryCount" form:"maxRetryCount" binding:"required"`
	MinConcurrency int    `json:"minConcurrency" form:"minConcurrency" binding:"required"`
	MaxConcurrency int    `json:"maxConcurrency" form:"maxConcurrency" binding:"required"`
	FuseSalt       int    `json:"fuseSalt" form:"fuseSalt" binding:"required"`
}

type startTopicRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	TopicId  int `json:"topicId" form:"topicId" binding:"required"`
}

type stopTopicRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	TopicId  int `json:"topicId" form:"topicId" binding:"required"`
}

type deleteTopicRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	TopicId  int `json:"topicId" form:"topicId" binding:"required"`
}

type setTopicCallbackRequest struct {
	Operator int                    `json:"operator" form:"operator" binding:"required"`
	TopicId  int                    `json:"topicId" form:"topicId" binding:"required"`
	Url      string                 `json:"url" form:"url" binding:"required"`
	Method   string                 `json:"method" form:"method" binding:"required"`
	Headers  map[string]interface{} `json:"headers" form:"headers" binding:"required"`
	Cookies  map[string]interface{} `json:"cookies" form:"cookies" binding:"required"`
}

type setTopicAlarmRequest struct {
	Operator   int           `json:"operator" form:"operator" binding:"required"`
	TopicId    int           `json:"topicId" form:"topicId" binding:"required"`
	Url        string        `json:"url" form:"url" binding:"required"`
	Method     string        `json:"method" form:"method" binding:"required"`
	Recipients []interface{} `json:"recipients" form:"recipients" binding:"required"`
}

type editTopicRequest struct {
	Operator       int `json:"operator" form:"operator" binding:"required"`
	TopicId        int `json:"topicId" form:"topicId" binding:"required"`
	MaxRetryCount  int `json:"maxRetryCount" form:"maxRetryCount" binding:"required"`
	MinConcurrency int `json:"minConcurrency" form:"minConcurrency" binding:"required"`
	MaxConcurrency int `json:"maxConcurrency" form:"maxConcurrency" binding:"required"`
	FuseSalt       int `json:"fuseSalt" form:"fuseSalt" binding:"required"`
}

type topicByIdRequest struct {
	Operator int `json:"operator" form:"operator" binding:"required"`
	TopicId  int `json:"topicId" form:"topicId" binding:"required"`
}

type topicPushMessageRequest struct {
	Operator int                    `json:"operator" form:"operator" binding:"required"`
	TopicId  int                    `json:"topicId" form:"topicId" binding:"required"`
	Message  map[string]interface{} `json:"message" form:"message" binding:"required"`
}

func (t topic) Add(c *gin.Context) {
	request := addTopicRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic Add: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.TopicService.Add(request.Operator, request.Name, request.MaxRetryCount, request.MaxConcurrency, request.MinConcurrency, request.FuseSalt)
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
	err := service.TopicService.SetCallback(request.Operator, request.TopicId, request.Url, request.Method, request.Headers, request.Cookies)
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
	err := service.TopicService.Delete(request.Operator, request.TopicId)
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
	err := service.TopicService.Start(request.Operator, request.TopicId)
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
	err := service.TopicService.Stop(request.Operator, request.TopicId)
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
	err := service.TopicService.SetAlarm(request.Operator, request.TopicId, request.Url, request.Method, request.Recipients)
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
	err := service.TopicService.Edit(request.Operator, request.TopicId, request.MaxRetryCount, request.MinConcurrency, request.MaxConcurrency, request.FuseSalt)
	if err != nil {
		err = fmt.Errorf("app topic SetTopic: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}

// TopicById 编辑主题信息
func (t topic) TopicById(c *gin.Context) {
	request := topicByIdRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic TopicById: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	topic, err := service.TopicService.TopicById(request.Operator, request.TopicId)
	if err != nil {
		err = fmt.Errorf("app topic TopicById: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, topic, "")
	return
}

// PushMessage 推送消息
func (t topic) PushMessage(c *gin.Context) {
	request := topicPushMessageRequest{}
	if err := c.ShouldBind(&request); err != nil {
		err = fmt.Errorf("app topic PushMessage: params %v error: %v", request, err)
		helper.Response(c, 1000, nil, err.Error())
		return
	}
	err := service.TopicService.PushMessage(request.Operator, request.TopicId, request.Message)
	if err != nil {
		err = fmt.Errorf("app topic PushMessage: params %v error: %v", request, err)
		helper.Response(c, 1001, nil, err.Error())
		return
	}
	helper.Response(c, 200, nil, "")
	return
}
