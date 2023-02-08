package logic

import (
	"context"

	"github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
	"github.com/BingguWang/bingBar/service/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMutualFollowedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMutualFollowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMutualFollowedLogic {
	return &GetMutualFollowedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取共同关注
func (l *GetMutualFollowedLogic) GetMutualFollowed(in *pb.GetMutualFollowedReq) (*pb.GetMutualFollowedResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMutualFollowedResp{}, nil
}
