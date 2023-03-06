package bingbar

import (
    "fmt"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/logic/bingbar"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/svc"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/types"
    "github.com/BingguWang/bingBar/service/bingBar/api/prom"
    "github.com/zeromicro/go-zero/rest/httpx"
    "net/http"
)

func BingbarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.SimpleReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }
        counter := prom.GetRequestCounter()
        // prom指标计数值加1
        counter.RequestCounter.Inc()

        //const (
        //    burst = 1000 // 令牌桶容量
        //    rate  = 200  // 令牌产生速率
        //)
        //store := svcCtx.Redis
        //fmt.Println(store.Ping())
        //// New tokenLimiter
        //limiter := limit.NewTokenLimiter(rate, burst, store, "rate-test")
        //if limiter.Allow() {
        //    atomic.AddInt32(&svcCtx.Allowed, 1)
        //} else {
        //    atomic.AddInt32(&svcCtx.Denied, 1)
        //    httpx.ErrorCtx(r.Context(), w, errors.New("被限流"))
        //    return
        //}
        fmt.Printf("allowed: %d, denied: %d ", svcCtx.Allowed, svcCtx.Denied)

        l := bingbar.NewBingbarLogic(r.Context(), svcCtx)
        resp, err := l.Bingbar(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
        } else {
            httpx.OkJsonCtx(r.Context(), w, resp)
        }
    }
}
