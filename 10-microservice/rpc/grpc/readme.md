## 步骤

1. 安装 protoc
```shell
sudo brew install protoc
```

安装不起可以到github下载对应平台的releases二进制文件安装。

安装go plugin

```shell
$ go install google.golang.org/protobuf/command/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/command/protoc-gen-go-grpc@v1.2
```
使用文档：https://developers.google.com/protocol-buffers/docs/reference/go-generated

更新 PATH 让 protoc 可以找到 plugin

```shell
export PATH="$PATH:$(go env GOPATH)/bin"
```

2. 定义服务

创建 `arith.proto` 文件，编写 proto 代码来定义服务接口

3. 生成gRpc go代码

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    arith.proto
```

可以看到生成了两个文件：`arith.pb.go`、`arith_grpc.pb.go`，前者定义了客户端、服务端之间的通信消息，也就是参数和响应，后者定义了客户端和服务端对象以及它们之间的调用

4. 编写 server 代码

server代码需要开启服务，并实现.proto文件定义的接口

5. 编写 client 代码

client 连接 server，并调用server的方法，获得结果
