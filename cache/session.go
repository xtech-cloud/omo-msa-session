package cache

import (
	"github.com/dgrijalva/jwt-go"
	"omo.msa.session/config"
	"time"
)

type SessionInfo struct {
	CreateTime time.Time
	User       string
	Token      *jwt.Token
}

type ClaimsInfo struct {
	//User string `json:"user"`
	jwt.StandardClaims
}

func CreateSession(user string) (string, error) {
	token, err := createToken(user)
	if err != nil {
		return "", err
	}
	//RemoveSession(user)
	//tmp := new(SessionInfo)
	//tmp.User = user
	//tmp.CreateTime = time.Now()
	//tmp.Token = token
	//cacheCtx.sessions = append(cacheCtx.sessions, tmp)
	//return tmp.Token.Raw
	return token.Raw, nil
}

//func GetSessionByUser(user string) *SessionInfo {
//	for i :=0 ;i < len(cacheCtx.sessions);i += 1 {
//		if cacheCtx.sessions[i].User == user {
//			return cacheCtx.sessions[i]
//		}
//	}
//	return nil
//}
//
//func GetSessionByToken(msg string) *SessionInfo {
//	for i :=0 ;i < len(cacheCtx.sessions);i += 1 {
//		if cacheCtx.sessions[i].Raw() == msg {
//			return cacheCtx.sessions[i]
//		}
//	}
//	return nil
//}

//func RemoveSession(user string) {
//	for i := 0;i < len(cacheCtx.sessions);i += 1 {
//		if cacheCtx.sessions[i].User == user {
//			cacheCtx.sessions = append(cacheCtx.sessions[:i], cacheCtx.sessions[i+1:]...)
//			break
//		}
//	}
//}

func createToken(user string) (*jwt.Token, error) {
	now := time.Now()
	expire := now.Add(time.Second * time.Duration(config.Schema.Basic.Timeout))
	info := ClaimsInfo{
		//User: user,
		StandardClaims: jwt.StandardClaims{
			Id:        user,
			ExpiresAt: expire.Unix(),
			Issuer:    "yumei",
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	msg, err := token.SignedString([]byte(config.Schema.Basic.Secret))
	token.Raw = msg
	return token, err
}

func generateToken(user string) string {
	token, err := createToken(user)
	if err != nil {
		return ""
	}
	return token.Raw
}

func ParseToken(token string) (*ClaimsInfo, error) {
	claims, err := jwt.ParseWithClaims(token, &ClaimsInfo{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Schema.Basic.Secret), nil
	})
	if claims != nil {
		info, ok := claims.Claims.(*ClaimsInfo)
		if ok && claims.Valid {
			return info, nil
		}
	}
	return nil, err
}

func (mine *ClaimsInfo) CheckNew() string {
	now := time.Now().Unix()
	rest := now - mine.StandardClaims.ExpiresAt
	if rest > 0 {
		return ""
	} else {
		if rest < -1800 {
			return ""
		} else {
			return generateToken(mine.Id)
		}
	}
}

func (mine *SessionInfo) IsExpired() bool {
	now := time.Now().Unix()
	diff := now - mine.CreateTime.Unix()
	if diff > config.Schema.Basic.Timeout {
		return true
	} else {
		mine.CreateTime = time.Now()
		return false
	}
}

func (mine *SessionInfo) UpdateTime() {
	mine.CreateTime = time.Now()
}

func (mine *SessionInfo) Raw() string {
	return mine.Token.Raw
}
