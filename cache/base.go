package cache

import (
	"time"
)

type BaseInfo struct {
	ID         uint64
	CreateTime time.Time
	UpdateTime time.Time
}

type cacheContext struct {
	sessions  []*SessionInfo
}

var cacheCtx *cacheContext

func InitData() error {
	cacheCtx = &cacheContext{}
	cacheCtx.sessions = make([]*SessionInfo, 0, 1000)
	return nil
}

