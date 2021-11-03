package main

import (
	"flag"
	"fmt"

	"blog/rpc/user/internal/config"
	"blog/rpc/user/internal/server"
	"blog/rpc/user/internal/svc"
	"blog/rpc/user/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUsersServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUsersServer(grpcServer, srv)

		switch c.Mode {
		case service.DevMode, service.TestMode:
			reflection.Register(grpcServer)
		default:
		}

	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
