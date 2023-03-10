// Code generated by goctl. DO NOT EDIT.
package handler

import (
    "net/http"

    bingbar "github.com/BingguWang/bingBar/service/bingBar/api/internal/handler/bingbar"
    "github.com/BingguWang/bingBar/service/bingBar/api/internal/svc"

    "github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
    server.AddRoutes(
        []rest.Route{
            {
                Method:  http.MethodPost,
                Path:    "/bingbar/a",
                Handler: bingbar.BingbarHandler(serverCtx),
            },
        },
        rest.WithPrefix("/bingbar/v1"),
    )

}

