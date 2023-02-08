package logic

import (
	"context"

	"github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
	"github.com/BingguWang/bingBar/service/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListByUserIDLogic {
	return &GetFriendListByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户好友列表
func (l *GetFriendListByUserIDLogic) GetFriendListByUserID(in *pb.GetFriendListByUserIDReq) (*pb.GetFriendListByUserIDResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFriendListByUserIDResp{}, nil
}
