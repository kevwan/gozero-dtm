package main

import (
	"context"
	"flag"
	"log"

	"github.com/kevwan/gozero-dtm/order/ordersvc"
	order "github.com/kevwan/gozero-dtm/order/pb"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
)

var configFile = flag.String("f", "config.yaml", "config file")

type Config struct {
	Order zrpc.RpcClientConf
}

func main() {
	flag.Parse()
	logx.Disable()

	var c Config
	conf.MustLoad(*configFile, &c)

	cli := zrpc.MustNewClient(c.Order)
	svc := ordersvc.NewOrderSvc(cli)
	svc.Transfer(context.Background(), &order.TransferInfo{Amount: 30, FromID: 1, ToID: 2})
	log.Printf("transfer finished\n")
}
