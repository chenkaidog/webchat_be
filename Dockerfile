# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.21

# 设置工作目录
WORKDIR /app

# 复制项目文件到工作目录
COPY output/ .

ENV log_level=debug\
    log_output_filename='web_chat'

# 使主程序可执行
RUN chmod +x ./bin/web_chat

# 设置容器启动时执行的命令
CMD ["./bin/web_chat"]
