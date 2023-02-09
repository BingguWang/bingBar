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

type UserDetailLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
    return &UserDetailLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *UserDetailLogic) UserDetail(req *types.UserInfoReq) (*types.UserInfoResp, error) {
    uid := ctxdata.GetUidFromCtx(l.ctx)
    info, err := l.svcCtx.UserServiceRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{Id: uid})
    if err != nil {
        return nil, err
    }
    var resp types.UserInfoResp
    _ = copier.Copy(&resp.UserInfo, info.GetUser())
    return &resp, nil
}
