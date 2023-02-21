package mointor

import (
    "github.com/prometheus/client_golang/prometheus"
    "sync"
)

type UserCountGauge struct {
    UserCountGauge *prometheus.GaugeVec
}

var (
    userCountGaugeonce     sync.Once
    userCountGaugeInstance *UserCountGauge
)

func GetUserCountGauge() *UserCountGauge {
    return userCountGaugeInstance
}

func NewUserCountGauge(g *prometheus.GaugeVec) *UserCountGauge {
    userCountGaugeonce.Do(func() {
        userCountGaugeInstance = &UserCountGauge{UserCountGauge: g}
    })
    return userCountGaugeInstance

}
