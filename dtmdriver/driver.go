package dtmdriver

import "fmt"

// Driver interface to do service register and discover
type Driver interface {
	GetName() string
	RegisterGrpcResolver()
	RegisterService(url string, endpoint string) error
}

var (
	drivers = map[string]Driver{}
)

// Register used by each driver writer
func Register(driver Driver) {
	drivers[driver.GetName()] = driver
}

func MustGetDriver(name string) Driver {
	v := drivers[name]
	if v == nil {
		panic(fmt.Errorf("no dtm driver with name: %s has been registered", name))
	}
	return v
}
