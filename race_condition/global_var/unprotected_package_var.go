package global_var

import (
	"net"
	"sync"
)

var (
	service map[string]net.Addr = make(map[string]net.Addr)
	mutex sync.Mutex
)

func RegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func ProtectRegisterService(name string, addr net.Addr) {
	mutex.Lock()
	defer mutex.Unlock()
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}
