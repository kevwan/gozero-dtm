package dtmdriverzero

import (
	"github.com/kevwan/gozero-dtm/dtmdriver"
)

type zeroDriver struct{}

func (z *zeroDriver) GetName() string {
	return "gozero"
}

func (z *zeroDriver) RegisterGrpcResolver() {
	// TODO, 注册gozero支持的几种schema，直连/etcd/k8s
}

func (z *zeroDriver) RegisterService(url string, endpoint string) {
	// TODO 将endpoint上的服务，注册到url，仅需要完成etcd这种形式
}

func init() {
	dtmdriver.Register(&zeroDriver{})
}

// RegisterAsDefault register zero driver as default driver
func RegisterAsDefault() {
	dtmdriver.RegisterAsDefault(&zeroDriver{})
}
