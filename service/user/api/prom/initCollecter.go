package prom

import "github.com/prometheus/client_golang/prometheus"

func InitCollector() {
    // 自定义GaugeVec 仪表盘类型指标
    userCountGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
        Namespace: "api_server",
        Name:      "fans_of_user_total",
    },
        []string{"user_id"}, // 设置标签后，在为指标赋值的时候也需要添加对应的 value
    )
    gauge := NewFansCounter(userCountGauge)
    prometheus.MustRegister(gauge.GaugeVec)
}
