package user

import (
    "context"
    "github.com/BingguWang/bingBar/common/ctxdata"
    "github.com/BingguWang/bingBar/service/user/model"
    "github.com/BingguWang/bingBar/service/user/rpc/userservice"
    "github.com/jinzhu/copier"

    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
    return &UnfollowLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *UnfollowLogic) Unfollow(req *types.UnFollowReq) (*types.UnFollowResp, error) {
    uid := ctxdata.GetUidFromCtx(l.ctx)
    logx.Infof("uid : %v", uid)
    followResp, err := l.svcCtx.UserServiceRpc.UnFollow(l.ctx, &userservice.UnFollowReq{
        UserId:   req.Uid,
        FollowBy: uid,
        AuthType: model.UserAuthTypeSystem,
    })
    if err != nil {
        return nil, err
    }
    var resp types.UnFollowResp
    _ = copier.Copy(&resp, followResp)
    return &resp, nil
}
