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

type GetFriendListByUserIDLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetFriendListByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListByUserIDLogic {
    return &GetFriendListByUserIDLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

// 获取用户好友列表
func (l *GetFriendListByUserIDLogic) GetFriendListByUserID(in *pb.GetFriendListByUserIDReq) (*pb.GetFriendListByUserIDResp, error) {
    UserFriendSetPrefixKey := fmt.Sprintf("%s%v", UserFriendSetPrefix, in.UserId)
    keys, cur, err := l.svcCtx.Redis.SscanCtx(l.ctx, UserFriendSetPrefixKey, uint64(in.GetPageNo()-1), "", in.GetPageSize())
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_CODE_SSCAN_FAILED), "redis Sscan faied ,key : %v , err : %s", UserFriendSetPrefixKey, err.Error())
    }
    fmt.Println(cur)
    resp := &pb.GetFriendListByUserIDResp{UserList: make([]*pb.User, len(keys))}
    var wg sync.WaitGroup
    for i := 0; i < len(keys); i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            parseInt, _ := strconv.ParseInt(keys[i], 10, 0)
            resp.UserList[i] = &pb.User{
                Id: parseInt,
            }
        }(i)
    }
    wg.Wait()
    return resp, nil
}
