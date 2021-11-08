package logic

import (
	"context"
	"log"

	"github.com/kevwan/gozero-dtm/order/internal/svc"
	"github.com/kevwan/gozero-dtm/order/pb"
	trans "github.com/kevwan/gozero-dtm/trans/pb"
	"github.com/tal-tech/go-zero/core/logx"
)

type TransferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferLogic {
	return &TransferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransferLogic) Transfer(in *pb.TransferInfo) (*pb.Response, error) {
	log.Printf("transfering %d cents from %d to %d", in.Amount, in.FromID, in.ToID)
	_, err := l.svcCtx.Trans.TransOut(l.ctx, &trans.AdjustInfo{Amount: 30, UserID: 1})
	if err != nil {
		log.Fatalf("Trans out called err: %v", err)
	}
	_, err = l.svcCtx.Trans.TransIn(l.ctx, &trans.AdjustInfo{Amount: 30, UserID: 1})
	if err != nil {
		log.Fatalf("Trans in called err: %v", err)
	}

	return &pb.Response{}, nil
}
