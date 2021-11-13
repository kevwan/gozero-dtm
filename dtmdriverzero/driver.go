package dtmdriverzero

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kevwan/gozero-dtm/dtmdriver"
	"github.com/tal-tech/go-zero/core/discov"
)

const kindEtcd = "etcd"

type zeroDriver struct{}

func (z *zeroDriver) GetName() string {
	return "gozero"
}

func (z *zeroDriver) RegisterGrpcResolver() {
	// TODO, 注册gozero支持的几种schema，直连/etcd/k8s
}

func (z *zeroDriver) RegisterService(target string, endpoint string) error {
	// TODO 将endpoint上的服务，注册到url，仅需要完成etcd这种形式
	u, err := url.Parse(target)
	if err != nil {
		return err
	}

	switch u.Scheme {
	case kindEtcd:
		pub := discov.NewPublisher(strings.Split(u.Host, ","), u.Path, endpoint)
		pub.KeepAlive()
	default:
		return fmt.Errorf("unknown scheme: %s", u.Scheme)
	}

	return nil
}

func init() {
	dtmdriver.Register(&zeroDriver{})
}

// RegisterAsDefault register zero driver as default driver
func RegisterAsDefault() {
	dtmdriver.RegisterAsDefault(&zeroDriver{})
}
