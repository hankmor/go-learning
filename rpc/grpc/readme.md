## 步骤

1. 安装 protoc
```shell
sudo brew install protoc
```

安装不起可以到github下载对应平台的releases二进制文件安装。

安装go plugin

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
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

4. server 代码

5. client 代码
