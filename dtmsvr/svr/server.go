package svr

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/kevwan/gozero-dtm/dtmsdk/dtmsdkimp"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DtmServer struct {
	pb.UnimplementedDtmSvcServer
}

// func (d *DtmServer) RegisterBranch(ctx context.Context, in *pb.DtmBranchRequest) (*emptypb.Empty, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method RegisterBranch not implemented")
// }

func (d *DtmServer) Submit(ctx context.Context, in *pb.DtmRequest) (*emptypb.Empty, error) {
	log.Printf("handling Submit with %v\n", in)
	steps := []map[string]string{}
	err := json.Unmarshal([]byte(in.Steps), &steps)
	pb.PanicIf(err != nil, err)
	for i, step := range steps {
		url := step["action"]
		svr, method := pb.GetServerAndMethod(url)
		r := []byte{}
		log.Printf("invoking %s\n", url)
		err = pb.MustGetRawConn(svr).Invoke(context.Background(), method, in.BinPayloads[i], &r)
		pb.PanicIf(err != nil, err)
	}
	return &emptypb.Empty{}, nil
}
