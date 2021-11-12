package dtmsdk

import (
	"github.com/kevwan/gozero-dtm/dtmsdk/dtmsdkimp"
	"google.golang.org/protobuf/proto"
)

// DtmAddr etcd url for DtmServer
const DtmAddr = "etcd://127.0.0.1:2379,127.0.0.1:2380/dtmservice" // TODO 改为实际的etcd服务注册/发现的地址

// SagaGrpc struct of saga
type SagaGrpc struct {
	dtmsdkimp.TransBase
}

// NewSagaGrpc create a saga
func NewSagaGrpc(server, gid string) *SagaGrpc {
	return &SagaGrpc{TransBase: dtmsdkimp.TransBase{Dtm: server, Gid: gid}}
}

// Add add a saga step
func (s *SagaGrpc) Add(action, compensate string, payload proto.Message) *SagaGrpc {
	s.Steps = append(s.Steps, map[string]string{"action": action, "compensate": compensate})
	s.BinPayloads = append(s.BinPayloads, dtmsdkimp.MustProtoMarshal(payload))
	return s
}

// AddBin add a saga step with bin payload
func (s *SagaGrpc) AddBin(action, compensate string, payload []byte) *SagaGrpc {
	s.Steps = append(s.Steps, map[string]string{"action": action, "compensate": compensate})
	s.BinPayloads = append(s.BinPayloads, payload)
	return s
}

// Submit submit the saga trans
func (s *SagaGrpc) Submit() error {
	return s.Call("Submit")
}
