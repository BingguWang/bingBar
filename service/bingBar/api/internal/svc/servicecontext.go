package svc

import (
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/config"
    "github.com/BingguWang/bingBar/service/testRpc/rpc/testservice"
    "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
    Config      config.Config
    TestService testservice.TestService
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:      c,
        TestService: testservice.NewTestService(zrpc.MustNewClient(c.TestServiceRpcConf)),
    }
}
