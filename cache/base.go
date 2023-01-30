package cache

type cacheContext struct {
	//sessions  []*SessionInfo
	outTokens map[string]string
}

var cacheCtx *cacheContext

func InitData() error {
	cacheCtx = &cacheContext{}
	//cacheCtx.sessions = make([]*SessionInfo, 0, 1000)
	cacheCtx.outTokens = make(map[string]string, 1000)
	return nil
}

func Content() *cacheContext {
	return cacheCtx
}

func (mine *cacheContext) SignOut(user, token string) {
	mine.outTokens[user] = token
}

func (mine *cacheContext) IsSignOut(user string) bool {
	_, ok := mine.outTokens[user]
	return ok
}

func (mine *cacheContext) Remove(user string) {
	_, ok := mine.outTokens[user]
	if ok {
		delete(mine.outTokens, user)
	}
}
