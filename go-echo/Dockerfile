# Dockerfile

# 使用 Go 编译器镜像作为构建阶段
FROM golang:1.18-alpine AS builder
WORKDIR /app
RUN go get github.com/labstack/echo/v4
COPY . .
RUN go build -o main .

# 使用 Alpine 作为最终运行环境
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]


#FROM hub.expvent.com.cn:1111/expvent/builder/golang:1.21 as builder
#LABEL maintainer=expvent@expvent.com
#
#WORKDIR /workspace
#COPY . .
#
#RUN --mount=type=cache,target=/go/pkg/mod,sharing=locked --mount=type=cache,target=/root/.cache,sharing=locked \
#    make -e BUILD_DEST_DIR=build build
#
#FROM hub.expvent.com.cn:1111/expvent/base/ubuntu:20.04
#LABEL maintainer=expvent@expvent.com
#
#ARG TARGET
#WORKDIR /
#ENV TARGET=${TARGET}
#USER root
#
#COPY --from=builder /workspace/build/* ./
#
#RUN chmod +x /${TARGET}
#
#ENTRYPOINT /${TARGET}
