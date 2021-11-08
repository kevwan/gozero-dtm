package main

import (
	"flag"
	"fmt"
	"syscall"
	"time"

	"github.com/kevwan/gozero-dtm/trans/internal/config"
	"github.com/kevwan/gozero-dtm/trans/internal/server"
	"github.com/kevwan/gozero-dtm/trans/internal/svc"
	"github.com/kevwan/gozero-dtm/trans/pb"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/trans.yaml", "the config file")

func main() {
	flag.Parse()
	logx.Disable()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewTransSvcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterTransSvcServer(grpcServer, srv)

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
