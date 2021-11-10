package main

import (
	"fmt"
	"log"
	"net"
	"time"

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

	// TARGET can also be "k8s://mynamespace/dtmservice:3456")
	// dtmdriver.GetDriver("zero").RegisterService("etcd://127.0.0.1:2379,127.0.0.1:2380/dtmservice", 59001, func(s *grpc.Server) {
	// 	s.RegisterService(&dtmsdkimp.DtmSvc_ServiceDesc, &svr.DtmServer{})
	// })
	time.Sleep(3000 * time.Second)
}
