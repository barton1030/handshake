package engine

import (
	"sync"
)

// 不只是对象需要实现单一职责，变量及其组成元素也应该都有自己的单一且明确的意义
type conduit struct {
	error          map[string]chan int
	errorLock      sync.Mutex
	statistics     map[string]chan int
	statisticsLock sync.Mutex
	fusing         map[string]chan int
	fusingLock     sync.Mutex
}

var conduitUnit conduit

func conduitInit() {
	conduitUnit.error = make(map[string]chan int)
	conduitUnit.statistics = make(map[string]chan int)
	conduitUnit.fusing = make(map[string]chan int)
}

func (c *conduit) setUpErrorConduit(name string, conduitCap int) {
	c.errorLock.Lock()
	defer c.errorLock.Unlock()
	if _, ok := c.error[name]; ok {
		return
	}
	c.error[name] = make(chan int, conduitCap)
}

func (c *conduit) setUpStatisticsConduit(name string, conduitCap int) {
	c.statisticsLock.Lock()
	defer c.statisticsLock.Unlock()
	if _, ok := c.statistics[name]; ok {
		return
	}
	c.statistics[name] = make(chan int, conduitCap)
}

func (c *conduit) setUpFusingConduit(name string, conduitCap int) {
	c.fusingLock.Lock()
	defer c.fusingLock.Unlock()
	if _, ok := c.fusing[name]; ok {
		return
	}
	c.fusing[name] = make(chan int, conduitCap)
}

func (c *conduit) errorConduitByName(name string) (errorConduit *chan int) {
	if cErrorConduit, ok := c.error[name]; ok {
		errorConduit = &cErrorConduit
	}
	return
}

func (c *conduit) statisticsConduitByName(name string) (statisticsConduit *chan int) {
	if cStatisticsConduit, ok := c.error[name]; ok {
		statisticsConduit = &cStatisticsConduit
	}
	return
}

func (c *conduit) fusingConduitByName(name string) (fusingConduit *chan int) {
	if cFusingConduit, ok := c.error[name]; ok {
		fusingConduit = &cFusingConduit
	}
	return
}