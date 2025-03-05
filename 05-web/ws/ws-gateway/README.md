websocket模板demo，仅仅展示基本的功能，包括:

- 服务注册、发现、健康检查, 这里通过 consul 实现
- 故障转移
- 负载均衡，这里通过 nginx + consul-template 实现


# 启动基础设施

```bash
# 启动Consul开发模式
consul agent -dev
# 启动Redis
docker run -p 6379:6379 redis
```

# 编译运行网关节点

```bash
go build -o gateway
./gateway
```

启动多个节点:

```bash
./gateway -p 8081 -i ws-gateway-node2
```

# 验证服务注册

```bash
curl <http://localhost:8500/v1/agent/services>
```

# 测试WebSocket连接

```javascript
// 客户端测试代码, 如果是集群环境这里配置nginx代理地址
const ws = new WebSocket('ws://localhost:8080/ws');
ws.onmessage = (event) => {
    console.log('Received:', event.data);
};
```

测试页面:

```bash
http://localhost:8080/static
```

# 连接负载均衡

```nginx
# Nginx配置示例

upstream websocket {
    server 192.168.1.10:8080;
    server 192.168.1.11:8080;

    # 使用最少连接算法
    least_conn;
}

server {
    listen 80;

    location /ws {
        proxy_pass http://websocket;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "Upgrade";
        proxy_set_header Host $host;
    }
}
```

## 使用consul-template动态管理nginx配置

### 安装consul-template

```bash
brew install consul-template
```

### 编辑consul-template模板

```nginx
# 这是一个consul-template模板文件
upstream websocket_servers {
  # 获取所有consul服务实例
{{ range service "ws-gateway" }}
    server {{ .Address }}:{{ .Port }};
{{ else }}
    server 127.0.0.1:9000; # 如果没有可用服务，默认 fallback
{{ end }}
}

server {
    listen 8889;

    location /ws {
        proxy_pass http://websocket_servers;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "Upgrade";
    }
}
```

### 启动consule-template

```bash
consul-template \
    -consul-addr=localhost:8500 \
    -template="/Users/hank/shell/ws-gateway.conf.ctmpl:/opt/homebrew/etc/nginx/servers/ws-gateway.conf:nginx -s reload" \
    -log-level=info
```

当consul中的服务状态发生变化，会自动更新nginx配置

# Redis集群配置

```go
redisClient = redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{
        "redis-node1:6379",
        "redis-node2:6379",
        "redis-node3:6379",
    },
    Password: "your_password",
})
```

# 性能调优参数

```go
// 调整Go运行参数
func init() {
    // 设置最大处理器数
    runtime.GOMAXPROCS(runtime.NumCPU())

    // 调整网络参数
    syscall.Setenv("GODEBUG", "netdns=go")
}
```
