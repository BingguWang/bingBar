package user

import (
    "context"
    "github.com/BingguWang/bingBar/common/ctxdata"
    "github.com/BingguWang/bingBar/common/xerr"
    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/internal/types"
    "github.com/BingguWang/bingBar/service/user/model"
    "github.com/BingguWang/bingBar/service/user/rpc/userservice"
    "github.com/jinzhu/copier"

    "github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
    return &FollowLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

// Follow 关注
func (l *FollowLogic) Follow(req *types.FollowReq) (*types.FollowResp, error) {
    uid := ctxdata.GetUidFromCtx(l.ctx)
    logx.Infof("uid : %v", uid)
    if uid == req.Uid {
        return nil, xerr.NewErrMsg("用户不能关注自己")
    }
    followResp, err := l.svcCtx.UserServiceRpc.Follow(l.ctx, &userservice.FollowReq{
        UserId:   req.Uid,
        FollowBy: uid,
        AuthType: model.UserAuthTypeSystem,
    })
    if err != nil {
        return nil, err
    }
    var resp types.FollowResp
    _ = copier.Copy(&resp, followResp)
    return &resp, nil
}
