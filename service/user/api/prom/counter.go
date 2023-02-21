package prom

import (
    "github.com/prometheus/client_golang/prometheus"
    "sync"
)

type FansCounter struct {
    *prometheus.GaugeVec
}

var (
    fansCounteronce     sync.Once
    fansCounterInstance *FansCounter
)

func GetFansCounter() *FansCounter {
    return fansCounterInstance
}

func NewFansCounter(gaugeVec *prometheus.GaugeVec) *FansCounter {
    fansCounteronce.Do(func() {
        fansCounterInstance = &FansCounter{gaugeVec}
    })
    return fansCounterInstance

}
