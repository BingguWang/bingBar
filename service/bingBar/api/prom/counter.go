package prom

import (
    "github.com/prometheus/client_golang/prometheus"
    "sync"
)

type RequestCounter struct {
    RequestCounter prometheus.Counter
}

var (
    requestCounteronce     sync.Once
    requestCounterInstance *RequestCounter
)

func GetRequestCounter() *RequestCounter {
    return requestCounterInstance
}

func NewRequestCounter(requestCounter prometheus.Counter) *RequestCounter {
    requestCounteronce.Do(func() {
        requestCounterInstance = &RequestCounter{RequestCounter: requestCounter}
    })
    return requestCounterInstance

}
