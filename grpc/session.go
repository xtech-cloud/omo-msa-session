package grpc

import (
	"context"
	"errors"
	"omo.msa.session/cache"
	pb "omo.msa.session/grpc/proto"
)

type SessionService struct {}

func (mine *SessionService)Create(ctx context.Context, in *pb.RequestInfo, out *pb.ReplyInfo) error {
	if len(in.Uid) < 1 {
		out.Status = pb.ResultStatus_Empty
		return errors.New("the account is empty")
	}
	token := cache.CreateSession(in.Uid)
	out.Token = token
	return nil
}

func (mine *SessionService)CheckAvailable(ctx context.Context, in *pb.RequestInfo, out *pb.ReplyAvailable) error {
	if len(in.Uid) < 1 {
		out.Status = pb.ResultStatus_Empty
		return errors.New("the account uid is empty")
	}
	info := cache.GetSession(in.Uid)
	if info == nil {
		out.Status = pb.ResultStatus_NotExisted
		return errors.New("the session not found")
	}
	out.Available = info.IsExpired()
	return nil
}

func (mine *SessionService)Remove(ctx context.Context, in *pb.RequestInfo, out *pb.ReplyInfo) error {
	if len(in.Uid) < 1 {
		out.Status = pb.ResultStatus_Empty
		return errors.New("the account uid is empty")
	}
	cache.RemoveSession(in.Uid)
	return nil
}

