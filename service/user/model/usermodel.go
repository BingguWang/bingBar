package model

import (
    "context"
    "fmt"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlc"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
    // UserModel is an interface to be customized, add more methods here,
    // and implement the added methods in customUserModel.
    UserModel interface {
        userModel
        FindOneByQuery(ctx context.Context, query string, values []interface{}) (*User, error)
        FindOneByMobile(ctx context.Context, mobile string) (*User, error)
        // Trans 执行事务操作
        Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
    }

    customUserModel struct {
        *defaultUserModel
    }
)

func (c *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
    userserviceUserMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, mobile)
    var resp User
    // 把指定的查询结果将主键保存在缓存内
    err := c.QueryRowIndexCtx(ctx, &resp, userserviceUserMobileKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
        query := fmt.Sprintf("select %s from %s where `mobile` = ? and deleted_at is NULL limit 1", userRows, c.table)
        if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil { // 唯一索引查询方法
            return nil, err
        }
        return resp.Id, nil
    },
        c.queryPrimary, // 主键索引查询方法
    )
    switch err {
    case nil:
        return &resp, nil
    case sqlc.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
    return &customUserModel{
        defaultUserModel: newUserModel(conn, c),
    }
}

func (c *defaultUserModel) FindOneByQuery(ctx context.Context, query string, values []interface{}) (*User, error) {
    var resp User
    err := c.QueryRowNoCacheCtx(ctx, &resp, query, values...)
    switch err {
    case nil:
        return &resp, nil
    default:
        return nil, err
    }
}

func (c *defaultUserModel) Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error {
    // 传入的func就是要放入事务的逻辑
    logx.Info("----using transaction----")
    return c.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
        return fn(ctx, session)
    })
}
