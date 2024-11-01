# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.21 AS builder

# 设置工作目录
WORKDIR /app

# 复制项目文件到工作目录
COPY . .

# 在这里可以添加其他构建步骤，例如运行 `go build`，如果需要的话

# 创建一个新的镜像，减少体积
FROM alpine:latest

ENV log_level=info\
    log_output_filename='web_chat'

# 设置工作目录
WORKDIR /app

# 复制 output 文件夹下的所有文件到 /app
COPY --from=builder /app/output/ .

# 使主程序可执行
RUN chmod +x ./bin/web_chat \
    mkdir /log

# 设置容器启动时执行的命令
CMD ["./bin/web_chat"]
