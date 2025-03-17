package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
)

const (
	Topic = "topic-test"
)

// ConsumerGroupHandler 实现消费者组处理逻辑
type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Received message: partition: %d, key=%s, value=%s, offset=%d\n", msg.Partition,
			string(msg.Key), string(msg.Value), msg.Offset)
		session.MarkMessage(msg, "") // 标记消息已处理
	}
	return nil
}

func main() {
	// 配置消费者组
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 创建消费者组
	group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "my-group", config)
	if err != nil {
		log.Fatalf("Failed to start consumer group: %v", err)
	}
	defer group.Close()

	// 处理信号以优雅退出
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			err := group.Consume(ctx, []string{Topic}, ConsumerGroupHandler{})
			if err != nil {
				log.Printf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	// 捕获退出信号
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm
	cancel()
	wg.Wait()
}
