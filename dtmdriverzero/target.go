package dtmdriverzero

import (
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc/resolver"
)

func BuildTarget(c RpcConfig) (string, error) {
	if len(c.Endpoints) > 0 {
		return resolver.BuildDirectTarget(c.Endpoints), nil
	} else if len(c.Target) > 0 {
		return c.Target, nil
	}

	if err := c.Etcd.Validate(); err != nil {
		return "", err
	}

	if c.Etcd.HasAccount() {
		discov.RegisterAccount(c.Etcd.Hosts, c.Etcd.User, c.Etcd.Pass)
	}

	return resolver.BuildDiscovTarget(c.Etcd.Hosts, c.Etcd.Key), nil
}
