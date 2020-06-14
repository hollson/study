#!/usr/bin/env bash

# 仅生成request和response对象
protoc --go_out=. hello.proto

# 生成request和response，
# 生成客户端/服务端接口
# 生成客户端实例和服务端注册方法
protoc --go_out=plugins=grpc:./ hello.proto