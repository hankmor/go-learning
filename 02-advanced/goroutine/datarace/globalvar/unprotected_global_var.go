package main

import (
	"net"
	"sync"
)

// 全局变量不受保护，容易发生竞态读写，不安全
// 确保并发环境下使用 mutex 保护之

var (
	service map[string]net.Addr
	mu      sync.Mutex
)

func RegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}

func SafeRegisterService(name string, addr net.Addr) {
	mu.Lock()
	defer mu.Unlock()
	service[name] = addr
}

func SafeLookupService(name string) net.Addr {
	mu.Lock()
	defer mu.Unlock()
	return service[name]
}

func main() {
}
