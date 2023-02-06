package user

import (
    "context"
    "github.com/BingguWang/bingBar/service/user/model"
    "github.com/BingguWang/bingBar/service/user/rpc/userservice"
    "github.com/jinzhu/copier"

    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
    return &LoginLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
    loginResp, err := l.svcCtx.UserServiceRpc.Login(l.ctx, &userservice.LoginReq{
        AuthType: model.UserAuthTypeSystem,
        AuthKey:  req.Mobile,
        Password: req.Password,
    })
    if err != nil {
        return nil, err
    }

    var resp types.LoginResp
    _ = copier.Copy(&resp, loginResp)

    return &resp, nil
}
