package logic

import (
    "context"
    "github.com/BingguWang/bingBar/common/xerr"

    "github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
    "github.com/BingguWang/bingBar/service/user/rpc/pb/pb"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

var ErrUserNoExistsError = xerr.NewErrCode(xerr.ERROR_CODE_USER_NOT_EXIST)

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
    return &GetUserInfoLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
    // todo: add your logic here and delete this line

    return &pb.GetUserInfoResp{}, nil
}
