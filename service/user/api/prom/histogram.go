package prom

import (
    "github.com/prometheus/client_golang/prometheus"
    "sync"
)

type FansHistogram struct {
    *prometheus.HistogramVec
}

var (
    fansHistogramonce     sync.Once
    fansHistogramInstance *FansHistogram
)

func GetFansHistogram() *FansCounter {
    return fansCounterInstance
}

func NewFansHistogram(histogram *prometheus.HistogramVec) *FansHistogram {
    fansHistogramonce.Do(func() {
        fansHistogramInstance = &FansHistogram{histogram}
    })
    return fansHistogramInstance
}
