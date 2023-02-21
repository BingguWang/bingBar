package mointor

import "github.com/prometheus/client_golang/prometheus"

func InitCollector() {
    // 自定义counter类型指标
    requestCounter := prometheus.NewCounter(prometheus.CounterOpts{
        Namespace: "api_server",
        Name:      "request_total",
    })
    newRequestCounter := NewRequestCounter(requestCounter)

    // 自定义GaugeVec 仪表盘类型指标
    userCountGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
        Namespace: "api_server",
        Name:      "user_count",
    },
        []string{"user_type"}, // 设置标签后，在为指标赋值的时候也需要添加对应的 value
    )
    userCountGauge.WithLabelValues("administration").Add(3)
    userCountGauge.WithLabelValues("user").Add(300)
    gauge := NewUserCountGauge(userCountGauge)

    prometheus.MustRegister(newRequestCounter.RequestCounter, gauge.UserCountGauge)
}
