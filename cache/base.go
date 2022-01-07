package cache

type cacheContext struct {
	//sessions  []*SessionInfo
}

var cacheCtx *cacheContext

func InitData() error {
	cacheCtx = &cacheContext{}
	//cacheCtx.sessions = make([]*SessionInfo, 0, 1000)
	return nil
}

