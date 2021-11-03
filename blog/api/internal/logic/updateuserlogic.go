package logic

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserLogic {
	return UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req types.ReqUpdateUser) (*types.CommResp, error) {
	resp, err := l.svcCtx.User.Update(l.ctx, &user.User{Id: req.Id, Password: req.Password, Username: req.Username})
	if err != nil {
		return nil, err
	}

	return &types.CommResp{Ok: resp.Ok}, nil
}
