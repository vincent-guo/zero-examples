package logic

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/rpc/user/users"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllUsersLogic {
	return GetAllUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUsersLogic) GetAllUsers() ([]types.User, error) {
	resp, err := l.svcCtx.User.GetAll(l.ctx, &users.ReqGetAll{})
	if err != nil {
		return nil, err
	}

	res := make([]types.User, len(resp.Users))
	for i := 0; i < len(resp.Users); i++ {
		item := resp.Users[i]
		res[i] = types.User{
			Id:       item.Id,
			Username: item.Username,
			Password: item.Password,
		}
	}

	return res, nil
}
