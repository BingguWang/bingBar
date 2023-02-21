// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "github.com/BingguWang/bingBar/service/user/api/internal/handler/user"
	"github.com/BingguWang/bingBar/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: user.UserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/editUser",
				Handler: user.EditUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/follow",
				Handler: user.FollowHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/unfollow",
				Handler: user.UnfollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/getFollowList",
				Handler: user.GetFollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/getFansList",
				Handler: user.GetFansListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/getFriendList",
				Handler: user.GetFriendListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/getMutualFollowList",
				Handler: user.GetMutualFollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/getMutualFriendList",
				Handler: user.GetMutualFriendListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/history/fans",
				Handler: user.HistoryfansHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/usercenter/v1"),
	)
}
