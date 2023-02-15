package file

import (
	"net/http"

	"github.com/BingguWang/bingBar/service/file/api/internal/logic/file"
	"github.com/BingguWang/bingBar/service/file/api/internal/svc"
	"github.com/BingguWang/bingBar/service/file/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewDownloadLogic(r.Context(), svcCtx)
		resp, err := l.Download(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
