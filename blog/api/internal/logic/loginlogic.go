package logic

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/rpc/user/users"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *LoginLogic) Login(req types.ReqUser) (*types.RespLogin, error) {
	resp, err := l.svcCtx.User.Login(l.ctx, &users.LoginReq{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &types.RespLogin{Token: resp.Token}, nil
}
