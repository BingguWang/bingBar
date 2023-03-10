// Code generated by goctl. DO NOT EDIT.
// Source: test.proto

package server

import (
	"context"

	"github.com/BingguWang/bingBar/service/testRpc/rpc/internal/logic"
	"github.com/BingguWang/bingBar/service/testRpc/rpc/internal/svc"
	"github.com/BingguWang/bingBar/service/testRpc/rpc/pb/pb"
)

type TestServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTestServiceServer
}

func NewTestServiceServer(svcCtx *svc.ServiceContext) *TestServiceServer {
	return &TestServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *TestServiceServer) TestRpcFunc(ctx context.Context, in *pb.SimpleReq) (*pb.SimpleResp, error) {
	l := logic.NewTestRpcFuncLogic(ctx, s.svcCtx)
	return l.TestRpcFunc(in)
}
