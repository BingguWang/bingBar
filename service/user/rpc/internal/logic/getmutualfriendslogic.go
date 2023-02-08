package logic

import (
	"context"

	"github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
	"github.com/BingguWang/bingBar/service/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMutualFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMutualFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMutualFriendsLogic {
	return &GetMutualFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取共同好友
func (l *GetMutualFriendsLogic) GetMutualFriends(in *pb.GetMutualFriendsReq) (*pb.GetMutualFriendsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMutualFriendsResp{}, nil
}
