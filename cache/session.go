package cache

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"omo.msa.session/config"
	"time"
)

type SessionInfo struct {
	BaseInfo
	User  string
	Token *jwt.Token
}

func CreateSession(user string) string {
	RemoveSession(user)
	tmp := new(SessionInfo)
	tmp.User = user
	tmp.CreateTime = time.Now()
	tmp.Token = createToken(user)
	cacheCtx.sessions = append(cacheCtx.sessions, tmp)
	return tmp.Token.Raw
}

func GetSessionByUser(user string) *SessionInfo {
	for i :=0 ;i < len(cacheCtx.sessions);i += 1 {
		if cacheCtx.sessions[i].User == user {
			return cacheCtx.sessions[i]
		}
	}
	return nil
}

func GetSessionByToken(msg string) *SessionInfo {
	for i :=0 ;i < len(cacheCtx.sessions);i += 1 {
		if cacheCtx.sessions[i].TokenString() == msg {
			return cacheCtx.sessions[i]
		}
	}
	return nil
}

func RemoveSession(user string) {
	for i := 0;i < len(cacheCtx.sessions);i += 1 {
		if cacheCtx.sessions[i].User == user {
			cacheCtx.sessions = append(cacheCtx.sessions[:i], cacheCtx.sessions[i+1:]...)
			break
		}
	}
}

func createToken(uid string) *jwt.Token {
	token := jwt.New(jwt.SigningMethodHS256)
	msg, _ := token.SignedString([]byte(uid+config.Schema.Basic.Secret))
	token.Raw = msg
	return token
}

func (mine *SessionInfo)IsExpired() bool {
	now := time.Now().Unix()
	diff := now - mine.CreateTime.Unix()
	if diff > config.Schema.Basic.Timeout {
		return true
	}else{
		mine.CreateTime = time.Now()
		return false
	}
}

func (mine *SessionInfo) UpdateTime()  {
	mine.CreateTime = time.Now()
}

func (mine *SessionInfo)TokenString() string {
	return mine.Token.Raw
}

func ParseToken(msg, uid string) (*jwt.Token,error) {
	token, err := jwt.Parse(msg, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(uid+config.Schema.Basic.Secret), nil
	})
	return token, err
}


