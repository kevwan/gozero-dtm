package svc

import (
	"github.com/kevwan/gozero-dtm/order/internal/config"
	"github.com/kevwan/gozero-dtm/trans/transsvc"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Trans  transsvc.TransSvc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Trans:  transsvc.NewTransSvc(zrpc.MustNewClient(c.Trans)),
	}
}
