package svc

import (
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/config"
    "github.com/BingguWang/bingBar/service/testRpc/rpc/testservice"
    "github.com/zeromicro/go-zero/core/stores/redis"
    "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
    Config      config.Config
    TestService testservice.TestService
    Redis       *redis.Redis
    Allowed     int32
    Denied      int32
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:      c,
        TestService: testservice.NewTestService(zrpc.MustNewClient(c.TestServiceRpcConf)),
        Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
            r.Type = c.Redis.Type
            r.Pass = c.Redis.Password
        }),
    }
}
