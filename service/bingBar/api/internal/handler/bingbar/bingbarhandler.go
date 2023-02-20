package bingbar

import (
	"net/http"

	"github.com/BingguWang/bingBar/service/bingBar/api/internal/logic/bingbar"
	"github.com/BingguWang/bingBar/service/bingBar/api/internal/svc"
	"github.com/BingguWang/bingBar/service/bingBar/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BingbarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SimpleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bingbar.NewBingbarLogic(r.Context(), svcCtx)
		resp, err := l.Bingbar(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}