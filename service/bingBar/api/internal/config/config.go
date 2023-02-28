package config

import (
    "github.com/zeromicro/go-zero/rest"
    "github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
    rest.RestConf
    TestServiceRpcConf zrpc.RpcClientConf
    Redis              struct {
        Password string
        Type     string
        Host     string
    }
}
