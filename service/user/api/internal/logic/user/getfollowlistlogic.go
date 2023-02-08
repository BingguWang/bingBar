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

type GetFollowListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
    return &GetFollowListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetFollowListLogic) GetFollowList(req *types.GetFollowListReq) (*types.GetFollowListResp, error) {
    uid := ctxdata.GetUidFromCtx(l.ctx)
    ret, err := l.svcCtx.UserServiceRpc.GetFollowedListByUserID(l.ctx, &pb.GetFollowedListByUserIDReq{
        UserId:   uid,
        PageNo:   req.PageNo,
        PageSize: req.PageSize,
    })
    if err != nil {
        return nil, err
    }
    var resp types.GetFollowListResp
    _ = copier.Copy(&resp.UserList, ret.GetUserList())
    return &resp, nil
}
