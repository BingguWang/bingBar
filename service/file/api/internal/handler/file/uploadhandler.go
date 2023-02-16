package file

import (
    "fmt"
    tool "github.com/BingguWang/bingBar/common/utils"
    fi "github.com/BingguWang/bingBar/service/file/api/internal/logic/file"
    "github.com/BingguWang/bingBar/service/file/api/internal/types"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/rest/httpx"
    "net/http"

    "github.com/BingguWang/bingBar/service/file/api/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.UploadFileRequest
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }
        // 文件接收
        if err2 := r.ParseMultipartForm(1 << 20); err2 != nil {
            httpx.ErrorCtx(r.Context(), w, err2)
        }
        logx.Info(r.MultipartForm)
        file, header, err := r.FormFile("file")
        if err != nil {
            fmt.Println(err)
            httpx.ErrorCtx(r.Context(), w, err)
        }
        logx.Errorf(tool.ToJsonString(header))
        defer file.Close()
        logx.Infof("成功获取文件")

        fileName := r.FormValue("filename")

        l := fi.NewUploadLogic(r.Context(), svcCtx)
        resp, err := l.Upload(file, fileName, &req)
        //resp, err := l.Upload(nil, "fileName", &req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
        } else {
            httpx.OkJsonCtx(r.Context(), w, resp)
        }
    }
}
