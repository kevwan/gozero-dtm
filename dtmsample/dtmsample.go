package main

import (
	"flag"

	"github.com/kevwan/gozero-dtm/dtmdriverzero"
	"github.com/kevwan/gozero-dtm/dtmsdk"
	"github.com/kevwan/gozero-dtm/dtmsdk/dtmsdkimp"
	trans "github.com/kevwan/gozero-dtm/trans/pb"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc/resolver"
)

var (
	configFile = flag.String("f", "etc/config.yaml", "the config file")
	outReq     = &trans.AdjustInfo{Amount: 30, UserID: 1}
	inReq      = &trans.AdjustInfo{Amount: 30, UserID: 2}
)

func sagaRawCase() {
	saga := dtmsdk.NewSagaGrpc("localhost:59001", "gid1").
		Add("127.0.0.1:50091/trans.TransSvc/TransOut", "127.0.0.1:50091/trans.TransSvc/TransOutRevert", outReq).
		Add("127.0.0.1:50091/trans.TransSvc/TransIn", "127.0.0.1:50091/trans.TransSvc/TransInRevert", inReq)
	err := saga.Submit()
	dtmsdkimp.PanicIf(err != nil, err)
}

func sagaCase(transsvr string) {
	saga := dtmsdk.NewSagaGrpc(dtmsdk.DtmAddr, "gid2").
		Add(transsvr+"/trans.TransSvc/TransOut", transsvr+"/TransOutRevert", outReq).
		Add(transsvr+"/trans.TransSvc/TransIn", transsvr+"/TransInRevert", inReq)
	err := saga.Submit()
	dtmsdkimp.PanicIf(err != nil, err)
}

func main() {
	flag.Parse()

	var c dtmdriverzero.RpcConfig
	conf.MustLoad(*configFile, &c)

	// sagaRawCase()
	resolver.Register()
	target, err := dtmdriverzero.BuildTarget(c)
	logx.Must(err)
	sagaCase(target)
}
