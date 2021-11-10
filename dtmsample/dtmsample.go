package main

import (
	"github.com/kevwan/gozero-dtm/dtmdriverzero"
	"github.com/kevwan/gozero-dtm/dtmsdk"
	"github.com/kevwan/gozero-dtm/dtmsdk/dtmsdkimp"
	trans "github.com/kevwan/gozero-dtm/trans/pb"
	"google.golang.org/protobuf/proto"
)

func sagaRawCase() {
	d1 := []byte{}
	info1 := trans.AdjustInfo{Amount: 30, UserID: 1}
	d1, err := info1.XXX_Marshal(d1, true)
	dtmsdkimp.PanicIf(err != nil, err)
	d2 := []byte{}
	info2 := trans.AdjustInfo{Amount: 30, UserID: 2}
	d2, err = info2.XXX_Marshal(d2, true)
	dtmsdkimp.PanicIf(err != nil, err)

	saga := dtmsdk.NewSagaGrpc("localhost:59001", "gid1").
		AddBin("127.0.0.1:50091/trans.TransSvc/TransOut", "127.0.0.1:50091/trans.TransSvc/TransOutRevert", d2).
		AddBin("127.0.0.1:50091/trans.TransSvc/TransIn", "127.0.0.1:50091/trans.TransSvc/TransInRevert", d1)
	err = saga.Submit()
	dtmsdkimp.PanicIf(err != nil, err)
}

func sagaCase2() {
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
	sagaRawCase()
	// sagaCase2()  // TARGET
}
