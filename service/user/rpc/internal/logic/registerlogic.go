package logic

import (
    "context"
    tool "github.com/BingguWang/bingBar/common/utils"
    "github.com/BingguWang/bingBar/common/xerr"
    "github.com/BingguWang/bingBar/service/user/model"
    "github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
    "github.com/BingguWang/bingBar/service/user/rpc/pb/pb"
    "github.com/BingguWang/bingBar/service/user/rpc/userservice"
    "github.com/pkg/errors"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
    return &RegisterLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
    user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
    if err != nil && err != model.ErrNotFound {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_DB_ERROR), "mobile:%s,err:%v", in.Mobile, err)
    }
    if user != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_USER_ALREADY_EXIST), "Register user exists mobile:%s,err:%v", in.Mobile, err)
    }

    // 不存在，通过事务新增用户
    var userId int64
    if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
        user := new(model.User)
        user.Mobile = in.Mobile
        if len(user.NickName) == 0 {
            user.NickName = tool.Krand(8, tool.KC_RAND_KIND_ALL)
        }
        if len(in.Password) > 0 {
            user.Password = tool.Md5ByString(in.Password)
        }
        // insert
        insertResult, err := l.svcCtx.UserModel.Insert(ctx, user, session)
        if err != nil {
            logx.Errorf("insert into user err: %v", err.Error())
            return errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
        }
        lastId, err := insertResult.LastInsertId()
        if err != nil {
            logx.Errorf("insert into user insertResult.LastInsertId  err: %v", err.Error())
            return errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_DB_ERROR), "Register db user insertResult.LastInsertId err:%v,user:%+v", err, user)
        }
        userId = lastId
        userAuth := &model.UserAuth{
            AuthKey:  in.AuthKey,
            AuthType: in.AuthType,
            UserId:   lastId,
        }
        // insert into user auth
        if _, err := l.svcCtx.UserAuthModel.Insert(ctx, userAuth, session); err != nil {
            logx.Errorf("insert into user_auth err: %v", err.Error())
            return errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_DB_ERROR), "Register db user_auth Insert err:%v,userAuth:%v", err, userAuth)
        }
        return nil
    }); err != nil {
        return nil, err
    }

    //  生成token
    generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
    tokenResp, err := generateTokenLogic.GenerateToken(&userservice.GenerateTokenReq{
        UserId: userId,
    })
    if err != nil {
        logx.Errorf("Generate Token err: %v", err.Error())
        return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
    }

    return &pb.RegisterResp{
        AccessToken:  tokenResp.AccessToken,
        AccessExpire: tokenResp.AccessExpire,
        RefreshAfter: tokenResp.RefreshAfter,
    }, nil
}
