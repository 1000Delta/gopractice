# grpc/hello-tls

- [grpc/hello-tls](#grpchello-tls)
  - [概述](#概述)
  - [运行 hello-tls](#运行-hello-tls)
    - [运行 GRPC 服务器](#运行-grpc-服务器)
    - [运行 GRPC 客户端](#运行-grpc-客户端)
    - [停止项目（服务器 & 客户端）](#停止项目服务器--客户端)

## 概述

基本 grpc 练习，创建了一组 server/client，使用了 grpc 的 tls 模块，对证书进行认证

## 运行 hello-tls

### 运行 GRPC 服务器

```bash
make server
```

日志重定向到了 `./server/server.log`

### 运行 GRPC 客户端

```bash
make client
```

日志重定向到了 `./client/client.log`

### 停止项目（服务器 & 客户端）

```bash
make stop
```
