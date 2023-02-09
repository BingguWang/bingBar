package user

import (
    "context"
    "github.com/BingguWang/bingBar/common/ctxdata"
    "github.com/BingguWang/bingBar/common/xerr"
    "github.com/BingguWang/bingBar/service/user/rpc/pb/pb"
    "github.com/jinzhu/copier"
    "github.com/pkg/errors"

    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetMutualFollowListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetMutualFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMutualFollowListLogic {
    return &GetMutualFollowListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

// GetMutualFollowList 从你的关注里找，从他的粉丝里找，然后求交集
func (l *GetMutualFollowListLogic) GetMutualFollowList(req *types.GetMutualFollowListReq) (*types.GetMutualFollowListResp, error) {
    uid := ctxdata.GetUidFromCtx(l.ctx)
    if uid == req.Followed {
        return nil, errors.Wrap(xerr.NewErrCode(xerr.ERROR_CODE_REUQEST_PARAM), "被关注者不能为当前用户")
    }
    ret, err := l.svcCtx.UserServiceRpc.GetMutualFollowed(l.ctx, &pb.GetMutualFollowedReq{
        UserId:   uid,
        Followed: req.Followed,
        PageSize: req.PageSize,
    })
    if err != nil {
        return nil, err
    }
    var resp types.GetMutualFollowListResp
    _ = copier.Copy(&resp.UserList, ret.GetUserList())
    resp.Total = ret.GetTotal()
    return &resp, nil
}
