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

type GetFriendListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
    return &GetFriendListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetFriendListLogic) GetFriendList(req *types.GetFriendListReq) (*types.GetFriendListResp, error) {
    uid := ctxdata.GetUidFromCtx(l.ctx)
    ret, err := l.svcCtx.UserServiceRpc.GetFriendListByUserID(l.ctx, &pb.GetFriendListByUserIDReq{
        UserId:   uid,
        PageNo:   req.PageNo,
        PageSize: req.PageSize,
    })
    if err != nil {
        return nil, err
    }
    var resp types.GetFriendListResp
    _ = copier.Copy(&resp.UserList, ret.GetUserList())
    return &resp, nil
}
