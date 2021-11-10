package dtmdriverzero

import (
	"github.com/kevwan/gozero-dtm/dtmdriver"
	"google.golang.org/grpc"
)

type zeroDriver struct {
}

func (z *zeroDriver) GetName() string {
	return "gozero"
}

func (z *zeroDriver) RegisterGrpcResolver() {
	// TODO, register gozero grpc resolver
}

func (z *zeroDriver) RegisterService(url string, localPort int, registerFn func(*grpc.Server)) {
	// TODO register a service to url
}

func init() {
	dtmdriver.Register(&zeroDriver{})
}

// RegisterAsDefault register zero driver as default driver
func RegisterAsDefault() {
	dtmdriver.RegisterAsDefault(&zeroDriver{})
}
