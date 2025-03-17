# 安装 Kafka

- 下载 Kafka：

从官网下载 Kafka（推荐版本：2.8.x 或更高）：Apache Kafka 下载

解压到本地，例如：tar -xzf kafka_2.13-3.6.0.tgz

# 启动 ZooKeeper（Kafka 依赖 ZooKeeper）：

```bash
cd kafka_2.13-3.6.0
bin/zookeeper-server-start.sh config/zookeeper.properties
```

- 启动 Kafka 服务：

```bash
bin/kafka-server-start.sh config/server.properties
```

创建 Topic：

```bash
bin/kafka-topics.sh --create --topic test-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

# 安装 Go 和依赖

```bash
go get github.com/IBM/sarama
```
