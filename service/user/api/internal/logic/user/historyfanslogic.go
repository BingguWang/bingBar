package user

import (
    "context"
    "fmt"
    "github.com/BingguWang/bingBar/common/ctxdata"
    "github.com/prometheus/client_golang/api"
    v1 "github.com/prometheus/client_golang/api/prometheus/v1"
    "github.com/prometheus/common/model"
    "time"

    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type HistoryfansLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewHistoryfansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryfansLogic {
    return &HistoryfansLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *HistoryfansLogic) Historyfans(req *types.FansHistoryReq) (*types.FansHistoryResp, error) {
    client, err := api.NewClient(api.Config{
        Address: "http://" + l.svcCtx.Config.DevServer.Host + ":9090",
    })
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        return nil, err
    }
    uid := ctxdata.GetUidFromCtx(l.ctx)
    ret, err := QueryFansHistoryOfUser(client, uid, req.Start, req.End)
    if err != nil {
        return nil, err
    }
    return ret, nil
}

func QueryFansHistoryOfUser(client api.Client, uid int64, start, end int64) (*types.FansHistoryResp, error) {
    v1api := v1.NewAPI(client)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    r := v1.Range{
        Start: time.Unix(start, 0),
        End:   time.Unix(end, 0),
        Step:  time.Minute,
    }
    result, warnings, err := v1api.QueryRange(ctx, fmt.Sprintf(`api_server_fans_of_user_total{user_id="%v"}`, uid), r, v1.WithTimeout(5*time.Second))
    if err != nil {
        fmt.Printf("Error querying Prometheus: %v\n", err)
        return nil, err
    }
    if len(warnings) > 0 {
        fmt.Printf("Warnings: %v\n", warnings)
    }

    resp := &types.FansHistoryResp{
        HistoryItems: []types.HistoryItem{},
    }
    for _, sample := range result.(model.Matrix) {
        for _, value := range sample.Values {
            historyItem := types.HistoryItem{
                Value:  int64(value.Value),
                TimeAt: int64(value.Timestamp),
            }
            resp.HistoryItems = append(resp.HistoryItems, historyItem)
        }
    }
    return resp, nil
}
