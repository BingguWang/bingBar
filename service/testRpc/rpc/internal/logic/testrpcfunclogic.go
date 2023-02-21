package logic

import (
    "context"
    "github.com/pkg/errors"
    "math/rand"

    "github.com/BingguWang/bingBar/service/testRpc/rpc/internal/svc"
    "github.com/BingguWang/bingBar/service/testRpc/rpc/pb/pb"

    "github.com/zeromicro/go-zero/core/logx"
)

type TestRpcFuncLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewTestRpcFuncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestRpcFuncLogic {
    return &TestRpcFuncLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *TestRpcFuncLogic) TestRpcFunc(in *pb.SimpleReq) (*pb.SimpleResp, error) {
    //return &pb.SimpleResp{}, nil
    // 随机设置错误
    if rand.Int()%7 == 0 {
        logx.Infof("调用失败") // 看这个输出的次数就可以知道被熔断器熔断了的请求有多少
        return nil, errors.New("call TestRpcFunc failed")
    }
    logx.Infof("调用成功")
    // 可以看到输出 “调用成功”和“调用失败”的次数总数< rpc请求总数， 可见有些请求被熔断了，直接再客户端就返回了
    return &pb.SimpleResp{}, nil
}
