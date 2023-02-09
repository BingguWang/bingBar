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

type GetMutualFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMutualFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMutualFriendListLogic {
	return &GetMutualFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMutualFriendListLogic) GetMutualFriendList(req *types.GetMutualFriendListReq) (  *types.GetMutualFriendListResp,   error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	if uid == req.Uid {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.ERROR_CODE_REUQEST_PARAM), "不能为当前用户")
	}
	ret, err := l.svcCtx.UserServiceRpc.GetMutualFriends(l.ctx, &pb.GetMutualFriendsReq{
		UserIds:   []int64{req.Uid, uid},
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var resp types.GetMutualFriendListResp
	_ = copier.Copy(&resp.UserList, ret.GetUserList())
	resp.Total = ret.GetTotal()
	return &resp, nil
}
