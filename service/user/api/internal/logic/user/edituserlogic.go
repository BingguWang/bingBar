package user

import (
    "context"
    "github.com/BingguWang/bingBar/common/ctxdata"
    "github.com/BingguWang/bingBar/service/user/rpc/pb/pb"
    "github.com/jinzhu/copier"

    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type EditUserLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewEditUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserLogic {
    return &EditUserLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *EditUserLogic) EditUser(req *types.EditUserReq) (*types.EditUserResp, error) {
    u := pb.User{
        Id: ctxdata.GetUidFromCtx(l.ctx),
    }
    _ = copier.Copy(&u, req)
    ret, err := l.svcCtx.UserServiceRpc.EditUserInfo(l.ctx, &pb.EditUserInfoReq{
        User: &u,
    })
    if err != nil {
        return nil, err
    }
    resp := &types.EditUserResp{RetMsg: ret.GetRetMsg()}
    return resp, nil
}
