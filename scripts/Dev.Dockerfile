FROM golang:1.19-alpine

WORKDIR /app

# 构建开发环境
COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify

COPY src/ ./