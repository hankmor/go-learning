package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

const (
	Topic = "topic-test"
)

func main() {
	// 配置消费者
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 创建消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
	defer consumer.Close()

	// 指定分区消费者
	partitionConsumer, err := consumer.ConsumePartition(Topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	// 消费消息
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: partition: %d, key=%s, value=%s, offset=%d\n", msg.Partition, string(msg.Key), string(msg.Value), msg.Offset)
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %v", err)
		}
	}
}
