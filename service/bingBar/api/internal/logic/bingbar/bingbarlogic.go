package bingbar

import (
    "context"
    tool "github.com/BingguWang/bingBar/common/utils"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/svc"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/types"
    "github.com/BingguWang/bingBar/service/testRpc/rpc/pb/pb"
    "github.com/zeromicro/go-zero/core/breaker"
    "github.com/zeromicro/go-zero/core/logx"
)

type BingbarLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewBingbarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BingbarLogic {
    return &BingbarLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *BingbarLogic) Bingbar(req *types.SimpleReq) (*types.SimpleResp, error) {
    logx.Infof("调用test服务")
    newBreaker := breaker.NewBreaker()
    if err := newBreaker.DoWithAcceptable(func() error {
        resp, err := l.svcCtx.TestService.TestRpcFunc(l.ctx, &pb.SimpleReq{})
        if err != nil {
            logx.Errorf("调用test 服务 失败 : ", err)
            return err
        }
        logx.Infof(tool.ToJsonString(resp))
        return nil
    }, func(err error) bool {
        return err == nil
    }); err != nil {
        return nil, err
    }

    //start := time.Now()
    //resp, err := l.svcCtx.TestService.TestRpcFunc(l.ctx, &pb.SimpleReq{})
    //if err != nil {
    //    logx.Errorf("调用test 服务 失败 : ", err)
    //    end := time.Now()
    //    fmt.Println(end.Sub(start))
    //    return nil, err
    //}
    //logx.Infof(tool.ToJsonString(resp))
    return nil, nil
}
