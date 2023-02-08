package logic

import (
    "context"
    "fmt"
    "github.com/BingguWang/bingBar/common/xerr"
    "github.com/pkg/errors"

    "github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
    "github.com/BingguWang/bingBar/service/user/rpc/pb/pb"

    "github.com/zeromicro/go-zero/core/logx"
)

type UnFollowLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFollowLogic {
    return &UnFollowLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

// UnFollow 取关对方
func (l *UnFollowLogic) UnFollow(in *pb.UnFollowReq) (*pb.UnFollowResp, error) {
    // 更新关注集合
    userFollowedSetPrefixkey := fmt.Sprintf("%s%v", UserFollowedSetPrefix, in.FollowBy)
    _, err := l.svcCtx.Redis.SremCtx(l.ctx, userFollowedSetPrefixkey, in.UserId)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SADD_FAILED), "redis srem faied ,key : %v , err : %s", userFollowedSetPrefixkey, err.Error())
    }
    // 关注粉丝集合
    userFansSetPrefixkey := fmt.Sprintf("%s%v", UserFansSetPrefix, in.UserId)
    _, err = l.svcCtx.Redis.SremCtx(l.ctx, userFansSetPrefixkey, in.FollowBy)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SADD_FAILED), "redis srem faied ,key : %v , err : %s", userFansSetPrefixkey, err.Error())
    }
    return &pb.UnFollowResp{
        RetMsg:  xerr.MapErrMsg(xerr.OK),
        RetCode: xerr.OK,
    }, nil
}
