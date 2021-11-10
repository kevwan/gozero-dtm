package dtmdriver

import (
	"fmt"

	"google.golang.org/grpc"
)

// Driver interface to do service register and discover
type Driver interface {
	GetName() string
	RegisterGrpcResolver()
	RegisterService(url string, localPort int, registerFn func(*grpc.Server))
}

var (
	drivers              = map[string]Driver{}
	defaultDriver Driver = nil
)

// Register register driver
func Register(driver Driver) {
	drivers[driver.GetName()] = driver
}

// RegisterAsDefault register driver as default. only one default is allowed
func RegisterAsDefault(driver Driver) error {
	if defaultDriver != nil && defaultDriver.GetName() != driver.GetName() {
		return fmt.Errorf("already registered a default driver: %s", defaultDriver.GetName())
	}
	defaultDriver = driver
	Register(driver)
	return nil
}

// GetDefault get the default driver
func GetDefault() Driver {
	return defaultDriver
}

// GetDriver get the driver by name
func GetDriver(name string) Driver {
	return drivers[name]
}
