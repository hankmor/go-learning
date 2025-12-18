# 多阶段构建：减小镜像体积
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# 最终镜像
FROM alpine:latest

WORKDIR /app

# 安装 CA 证书（用于 HTTPS 请求）
RUN apk --no-cache add ca-certificates

# 从 builder 阶段复制二进制文件
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
