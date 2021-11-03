package logic

import (
	"context"

	"blog/rpc/model"
	"blog/rpc/user/internal/svc"
	"blog/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *user.User) (*user.CommResp, error) {
	err := l.svcCtx.Model.Update(model.User{
		Password: in.Password,
		Username: in.Username,
		Id:       in.Id,
	})
	if err != nil {
		return nil, err
	}

	return &user.CommResp{Ok: true}, nil
}
