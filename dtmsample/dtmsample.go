package main

import (
	"github.com/kevwan/gozero-dtm/dtmdriverzero"
	"github.com/kevwan/gozero-dtm/dtmsdk"
	"github.com/kevwan/gozero-dtm/dtmsdk/dtmsdkimp"
	trans "github.com/kevwan/gozero-dtm/trans/pb"
	"google.golang.org/protobuf/proto"
)

func sagaCase() {
	dtmdriverzero.RegisterAsDefault()
	dtmurl := "etcd://127.0.0.1:2379,127.0.0.1:2380/dtmservice"       // TARGET can also be "k8s://mynamespace/dtmservice:3456"
	transsvr := "etcd://127.0.0.1:2379,127.0.0.1:2380/trans.TransSvc" // TARGET can also be "k8s://mynamespace/trans.TransSvc:3456"
	var req proto.Message = nil                                       // TARGET can be &trans.AdjustInfo{Amount: 30, UserID: 1}
	saga := dtmsdk.NewSagaGrpc(dtmurl, "gid2").
		Add(transsvr+"/TransOut", transsvr+"/TransOutRevert", req).
		Add(transsvr+"/TransIn", transsvr+"/TransInRevert", req)
	err := saga.Submit()
	dtmsdkimp.PanicIf(err != nil, err)
}

func main() {
	sagaCase()
}
