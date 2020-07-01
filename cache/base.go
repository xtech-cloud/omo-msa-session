package cache

import (
	"time"
)

type BaseInfo struct {
	ID         uint64 `json:"-"`
	UID        string `json:"uid"`
	Name       string `json:"name"`
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

	//err := nosql.InitDB(config.Schema.Database.IP, config.Schema.Database.Port, config.Schema.Database.Name, config.Schema.Database.Type)
	//if nil != err {
	//	return err
	//}

	return nil
}

