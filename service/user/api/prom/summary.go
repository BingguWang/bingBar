package prom

import (
    "github.com/prometheus/client_golang/prometheus"
    "sync"
)

type FansSummary struct {
    *prometheus.SummaryVec
}

var (
    fansSummaryonce     sync.Once
    fansSummaryInstance *FansSummary
)

func GetFansSummary() *FansCounter {
    return fansCounterInstance
}

func NewFansSummary(summaryVec *prometheus.SummaryVec) *FansSummary {
    fansSummaryonce.Do(func() {
        fansSummaryInstance = &FansSummary{summaryVec}
    })
    return fansSummaryInstance

}
