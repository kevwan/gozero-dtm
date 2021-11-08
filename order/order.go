package main

import (
	"flag"
	"fmt"
	"syscall"
	"time"

	"github.com/kevwan/gozero-dtm/order/internal/config"
	"github.com/kevwan/gozero-dtm/order/internal/server"
	"github.com/kevwan/gozero-dtm/order/internal/svc"
	"github.com/kevwan/gozero-dtm/order/pb"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()
	logx.Disable()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewOrderSvcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderSvcServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	go func() {
		time.Sleep(time.Second * 5)
		syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	}()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
