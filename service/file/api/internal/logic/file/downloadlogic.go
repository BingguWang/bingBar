package file

import (
	"context"

	"github.com/BingguWang/bingBar/service/file/api/internal/svc"
	"github.com/BingguWang/bingBar/service/file/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadFileRequest) (resp *types.DownloadFileResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
