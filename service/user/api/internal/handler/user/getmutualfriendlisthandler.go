package user

import (
    "net/http"

    "github.com/BingguWang/bingBar/service/user/api/internal/logic/user"
    "github.com/BingguWang/bingBar/service/user/api/internal/svc"
    "github.com/BingguWang/bingBar/service/user/api/internal/types"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func GetMutualFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.GetMutualFriendListReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        l := user.NewGetMutualFriendListLogic(r.Context(), svcCtx)
        resp, err := l.GetMutualFriendList(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
        } else {
            httpx.OkJsonCtx(r.Context(), w, resp)
        }
    }
}
