package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// 全局变量不受保护，容易发生竞态读写，不安全
// 确保并发环境下使用 mutex 保护之

var (
	service     = make(map[string]net.Addr)
	mu          sync.Mutex
	safeService sync.Map
)

func UnsafeRegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func UnsafeLookupService(name string) net.Addr {
	return service[name]
}

func SafeRegisterService(name string, addr net.Addr) {
	mu.Lock()
	defer mu.Unlock()
	service[name] = addr
	fmt.Println(service)
}

func SafeLookupService(name string) net.Addr {
	mu.Lock()
	defer mu.Unlock()
	return service[name]
}

func SafeRegisterService1(name string, addr net.Addr) {
	safeService.Store(name, addr)
}

func SafeLookupService1(name string) net.Addr {
	v, ok := safeService.Load(name)
	if ok {
		return v.(net.Addr)
	} else {
		return nil
	}
}

func main() {
	key := "net"
	_, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		UnsafeRegisterService(key, ipv4Net)
	}()
	time.Sleep(time.Second)
	ret := UnsafeLookupService(key)
	fmt.Println(ret)

	go func() {
		SafeRegisterService(key, ipv4Net)
	}()
	time.Sleep(time.Second)
	fmt.Println(SafeLookupService(key))

	time.Sleep(time.Second)
	go func() {
		SafeRegisterService1(key, ipv4Net)
	}()
	time.Sleep(time.Second)
	fmt.Println(SafeLookupService1(key))
}
