package logic

import (
    "context"
    "fmt"
    "github.com/BingguWang/bingBar/common/xerr"
    "github.com/pkg/errors"
    "strconv"
    "sync"

    "github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
    "github.com/BingguWang/bingBar/service/user/rpc/pb/pb"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetMutualFollowedLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetMutualFollowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMutualFollowedLogic {
    return &GetMutualFollowedLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

// 获取共同关注
func (l *GetMutualFollowedLogic) GetMutualFollowed(in *pb.GetMutualFollowedReq) (*pb.GetMutualFollowedResp, error) {
    // 从你的关注里找，从他的粉丝里找，然后求交集
    UserFollowedSetPrefixKey := fmt.Sprintf("%s%v", UserFollowedSetPrefix, in.UserId)
    UserFansSetPrefixKey := fmt.Sprintf("%s%v", UserFansSetPrefix, in.Followed)
    mutualIds, err := l.svcCtx.Redis.SinterCtx(l.ctx, UserFollowedSetPrefixKey, UserFansSetPrefixKey)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SSCAN_FAILED), "redis Sinter faied  , err : %s", err.Error())
    }
    fmt.Println(mutualIds)
    length := len(mutualIds)
    if length > int(in.PageSize) {
        length = int(in.PageSize)
    }
    resp := &pb.GetMutualFollowedResp{UserList: make([]*pb.User, length)}
    var wg sync.WaitGroup
    for i := 0; i < len(mutualIds); i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            parseInt, _ := strconv.ParseInt(mutualIds[i], 10, 0)
            resp.UserList[i] = &pb.User{
                Id: parseInt,
            }
        }(i)
    }
    wg.Wait()
    return resp, nil
}
