package user

import (
    "context"
    "fmt"
    tool "github.com/BingguWang/bingBar/common/utils"
    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/internal/types"
    "github.com/prometheus/client_golang/api"
    v1 "github.com/prometheus/client_golang/api/prometheus/v1"
    "github.com/prometheus/common/model"
    "time"

    "github.com/zeromicro/go-zero/core/logx"
)

type SummaryLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SummaryLogic {
    return &SummaryLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *SummaryLogic) Summary(req *types.SummaryFansReq) (*types.SummaryFansResp, error) {
    client, err := api.NewClient(api.Config{
        Address: "http://" + l.svcCtx.Config.DevServer.Host + ":9090",
    })
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        return nil, err
    }
    ret, err := QueryFansSummaryOfUsers(client)
    if err != nil {
        return nil, err
    }
    return ret, nil
}

func QueryFansSummaryOfUsers(client api.Client) (*types.SummaryFansResp, error) {
    v1api := v1.NewAPI(client)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    result, warnings, err := v1api.Query(ctx, fmt.Sprintf(`fans_of_users_summary{name="fans"}`), time.Now())
    if err != nil {
        fmt.Printf("Error querying Prometheus: %v\n", err)
        return nil, err
    }
    if len(warnings) > 0 {
        fmt.Printf("Warnings: %v\n", warnings)
    }
    resp := &types.SummaryFansResp{
        SummaryItems: []types.SummaryFansItem{},
    }
    for _, sample := range result.(model.Vector) {
        set := model.LabelSet(sample.Metric)
        if labelValue, ok := set[model.QuantileLabel]; ok {
            p := types.SummaryFansItem{
                Percentile: string(labelValue),
                Value:      float64(sample.Value),
            }
            resp.SummaryItems = append(resp.SummaryItems, p)
        }
    }
    fmt.Println("-------",tool.ToJsonString(resp))
    return resp, nil
}
