package logic

import (
	"context"

	"github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
	"github.com/BingguWang/bingBar/service/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFansListByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFansListByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFansListByUserIDLogic {
	return &GetFansListByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取粉丝列表
func (l *GetFansListByUserIDLogic) GetFansListByUserID(in *pb.GetFansListByUserIDReq) (*pb.GetFansListByUserIDResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFansListByUserIDResp{}, nil
}
