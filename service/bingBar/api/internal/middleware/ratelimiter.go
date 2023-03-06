package middleware

import (
    "errors"
    "fmt"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/svc"
    "github.com/zeromicro/go-zero/core/limit"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/rest/httpx"
    "net/http"
    "sync/atomic"
)

func MyLimiterHandler(ctx *svc.ServiceContext) func(next http.HandlerFunc) http.HandlerFunc {
    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            logx.Info("global middleware")

            const (
                burst = 1000 // 令牌桶容量
                rate  = 200  // 令牌产生速率
            )
            store := ctx.Redis
            fmt.Println(store.Ping())
            // New tokenLimiter
            limiter := limit.NewTokenLimiter(rate, burst, store, "rate-test")
            if limiter.Allow() {
                atomic.AddInt32(&ctx.Allowed, 1)
            } else {
                atomic.AddInt32(&ctx.Denied, 1)
                httpx.ErrorCtx(r.Context(), w, errors.New("被限流"))
                return
            }
            next(w, r)
        }
    }
}
