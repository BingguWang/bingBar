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

type GetMutualFriendsLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetMutualFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMutualFriendsLogic {
    return &GetMutualFriendsLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

// 获取共同好友
func (l *GetMutualFriendsLogic) GetMutualFriends(in *pb.GetMutualFriendsReq) (*pb.GetMutualFriendsResp, error) {
    ids := in.GetUserIds()
    keys := make([]string, len(ids))
    for i := 0; i < len(ids); i++ {
        keys[i] = fmt.Sprintf("%s%v", UserFriendSetPrefix, ids[i])
    }
    mutualIds, err := l.svcCtx.Redis.SinterCtx(l.ctx, keys...)
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SSCAN_FAILED), "redis Sinter faied  , err : %s", err.Error())
    }
    fmt.Println(mutualIds)
    length := len(mutualIds)
    if length > int(in.PageSize) {
        length = int(in.PageSize)
    }
    resp := &pb.GetMutualFriendsResp{UserList: make([]*pb.User, length)}
    resp.Total = int64(len(mutualIds))
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
