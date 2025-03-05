// consul_registry.go
package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

const (
	consulAddress = "localhost:8500"
	serviceName   = "ws-gateway"
)

func registerService(servicePort int) {
	config := api.DefaultConfig()
	config.Address = consulAddress
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	registration := &api.AgentServiceRegistration{
		ID:   *serviceId,
		Name: serviceName,
		Port: servicePort,
		Tags: []string{"v1", "zone-a"},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://localhost:%d/health", servicePort),
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		panic(err)
	}
	fmt.Println("Service registered with Consul")
}

func deregisterService() {
	config := api.DefaultConfig()
	config.Address = consulAddress
	client, _ := api.NewClient(config)
	client.Agent().ServiceDeregister(*serviceId)
	fmt.Println("Service deregistered from Consul")
}
