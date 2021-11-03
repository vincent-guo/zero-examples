package logic

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/rpc/user/users"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req types.ReqUser) (*types.CommResp, error) {
	resp, err := l.svcCtx.User.Create(l.ctx, &users.ReqUser{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &types.CommResp{Ok: resp.Ok}, nil
}
