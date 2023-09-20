FROM golang:1.19-alpine

WORKDIR /app

# 指定gopath，给IDE做提示用
ENV GOPATH /app/.gopath

# 安装zsh
RUN apk update --no-cache && \
    apk add zsh git && \
    wget -O install.sh https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh && \
    chmod +x ./install.sh && \
    ./install.sh
