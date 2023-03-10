package main

import (
    "flag"
    "fmt"
    "github.com/BingguWang/bingBar/service/user/api/internal/config"
    "github.com/BingguWang/bingBar/service/user/api/internal/handler"
    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/prom"
    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userservice.yaml", "the config file")

func init() {
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

    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
