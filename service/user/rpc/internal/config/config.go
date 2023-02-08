package config

import (
    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
    zrpc.RpcServerConf
    JwtAuth struct {
        AccessSecret string
        AccessExpire int64
    }
    DB struct {
        DataSource string
    }
    Cache cache.CacheConf // 缓存，数据库查询的时候会缓存到Redis里
    Redis struct {
        Password string
        Type     string
        Host     string
    }
}
