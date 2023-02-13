package domain

import (
	"fmt"
	topic2 "handshake/domain/topic"
)

func Init() {
	defer func() {
		fmt.Println("初始化结束")
	}()
	offset := 0
	limit := 10
	for {
		topicList, err := topic2.ListExample.List(offset, limit)
		if err != nil {
			continue
		}
		if len(topicList) == 0 {
			return
		}
		for _, topic := range topicList {
			currentTopic := topic
			currentTopic.StartUp()
		}
		offset += limit
		limit += limit
	}
}
