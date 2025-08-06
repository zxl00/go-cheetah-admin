FROM golang:1.23.5 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o main  .
FROM alpine:3.21.3
WORKDIR /app
COPY --from=builder /app/main  .
COPY --from=builder /app/rbac_model.conf .
ENV TZ=Asia/Shanghai
RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
    && apk upgrade \
    && apk add --no-cache --virtual .build-deps \
    ca-certificates upx tzdata \
    && chmod +x /app/main \
    && mkdir /app/db \
    && chmod 777 /app/db
ENTRYPOINT [ "/app/main" ]
