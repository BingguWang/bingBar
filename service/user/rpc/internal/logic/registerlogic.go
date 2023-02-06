package logic

import (
	"context"
	"github.com/BingguWang/bingBar/service/user/rpc/internal/svc"
	"github.com/BingguWang/bingBar/service/user/rpc/pb/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// todo 查找用户是否存在
	//if user != nil {
	//	return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", in.Mobile, err)
	//}

	// todo 不存在，通过事务插入用户

	// todo 生成token

	return &pb.RegisterResp{}, nil
}
