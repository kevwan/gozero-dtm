package dtmsdkimp

import (
	context "context"
	"log"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// TransBase base for all trans
type TransBase struct {
	Gid        string `json:"gid"`
	TransType  string `json:"trans_type"`
	Dtm        string `json:"-"`
	CustomData string `json:"custom_data,omitempty"`

	Steps       []map[string]string `json:"steps,omitempty"`    // use in MSG/SAGA
	Payloads    []string            `json:"payloads,omitempty"` // used in MSG/SAGA
	BinPayloads [][]byte            `json:"-"`
	Op          string              `json:"-"` // used in XA/TCC

	QueryPrepared string `json:"query_prepared,omitempty"` // used in MSG
}

// Call call an operation
func (s *TransBase) Call(operation string) error {
	reply := emptypb.Empty{}
	log.Printf("Calling operation %s with params %v\n", operation, s)
	return MustGetConn(s.Dtm).Invoke(context.Background(), "/dtmsdkimp.DtmSvc/"+operation, &DtmRequest{
		Gid:           s.Gid,
		TransType:     s.TransType,
		QueryPrepared: s.QueryPrepared,
		CustomedData:  s.CustomData,
		BinPayloads:   s.BinPayloads,
		Steps:         MustMarshalString(s.Steps),
	}, &reply)
}
