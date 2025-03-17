package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
)

const (
	Topic = "topic-test"
)

func main() {
	// 配置生产者
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true // 返回成功消息
	config.Producer.Return.Errors = true    // 返回错误消息

	// 创建异步生产者
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start async producer: %v", err)
	}
	defer producer.Close()

	ticker := time.NewTicker(2 * time.Second)
	partition := 3
	// 消息通道
	go func() {
		for c := range ticker.C {
			msg := &sarama.ProducerMessage{
				Topic: Topic,
				// 根据partition生成random key，使消息路由到不同的partition中，这样不同的消费者都可以消费
				Key:   sarama.StringEncoder(fmt.Sprintf("%d", rand.Intn(partition))),
				Value: sarama.StringEncoder(fmt.Sprintf("Async message from Sarama, now: %v", c)),
			}
			producer.Input() <- msg
		}
	}()

	// 处理发送结果
	for {
		select {
		case success := <-producer.Successes():
			fmt.Printf("Message sent to partition %d at offset %d\n", success.Partition, success.Offset)
		case err := <-producer.Errors():
			log.Printf("Failed to send message: %v", err)
		case <-time.After(5 * time.Second):
			log.Println("Timeout waiting for response")
		}
	}
}
