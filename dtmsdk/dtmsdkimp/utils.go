package dtmsdkimp

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/tal-tech/go-zero/zrpc"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// MustProtoMarshal must version of proto.Marshal
func MustProtoMarshal(msg proto.Message) []byte {
	b, err := proto.Marshal(msg)
	PanicIf(err != nil, err)
	return b
}

// PanicIf name is clear
func PanicIf(cond bool, err error) {
	if cond {
		panic(err)
	}
}

// MustGetConn get grpc conn
func MustGetConn(grpcServer string) *grpc.ClientConn {
	cli, err := zrpc.NewClientWithTarget(grpcServer)
	PanicIf(err != nil, err)
	return cli.Conn()
}

// GetServerAndMethod 将grpc的url分解为server和method
func GetServerAndMethod(grpcURL string) (string, string) {
	// caution on wrong and complex grpcURL
	u, err := url.Parse(grpcURL)
	if err != nil {
		panic(err)
	}
	index := strings.IndexByte(u.Path[1:], '/') + 1
	server := u.Scheme + "://" + u.Host + u.Path[:index]
	method := u.Path[index:]
	return server, method
}

// MustMarshalString checked version for marshal
func MustMarshalString(v interface{}) string {
	b, err := json.Marshal(v)
	PanicIf(err != nil, err)
	return string(b)
}

type rawCodec struct{}

func (cb rawCodec) Marshal(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}

func (cb rawCodec) Unmarshal(data []byte, v interface{}) error {
	ba, _ := v.([]byte)
	for index, byte := range data {
		ba[index] = byte
	}
	return nil
}

func (cb rawCodec) Name() string { return "dtm_raw" }

// MustGetRawConn raw []byte in and out
func MustGetRawConn(grpcServer string) *grpc.ClientConn {
	cli, err := zrpc.NewClientWithTarget(grpcServer, zrpc.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(rawCodec{}))))
	PanicIf(err != nil, err)
	return cli.Conn()
}
