// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package server

import (
	"context"

	"github.com/kevwan/gozero-dtm/order/internal/logic"
	"github.com/kevwan/gozero-dtm/order/internal/svc"
	"github.com/kevwan/gozero-dtm/order/pb"
)

type OrderSvcServer struct {
	svcCtx *svc.ServiceContext
}

func NewOrderSvcServer(svcCtx *svc.ServiceContext) *OrderSvcServer {
	return &OrderSvcServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderSvcServer) Transfer(ctx context.Context, in *pb.TransferInfo) (*pb.Response, error) {
	l := logic.NewTransferLogic(ctx, s.svcCtx)
	return l.Transfer(in)
}
