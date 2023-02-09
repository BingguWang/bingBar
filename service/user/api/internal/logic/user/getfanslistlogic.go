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

type GetFansListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFansListLogic {
    return &GetFansListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetFansListLogic) GetFansList(req *types.GetFansListReq) (*types.GetFansListResp, error) {
    uid := ctxdata.GetUidFromCtx(l.ctx)
    ret, err := l.svcCtx.UserServiceRpc.GetFansListByUserID(l.ctx, &pb.GetFansListByUserIDReq{
        UserId:   uid,
        PageSize: req.PageSize,
        PageNo:   req.PageNo,
    })
    if err != nil {
        return nil, err
    }
    resp := types.GetFansListResp{}
    _ = copier.Copy(&resp.UserList, ret.GetUserList())
    return &resp, nil
}
