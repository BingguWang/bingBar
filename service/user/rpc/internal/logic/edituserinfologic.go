package logic

import (
    "context"
    "fmt"
    tool "github.com/BingguWang/bingBar/common/utils"
    "github.com/BingguWang/bingBar/common/xerr"
    "github.com/BingguWang/bingBar/service/user/model"
    "github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
    "github.com/BingguWang/bingBar/service/user/rpc/pb/pb"
    "github.com/jinzhu/copier"
    "github.com/pkg/errors"
    "github.com/zeromicro/go-zero/core/stores/sqlx"

    "github.com/zeromicro/go-zero/core/logx"
)

type EditUserInfoLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewEditUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserInfoLogic {
    return &EditUserInfoLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

// 修改用户信息
func (l *EditUserInfoLogic) EditUserInfo(in *pb.EditUserInfoReq) (*pb.EditUserInfoResp, error) {
    u := model.User{
        Id: in.User.Id,
    }
    _ = copier.Copy(&u, &in.User)
    fmt.Println(tool.ToJsonString(u))
    if err := l.svcCtx.UserModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
        if err := l.svcCtx.UserModel.Update(l.ctx, &u, session); err != nil {
            return errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_DB_ERROR), "update user err  :%v,user:%+v", err, u.Id)
        }
        return nil
    }); err != nil {
        return nil, err
    }
    return &pb.EditUserInfoResp{
        RetMsg:  xerr.MapErrMsg(xerr.OK),
        RetCode: xerr.OK,
    }, nil
}
