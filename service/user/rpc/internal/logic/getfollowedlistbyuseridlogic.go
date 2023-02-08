package logic

import (
	"context"

	"github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
	"github.com/BingguWang/bingBar/service/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowedListByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowedListByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowedListByUserIDLogic {
	return &GetFollowedListByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取关注列表
func (l *GetFollowedListByUserIDLogic) GetFollowedListByUserID(in *pb.GetFollowedListByUserIDReq) (*pb.GetFollowedListByUserIDResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFollowedListByUserIDResp{}, nil
}
