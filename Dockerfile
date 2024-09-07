FROM ubuntu:20.04

# 设置时区，避免交互式安装中的时区选择
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Shanghai

RUN apt update -y
RUN apt install gcc ca-certificates -y && update-ca-certificates
RUN apt install rsync ssh -y

ADD https://golang.google.cn/dl/go1.23.0.linux-arm64.tar.gz /
RUN tar -zxvf go1.23.0.linux-arm64.tar.gz

# 设置环境变量
ENV GO_VERSION 1.23
ENV GOROOT /go
ENV PATH /root/go/bin:/go/bin:/usr/local/go/bin:$PATH
RUN go install github.com/gogf/gf/cmd/gf/v2@latest

#RUN go mod download github.com/gogf/gf/v2@v2.7.2 && \
#    go mod download github.com/torykit/go-modbus@v1.0.0 && \
#    go mod download github.com/fsnotify/fsnotify@v1.7.0


# 创建工作目录并将当前目录挂载到容器的家目录
WORKDIR /root/work
# 挂载当前目录到容器的家目录
VOLUME ["/root/work"]

#COPY . .  # 再启动的时候把 当前目录挂载到容器的目录下
# 设置默认的执行命令为一个bash终端
CMD ["./docker_run.sh"]
#CMD ["/bin/bash"]