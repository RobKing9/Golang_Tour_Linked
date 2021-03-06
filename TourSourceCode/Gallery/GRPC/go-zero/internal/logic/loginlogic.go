package logic

import (
	"context"

	"os_bufio_ioutil_file/ACode/Golang/GRPC/go-zero/internal/svc"
	"os_bufio_ioutil_file/ACode/Golang/GRPC/go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line

	return
}
