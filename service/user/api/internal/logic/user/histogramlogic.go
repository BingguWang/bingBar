package user

import (
    "context"
    "fmt"
    "github.com/prometheus/client_golang/api"
    v1 "github.com/prometheus/client_golang/api/prometheus/v1"
    "github.com/prometheus/common/model"
    "time"

    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type HistogramLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewHistogramLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistogramLogic {
    return &HistogramLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *HistogramLogic) Histogram(req *types.HistogramFansReq) (resp *types.HistogramFansResp, err error) {
    client, err := api.NewClient(api.Config{
        Address: "http://" + l.svcCtx.Config.DevServer.Host + ":9090",
    })
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        return nil, err
    }
    ret, err := QueryFansHistogramOfUsers(client)
    if err != nil {
        return nil, err
    }
    return ret, nil
}

func QueryFansHistogramOfUsers(client api.Client) (*types.HistogramFansResp, error) {
    v1api := v1.NewAPI(client)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    result, warnings, err := v1api.Query(ctx, fmt.Sprintf(`api_server_fans_of_users_histogram_bucket{}`), time.Now())
    if err != nil {
        fmt.Printf("Error querying Prometheus: %v\n", err)
        return nil, err
    }
    if len(warnings) > 0 {
        fmt.Printf("Warnings: %v\n", warnings)
    }
    resp := &types.HistogramFansResp{
        HistogramFansItems: []types.HistogramFansItem{},
    }

    // 注意，返回的bucket的值是向下包含的!
    for _, sample := range result.(model.Vector) {
        set := model.LabelSet(sample.Metric)
        if labelValue, ok := set[model.BucketLabel]; ok {
            p := types.HistogramFansItem{
                Le:    string(labelValue),
                Value: float64(sample.Value),
            }
            resp.HistogramFansItems = append(resp.HistogramFansItems, p)
        }
    }
    return resp, nil
}
