package grpc

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/v2/logger"
	pb "github.com/xtech-cloud/omo-msp-session/proto/session"
	"omo.msa.session/cache"
)

type SessionService struct {}

func inLog(name, data interface{})  {
	bytes, _ := json.Marshal(data)
	msg := ByteString(bytes)
	logger.Infof("[in.%s]:data = %s", name, msg)
}

func outLog(name, data interface{}) *pb.ReplyStatus {
	bytes, _ := json.Marshal(data)
	msg := ByteString(bytes)
	logger.Infof("[out.%s]:data = %s", name, msg)
	tmp := &pb.ReplyStatus{
		Code: 0,
		Message: "",
	}
	return tmp
}

func outError(name, msg string, code pb.ResultStatus) *pb.ReplyStatus {
	logger.Warnf("[error.%s]:code = %d, msg = %s", name, code, msg)
	tmp := &pb.ReplyStatus{
		Code: code,
		Message: msg,
	}
	return tmp
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
		out.Status = outError(path, "the account is empty",pb.ResultStatus_Empty)
		return nil
	}
	token := cache.CreateSession(in.User)
	out.Token = token
	out.Status = outLog(path, out)
	return nil
}

func (mine *SessionService)CheckAvailable(ctx context.Context, in *pb.RequestInfo, out *pb.ReplyAvailable) error {
	path := "session.available"
	inLog(path, in)
	if len(in.Token) < 1 {
		out.Status = outError(path, "the token is empty",pb.ResultStatus_Empty)
		return nil
	}
	info := cache.GetSessionByToken(in.Token)
	if info == nil {
		out.Status = outError(path, "the token not exited",pb.ResultStatus_NotExisted)
		return nil
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
	out.Available = !info.IsExpired()
	out.Status = outLog(path, out)
	return nil
}

func (mine *SessionService)Remove(ctx context.Context, in *pb.ReqSessionRemove, out *pb.ReplyInfo) error {
	path := "session.remove"
	if len(in.User) < 1 {
		out.Status = outError(path, "the user not empty",pb.ResultStatus_Empty)
		return nil
	}
	info := cache.GetSessionByUser(in.User)
	if info == nil {
		out.Status = outError(path, "the session not found",pb.ResultStatus_NotExisted)
		return nil
	}
	out.User = in.User
	cache.RemoveSession(info.User)
	out.Status = outLog(path, out)
	return nil
}

