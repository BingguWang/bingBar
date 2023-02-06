package svc

import (
    "github.com/BingguWang/bingBar/service/user/api/internal/config"
    "github.com/BingguWang/bingBar/service/user/rpc/userservice"
    "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
    Config         config.Config
    UserServiceRpc userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:         c,
        UserServiceRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserServiceRpcConf)),
    }
}
