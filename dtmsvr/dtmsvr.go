package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"

	"github.com/kevwan/gozero-dtm/dtmdriver"
	_ "github.com/kevwan/gozero-dtm/dtmdriverzero"
	"github.com/kevwan/gozero-dtm/dtmsdk"
	"github.com/kevwan/gozero-dtm/dtmsdk/dtmsdkimp"
	"github.com/kevwan/gozero-dtm/dtmsvr/svr"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc"
)

func startRPCSvc(sd *grpc.ServiceDesc, svc interface{}, port int64) {
	s := grpc.NewServer()
	s.RegisterService(sd, svc)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("listen at port %d error: %v", port, err)
	}

	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("Serve error: %v", err)
		}
	}()
	log.Printf("service: %s serv at port %d\n", sd.ServiceName, port)
}

func main() {
	startRPCSvc(&dtmsdkimp.DtmSvc_ServiceDesc, &svr.DtmServer{}, 59001)
	driver := dtmdriver.MustGetDriver("zero")
	driver.RegisterGrpcResolver() // 服务器端，可能有多个driver，需要手动指定使用driver的Resolver
	// TODO 把本地端口59001启动的DtmServer注册到etcd
	// driver.RegisterService(dtmsdk.DtmAddr, "localhost:59001")
	u, err := url.Parse(dtmsdk.DtmAddr)
	logx.Must(err)
	if u.Scheme == "etcd" {
		pub := discov.NewPublisher(strings.Split(u.Host, ","), u.Path, "localhost:59001")
		pub.KeepAlive()
	}

	select {}
}
