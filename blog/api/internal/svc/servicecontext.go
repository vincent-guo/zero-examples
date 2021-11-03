package svc

import (
	"blog/api/internal/config"
	"blog/rpc/user/users"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   users.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   users.NewUsers(zrpc.MustNewClient(c.User)),
	}
}
