使用 `gin` 框架编写的一个 demo web server，用来测试 nginx 作为负载均衡时在服务出错、超时等情况下的行为。

1、nginx 在本地安装，并配置有 stream 做负载

```shell
upstream test-server {
    server 127.0.0.1:8002 fail_timeout=10s max_fails=1;
    server 127.0.0.1:8001;
}

server {
	listen 9000;
	server_name 127.0.0.1;
	location / {
		proxy_pass http://test-server;
	}
}
```

2、启动 go 服务，服务中注册有 `/err`, `/timeout` 接口，通过 nginx 代理 url 访问接口，观察 nginx 的行为，是 502？
504？还是其他错误信息？