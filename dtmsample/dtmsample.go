package main

import (
	_ "github.com/kevwan/gozero-dtm/dtmdriverzero"
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
	transsvr := "etcd://127.0.0.1:2379/trans" // TODO 这个地方能够从trans的客户端便捷的获取url，而不是写死或手动拼凑
	saga := dtmsdk.NewSagaGrpc(dtmsdk.DtmAddr, "gid2").
		Add(transsvr+"/trans.TransSvc/TransOut", transsvr+"/TransOutRevert", outReq).
		Add(transsvr+"/trans.TransSvc/TransIn", transsvr+"/TransInRevert", inReq)
	err := saga.Submit()
	dtmsdkimp.PanicIf(err != nil, err)
}

func main() {
	// sagaRawCase()
	sagaCase()
}
