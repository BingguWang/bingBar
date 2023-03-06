package main

import (
    "flag"
    "fmt"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/config"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/handler"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/middleware"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/svc"
    "github.com/BingguWang/bingBar/service/bingBar/api/prom"
    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/bingbarservice.yaml", "the config file")

func init() {
    // 自定义监控指标
    /*jobsInQueue := prometheus.NewGaugeVec(prometheus.GaugeOpts{
          Namespace: "api_server",
          Name:      "call",
          Help:      "Current number of jobs in the queue",
      }, []string{"job_type"})
      jobsInQueue.WithLabelValues("myJob-1").Add(3)

      jobsInQueueRpc := prometheus.NewGaugeVec(prometheus.GaugeOpts{
          Namespace: "rpc_server",
          Name:      "job_in_queue",
          Help:      "Current number of jobs in the queue",
      }, []string{"job_type"})
      jobsInQueueRpc.WithLabelValues("myJob-2").Add(10)*/

    prom.InitCollector()
}
func main() {
    flag.Parse()

    var c config.Config
    conf.MustLoad(*configFile, &c)

    server := rest.MustNewServer(c.RestConf)
    defer server.Stop()

    ctx := svc.NewServiceContext(c)
    handler.RegisterHandlers(server, ctx)
    // 全局中间件
    server.Use(middleware.MyLimiterHandler(ctx))

    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
