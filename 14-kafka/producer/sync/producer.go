package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

const (
	TOPIC = "topic-test"
)

func main() {
	// 配置生产者
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 等待所有副本确认
	config.Producer.Retry.Max = 5                    // 最大重试次数
	config.Producer.Return.Successes = true          // 返回成功消息

	// 创建同步生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start producer: %v", err)
	}
	// defer producer.Close()

	ctx := context.Background()
	ticker := time.NewTicker(2 * time.Second)
	go func(ctx context.Context) {
		for {
			select {
			case c := <-ticker.C:
				// 要发送的消息
				msg := &sarama.ProducerMessage{
					Topic: TOPIC,
					Key:   sarama.StringEncoder("key-1"),
					Value: sarama.StringEncoder(fmt.Sprintf("Hello, Kafka with Sarama, now: %v", c)),
				}

				// 发送消息
				partition, offset, err := producer.SendMessage(msg)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				} else {
					fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
				}
			case <-ctx.Done():
				fmt.Println("producer done")
                producer.Close()
				return
			}
		}
	}(ctx)
	time.Sleep(time.Second * 6)
	ctx.Done()
}
