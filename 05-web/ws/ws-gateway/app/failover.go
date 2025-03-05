// failover.go
package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
)

// failover.go 完整实现
func watchServiceChanges(stopCh <-chan struct{}) {
	config := api.DefaultConfig()
	config.Address = consulAddress
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Consul client init failed:", err)
	}

	// 使用带超时的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 首次获取所有服务节点
	initialServices, _, err := client.Health().Service(serviceName, "", true, nil)
	if err == nil {
		checkFailedNodes(ctx, initialServices)
	}

	// 长轮询配置
	params := &api.QueryOptions{
		WaitTime:  30 * time.Second, // 长轮询等待时间
		WaitIndex: 0,
	}

	for {
		select {
		case <-stopCh:
			log.Println("Stopping service watcher")
			return
		default:
			services, meta, err := client.Health().Service(
				serviceName,
				"",
				true,
				params.WithContext(ctx),
			)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}
				log.Printf("Watch error: %v, retrying...", err)
				time.Sleep(2 * time.Second)
				continue
			}

			params.WaitIndex = meta.LastIndex
			checkFailedNodes(ctx, services)
		}
	}
}

// 改进后的节点检查
func checkFailedNodes(ctx context.Context, services []*api.ServiceEntry) {
	activeNodes := make(map[string]struct{})
	for _, s := range services {
		activeNodes[s.Service.ID] = struct{}{}
	}

	// 使用SCAN迭代避免KEYS命令阻塞
	iter := redisClient.Scan(ctx, 0, "ws:session:*", 100).Iterator()
	for iter.Next(ctx) {
		// 确保每次迭代都释放资源
		// defer func() {
		// 	iter.Close()
		// 	cancel()
		// }()
		key := iter.Val()
		nodeID, err := redisClient.HGet(ctx, key, "gateway_node").Result()
		if err != nil {
			continue
		}

		if _, exists := activeNodes[nodeID]; !exists {
			go migrateSession(key, nodeID)
		}
	}
	if err := iter.Err(); err != nil {
		log.Printf("Redis scan error: %v", err)
	}
}

func migrateSession(key, oldNode string) {
	// 实现会话迁移逻辑
	// 1. 通知客户端重新连接
	// 2. 更新Redis中的节点信息
	// 3. 清理旧节点数据
}
