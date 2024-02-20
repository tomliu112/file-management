# 基础镜像
FROM golang:1.19

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

# 切换到app目录
WORKDIR /app
# 将源码复制到app目录
COPY . .

# 编译源码
RUN go build -o main .

# 运行
CMD ["./main"]