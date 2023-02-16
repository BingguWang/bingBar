package file

import (
    "bufio"
    "context"
    "fmt"
    "io"
    "mime/multipart"
    "os"

    "github.com/BingguWang/bingBar/service/file/api/internal/svc"
    "github.com/BingguWang/bingBar/service/file/api/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
    return &UploadLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *UploadLogic) Upload(file multipart.File, filename string, req *types.UploadFileRequest) (resp *types.UploadFileResponse, err error) {
    path := "/home/tmp/" + filename
    if _, e := os.Stat(path); e != nil {
        if os.IsNotExist(e) {
            _, err := os.Create(path)
            if err != nil {
                return nil, err
            }
        } else {
            return nil, e
        }
    }
    fmt.Println("文件已存在")
    f, err := os.OpenFile(path, 777, os.ModePerm)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    reader := bufio.NewReader(file)
    buf := make([]byte, 1<<10)
    file.Seek(0, 0) // 从文件流最开始读
    var (
        writeByteCount int // 写入文件的字节数
        offset         int
    )
    for {
        // 读出
        n, err := reader.Read(buf) // 如果buf长度大于reader设置的缓存大小，就认为是大文件上传,会避免copy,直接从reader里read到buf,否则是copy到buf里
        //fmt.Println("读出:", n)
        if err == io.EOF {
            fmt.Println("上传完成")
            break
        }
        // 写入
        wn, e := f.WriteAt(buf[:n], int64(offset))
        if e != nil {
            fmt.Println(e.Error())
            return nil, e
        }
        writeByteCount += wn
        offset += wn
    }
    logx.Infof("文件写入完毕!")
    return
}
