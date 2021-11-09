package main

import (
	"github.com/kevwan/gozero-dtm/dtmsdk"
	"github.com/kevwan/gozero-dtm/dtmsdk/dtmsdkimp"
	trans "github.com/kevwan/gozero-dtm/trans/pb"
)

func sagaCase() {
	d1 := []byte{}
	info1 := trans.AdjustInfo{Amount: 30, UserID: 1}
	d1, err := info1.XXX_Marshal(d1, true)
	dtmsdkimp.PanicIf(err != nil, err)
	d2 := []byte{}
	info2 := trans.AdjustInfo{Amount: 30, UserID: 2}
	d2, err = info2.XXX_Marshal(d2, true)
	dtmsdkimp.PanicIf(err != nil, err)

	saga := dtmsdk.NewSagaGrpc("localhost:59001", "gid1").
		// TODO dtm的设计是下面Add函数，采用的负载参数为proto.Message，但是gozero里面的消息，并不满足proto.Message接口
		// Add("127.0.0.1:50091/trans.TransSvc/TransOut", "127.0.0.1:50091/trans.TransSvc/TransOutRevert", &trans.AdjustInfo{Amount: 30, UserID: 1}).

		// TODO 目前是写死的127.0.0.1:50091，需要让用户从trans Client里面获取
		AddBin("127.0.0.1:50091/trans.TransSvc/TransOut", "127.0.0.1:50091/trans.TransSvc/TransOutRevert", d2).
		AddBin("127.0.0.1:50091/trans.TransSvc/TransIn", "127.0.0.1:50091/trans.TransSvc/TransInRevert", d1)
	err = saga.Submit()
	dtmsdkimp.PanicIf(err != nil, err)
}

func main() {
	sagaCase()
}
