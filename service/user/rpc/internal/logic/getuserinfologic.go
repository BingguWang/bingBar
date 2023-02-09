package logic

import (
    "context"
    "github.com/BingguWang/bingBar/common/xerr"
    "github.com/BingguWang/bingBar/service/user/model"
    "github.com/jinzhu/copier"
    "github.com/pkg/errors"

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

// GetUserInfo 获取user的详细信息
func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
    one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
    if err != nil && err != model.ErrNotFound {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_DB_ERROR), "查询用户信息失败，id:%v,err:%v", in.Id, err)
    }
    if one == nil {
        return nil, errors.Wrapf(ErrUserNoExistsError, "id:%v", in.Id)
    }
    resp := pb.GetUserInfoResp{User: &pb.User{}}
    _ = copier.Copy(&resp.User, &one)
    return &resp, nil
}
