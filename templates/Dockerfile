# 第一阶段：构建 Go 应用程序
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 Go 模块文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/myapp ./cmd/myapp

# 第二阶段：运行 Go 应用程序
FROM alpine:latest AS runtime

# 安装必要的运行时依赖（如果需要）
RUN apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /bin/myapp /app/myapp

# 暴露端口
EXPOSE 8080

# 设置环境变量（可选）
ENV ENV=production

# 运行应用程序
CMD ["/app/myapp"]