package logic

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteUserLogic {
	return DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req types.ReqUserId) (*types.CommResp, error) {
	resp, err := l.svcCtx.User.Delete(l.ctx, &user.ReqUserId{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.CommResp{Ok: resp.Ok}, nil
}
