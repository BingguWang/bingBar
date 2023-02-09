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

type FollowLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
    return &FollowLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

// ============================ 用户关系 =======================

// Follow 关注
func (l *FollowLogic) Follow(in *pb.FollowReq) (*pb.FollowResp, error) {
    // 检查被关注的用户是否存在
    _, err2 := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
    if err2 != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_DB_ERROR), "find the followed user faied ,key : %v , err : %s", in.UserId, err2.Error())
    }

    // 更新关注集合
    UserFollowedSetPrefixKey := fmt.Sprintf("%s%v", UserFollowedSetPrefix, in.FollowBy)
    _, err := l.svcCtx.Redis.SaddCtx(l.ctx, UserFollowedSetPrefixKey, in.UserId)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SADD_FAILED), "redis sadd faied ,key : %v , err : %s", UserFollowedSetPrefixKey, err.Error())
    }
    // 更新粉丝集合
    UserFansSetPrefixKey := fmt.Sprintf("%s%v", UserFansSetPrefix, in.UserId)
    _, err = l.svcCtx.Redis.SaddCtx(l.ctx, UserFansSetPrefixKey, in.FollowBy)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SADD_FAILED), "redis sadd faied ,key : %v , err : %s", UserFansSetPrefixKey, err.Error())
    }
    // 更新好友集合
    Key1 := fmt.Sprintf("%s%v", UserFansSetPrefix, in.FollowBy)
    r1, err := l.svcCtx.Redis.SismemberCtx(l.ctx, Key1, in.UserId)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SISMEMBER_FAILED), "redis sismember faied ,key : %v , err : %s", Key1, err.Error())
    }
    Key2 := fmt.Sprintf("%s%v", UserFansSetPrefix, in.UserId)
    r2, err := l.svcCtx.Redis.SismemberCtx(l.ctx, Key2, in.FollowBy)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SISMEMBER_FAILED), "redis sismember faied ,key : %v , err : %s", Key2, err.Error())
    }
    Key3 := fmt.Sprintf("%s%v", UserFollowedSetPrefix, in.FollowBy)
    r3, err := l.svcCtx.Redis.SismemberCtx(l.ctx, Key3, in.UserId)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SISMEMBER_FAILED), "redis sismember faied ,key : %v , err : %s", Key3, err.Error())
    }
    Key4 := fmt.Sprintf("%s%v", UserFollowedSetPrefix, in.UserId)
    r4, err := l.svcCtx.Redis.SismemberCtx(l.ctx, Key4, in.FollowBy)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SISMEMBER_FAILED), "redis sismember faied ,key : %v , err : %s", Key4, err.Error())
    }
    // 是好友关系
    if r1 && r2 && r3 && r4 {
        logx.Infof("他们是好友关系:%v :%v", in.FollowBy, in.UserId)
        UserFriendSetPrefixKey1 := fmt.Sprintf("%s%v", UserFriendSetPrefix, in.FollowBy)
        _, err = l.svcCtx.Redis.SaddCtx(l.ctx, UserFriendSetPrefixKey1, in.FollowBy)
        if err != nil {
            return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SADD_FAILED), "redis sadd faied ,key : %v , err : %s", UserFriendSetPrefixKey1, err.Error())
        }
        UserFriendSetPrefixKey2 := fmt.Sprintf("%s%v", UserFriendSetPrefix, in.UserId)
        _, err = l.svcCtx.Redis.SaddCtx(l.ctx, UserFriendSetPrefixKey2, in.FollowBy)
        if err != nil {
            return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SADD_FAILED), "redis sadd faied ,key : %v , err : %s", UserFriendSetPrefixKey2, err.Error())
        }
    }

    return &pb.FollowResp{
        RetMsg:  xerr.MapErrMsg(xerr.OK),
        RetCode: xerr.OK,
    }, nil
}
