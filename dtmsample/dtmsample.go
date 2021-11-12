package main

import (
	"github.com/kevwan/gozero-dtm/dtmdriverzero"
	"github.com/kevwan/gozero-dtm/dtmsdk"
	"github.com/kevwan/gozero-dtm/dtmsdk/dtmsdkimp"
	trans "github.com/kevwan/gozero-dtm/trans/pb"
)

var outReq = &trans.AdjustInfo{Amount: 30, UserID: 1}
var inReq = &trans.AdjustInfo{Amount: 30, UserID: 2}

func sagaRawCase() {
	saga := dtmsdk.NewSagaGrpc("localhost:59001", "gid1").
		Add("127.0.0.1:50091/trans.TransSvc/TransOut", "127.0.0.1:50091/trans.TransSvc/TransOutRevert", outReq).
		Add("127.0.0.1:50091/trans.TransSvc/TransIn", "127.0.0.1:50091/trans.TransSvc/TransInRevert", inReq)
	err := saga.Submit()
	dtmsdkimp.PanicIf(err != nil, err)
}

func sagaCase() {
	dtmdriverzero.RegisterAsDefault()
	dtmurl := "etcd://127.0.0.1:2379,127.0.0.1:2380/dtmservice"       // TARGET can also be "k8s://mynamespace/dtmservice:3456"
	transsvr := "etcd://127.0.0.1:2379,127.0.0.1:2380/trans.TransSvc" // TARGET can also be "k8s://mynamespace/trans.TransSvc:3456"
	saga := dtmsdk.NewSagaGrpc(dtmurl, "gid2").
		Add(transsvr+"/TransOut", transsvr+"/TransOutRevert", outReq).
		Add(transsvr+"/TransIn", transsvr+"/TransInRevert", inReq)
	err := saga.Submit()
	dtmsdkimp.PanicIf(err != nil, err)
}

func main() {
	sagaRawCase()
	sagaCase()
}
