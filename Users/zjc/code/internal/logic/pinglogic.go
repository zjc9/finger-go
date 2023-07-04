package logic

import (
	"context"

	"G2/Users/zjc/code/g2"
	"G2/Users/zjc/code/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *g2.Request) (*g2.Response, error) {
	// todo: add your logic here and delete this line

	return &g2.Response{}, nil
}
