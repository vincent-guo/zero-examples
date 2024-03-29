package logic

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/rpc/user/users"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.ReqUserId) (*types.User, error) {
	resp, err := l.svcCtx.User.Get(l.ctx, &users.ReqUserId{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.User{
		Id:       resp.Id,
		Username: resp.Username,
		Password: resp.Password,
	}, nil
}
