package prom

import (
    "github.com/prometheus/client_golang/prometheus"
)

func InitCollector() {
    // 自定义GaugeVec 仪表盘类型指标
    userCountGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
        Namespace: "api_server",
        Name:      "fans_of_user_total",
    },
        []string{"user_id"}, // 设置标签后，在为指标赋值的时候也需要添加对应的 value
    )
    gauge := NewFansCounter(userCountGauge)

    // summary
    summary := prometheus.NewSummaryVec(
        prometheus.SummaryOpts{
            Name:       "fans_of_users_summary",
            Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
        },
        []string{"name"},
    )
    su := NewFansSummary(summary)

    // 统计一下粉丝数在2-10-20的用户
    histogram := prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Namespace: "api_server",
            Name:      "fans_of_users_histogram",
            Buckets:   []float64{10, 25, 50, 75, 100}, // 0-10 10-25 ...
        },
        []string{"count"},
    )
    hi := NewFansHistogram(histogram)

    // 传入的i假设就是我们每个用户的粉丝数
    //for i := 1; i <= 100; i++ {
    //    //a := (time.Now().Second() + rand.Int()) % 100
    //    a := i
    //    summary.WithLabelValues("fans").Observe(float64(a))
    //    histogram.WithLabelValues("fans_histogram").Observe(float64(a))
    //}

    prometheus.MustRegister(gauge.GaugeVec, su.SummaryVec, hi.HistogramVec)

}
