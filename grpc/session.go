package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/v2/logger"
	pb "github.com/xtech-cloud/omo-msp-session/proto/session"
	"omo.msa.session/cache"
)

type SessionService struct {}

func inLog(name, data interface{})  {
	bytes, _ := json.Marshal(data)
	msg := ByteString(bytes)
	logger.Infof("[request.%s]:data = %s", name, msg)
}

func outLog(name, data interface{})  {
	bytes, _ := json.Marshal(data)
	msg := ByteString(bytes)
	logger.Infof("[response.%s]:data = %s", name, msg)
}

func ByteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func (mine *SessionService)Create(ctx context.Context, in *pb.ReqSessionAdd, out *pb.ReplyInfo) error {
	path := "session.create"
	inLog(path, in)
	if len(in.User) < 1 {
		out.Status = pb.ResultStatus_Empty
		return errors.New("the account is empty")
	}
	token := cache.CreateSession(in.User)
	out.Token = token
	outLog(path, in)
	return nil
}

func (mine *SessionService)CheckAvailable(ctx context.Context, in *pb.RequestInfo, out *pb.ReplyAvailable) error {
	path := "session.available"
	inLog(path, in)
	if len(in.Token) < 1 {
		out.Status = pb.ResultStatus_Empty
		return errors.New("the account uid is empty")
	}
	info := cache.GetSessionByToken(in.Token)
	if info == nil {
		out.Status = pb.ResultStatus_NotExisted
		return errors.New("the session not found")
	}
	//token,err := cache.ParseToken(info.User, in.Token)
	//if err != nil {
	//	out.Status = pb.ResultStatus_DBException
	//	return err
	//}
	//if !token.Valid {
	//	out.Status = pb.ResultStatus_DBException
	//	return errors.New("the token valid failed")
	//}

	out.User = info.User
	out.Available = info.IsExpired()
	outLog(path, in)
	return nil
}

func (mine *SessionService)Remove(ctx context.Context, in *pb.ReqSessionRemove, out *pb.ReplyInfo) error {
	if len(in.User) < 1 {
		out.Status = pb.ResultStatus_Empty
		return errors.New("the account uid is empty")
	}
	info := cache.GetSessionByUser(in.User)
	if info == nil {
		out.Status = pb.ResultStatus_NotExisted
		return errors.New("the session not found")
	}
	out.User = in.User
	cache.RemoveSession(info.User)
	return nil
}

