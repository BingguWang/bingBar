package logic

import (
    "context"
    "fmt"
    tool "github.com/BingguWang/bingBar/common/utils"
    "github.com/BingguWang/bingBar/common/xerr"
    "github.com/BingguWang/bingBar/service/user/model"
    "github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
    "github.com/BingguWang/bingBar/service/user/rpc/pb/pb"
    "github.com/BingguWang/bingBar/service/user/rpc/userservice"
    "github.com/pkg/errors"
    "github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrCode(xerr.ERROR_CODE_TOKEN_GENERATE)
var ErrUsernamePwdError = xerr.NewErrCode(xerr.ERROR_CODE_USER_PASSWORD_ERROR)

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
    return &LoginLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
    var userId int64
    var err error
    switch in.AuthType {
    case model.UserAuthTypeSystem:
        userId, err = l.loginByMobile(in.AuthKey, in.Password)
    default:
        return nil, xerr.NewErrCode(xerr.ERROR_CODE_REUQEST_PARAM)
    }
    if err != nil {
        return nil, err
    }

    // 生成token
    generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
    tokenResp, err := generateTokenLogic.GenerateToken(&userservice.GenerateTokenReq{
        UserId: userId,
    })
    if err != nil {
        return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
    }

    return &userservice.LoginResp{
        AccessToken:  tokenResp.AccessToken,
        AccessExpire: tokenResp.AccessExpire,
        RefreshAfter: tokenResp.RefreshAfter,
    }, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {
    //   查找用户
    one, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
    if err != nil && err != model.ErrNotFound {
        return 0, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
    }
    if one == nil {
        return 0, errors.Wrapf(ErrUserNoExistsError, "mobile:%s", mobile)
    }

    fmt.Println("查询结果:", tool.ToJsonString(one))

    //   解码比较
    if !(tool.Md5ByString(password) == one.Password) {
       return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
    }

    return one.Id, nil
}
