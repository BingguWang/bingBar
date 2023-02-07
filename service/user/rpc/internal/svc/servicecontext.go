package svc

import (
    "github.com/BingguWang/bingBar/service/user/model"
    "github.com/BingguWang/bingBar/service/user/rpc/internal/config"
    "github.com/zeromicro/go-zero/core/stores/redis"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

// ServiceContext 装配依赖
type ServiceContext struct {
    Config config.Config

    UserModel     model.UserModel
    UserAuthModel model.UserAuthModel
    Redis         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
    sqlConn := sqlx.NewMysql(c.DB.DataSource)
    return &ServiceContext{
        Config: c,

        Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
            r.Type = c.Redis.Type
            r.Pass = c.Redis.Pass
        }),

        UserModel:     model.NewUserModel(sqlConn, c.Cache),
        UserAuthModel: model.NewUserAuthModel(sqlConn, c.Cache),
    }
}
