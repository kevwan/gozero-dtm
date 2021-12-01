package dtmdriverzero

import (
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
)

const (
	// directScheme stands for direct scheme.
	directScheme = "direct"
	// etcdScheme stands for etcd scheme.
	etcdScheme = "etcd"
	// kubernetesScheme stands for k8s scheme.
	kubernetesScheme = "k8s"
	// endpointSepChar is the separator cha in endpoints.
	endpointSep = ","
)

func BuildTarget(c zrpc.RpcClientConf) (string, error) {
	if len(c.Endpoints) > 0 {
		return buildDirectTarget(c.Endpoints), nil
	} else if len(c.Target) > 0 {
		return c.Target, nil
	} else {
		if err := c.Etcd.Validate(); err != nil {
			return "", err
		}

		if c.Etcd.HasAccount() {
			discov.RegisterAccount(c.Etcd.Hosts, c.Etcd.User, c.Etcd.Pass)
		}

		return buildDiscovTarget(c.Etcd.Hosts, c.Etcd.Key), nil
	}
}

// buildDirectTarget returns a string that represents the given endpoints with direct schema.
func buildDirectTarget(endpoints []string) string {
	return fmt.Sprintf("%s:///%s", directScheme, strings.Join(endpoints, endpointSep))
}

// buildDiscovTarget returns a string that represents the given endpoints with discov schema.
func buildDiscovTarget(endpoints []string, key string) string {
	return fmt.Sprintf("%s://%s/%s", etcdScheme, strings.Join(endpoints, endpointSep), key)
}
