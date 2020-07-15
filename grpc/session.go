package grpc

import (
	"context"
	"errors"
	pb "github.com/xtech-cloud/omo-msp-session/proto/session"
	"omo.msa.session/cache"
)

type SessionService struct {}

func (mine *SessionService)Create(ctx context.Context, in *pb.ReqSessionAdd, out *pb.ReplyInfo) error {
	if len(in.User) < 1 {
		out.Status = pb.ResultStatus_Empty
		return errors.New("the account is empty")
	}
	token := cache.CreateSession(in.User)
	out.Token = token
	return nil
}

func (mine *SessionService)CheckAvailable(ctx context.Context, in *pb.RequestInfo, out *pb.ReplyAvailable) error {
	if len(in.Token) < 1 {
		out.Status = pb.ResultStatus_Empty
		return errors.New("the account uid is empty")
	}
	token,err := cache.ParseToken(in.Token)
	if err != nil {
		out.Status = pb.ResultStatus_DBException
		return err
	}
	if !token.Valid {
		out.Status = pb.ResultStatus_DBException
		return errors.New("the token valid failed")
	}
	info := cache.GetSessionByToken(in.Token)
	if info == nil {
		out.Status = pb.ResultStatus_NotExisted
		return errors.New("the session not found")
	}
	out.User = info.User
	out.Available = info.IsExpired()
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

