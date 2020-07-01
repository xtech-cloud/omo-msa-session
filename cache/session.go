package cache

import (
	"github.com/dgrijalva/jwt-go"
	"omo.msa.session/config"
	"time"
)

type SessionInfo struct {
	BaseInfo
	Token string
	Account string
}

func CreateSession(account string) string {
	info := GetSession(account)
	if info == nil {
		tmp := new(SessionInfo)
		tmp.Account = account
		tmp.CreateTime = time.Now()
		tmp.Token = createToken()
		cacheCtx.sessions = append(cacheCtx.sessions, tmp)
		return tmp.Token
	}else{
		return info.Token
	}
}

func GetSession(account string) *SessionInfo {
	for i :=0 ;i < len(cacheCtx.sessions);i += 1 {
		if cacheCtx.sessions[i].Account == account {
			return cacheCtx.sessions[i]
		}
	}
	return nil
}

func RemoveSession(account string) {
	for i := 0;i < len(cacheCtx.sessions);i += 1 {
		if cacheCtx.sessions[i].Account == account {
			cacheCtx.sessions = append(cacheCtx.sessions[:i], cacheCtx.sessions[i+1:]...)
			break
		}
	}
}

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString([]byte(config.Schema.Basic.Secret))
	return tokenString
}

func (mine *SessionInfo)IsExpired() bool {
	now := time.Now().Unix()
	diff := now - mine.CreateTime.Unix()
	if diff > config.Schema.Basic.Timeout {
		return false
	}else{
		return true
	}
}


