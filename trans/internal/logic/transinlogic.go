package logic

import (
	"context"
	"log"

	"github.com/kevwan/gozero-dtm/trans/internal/svc"
	"github.com/kevwan/gozero-dtm/trans/pb"
	"github.com/tal-tech/go-zero/core/logx"
)

type TransInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransInLogic {
	return &TransInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransInLogic) TransIn(in *pb.AdjustInfo) (*pb.Response, error) {
	log.Printf("transfer in %d cents to %d", in.Amount, in.UserID)
	return &pb.Response{}, nil
}
