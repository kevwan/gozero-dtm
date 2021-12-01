package dtmdriverzero

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kevwan/gozero-dtm/dtmdriver"
	"github.com/tal-tech/go-zero/core/discov"
)

const (
	DriverName = "go-zero"
	kindEtcd   = "etcd"
)

type zeroDriver struct{}

func (z *zeroDriver) GetName() string {
	return DriverName
}

func (z *zeroDriver) RegisterGrpcResolver() {
	// import is enough, nothing to do here
}

func (z *zeroDriver) RegisterService(target string, endpoint string) error {
	// TODO 将endpoint上的服务，注册到url，仅需要完成etcd这种形式
	u, err := url.Parse(target)
	if err != nil {
		return err
	}

	switch u.Scheme {
	case kindEtcd:
		pub := discov.NewPublisher(strings.Split(u.Host, ","), strings.TrimPrefix(u.Path, "/"), endpoint)
		pub.KeepAlive()
	default:
		return fmt.Errorf("unknown scheme: %s", u.Scheme)
	}

	return nil
}

func init() {
	dtmdriver.Register(&zeroDriver{})
}
