package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/kevwan/gozero-dtm/dtmdriver"
	"github.com/kevwan/gozero-dtm/dtmsdk"
	"github.com/kevwan/gozero-dtm/dtmsdk/dtmsdkimp"
	"github.com/kevwan/gozero-dtm/dtmsvr/svr"

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

	// TODO 把本地端口59001启动的DtmServer注册到etcd
	dtmdriver.GetDriver("zero").RegisterService(dtmsdk.DtmAddr, "localhost:59001")

	time.Sleep(3000 * time.Second)
}
