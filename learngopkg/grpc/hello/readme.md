# grpc/hello

- [grpc/hello](#grpchello)
  - [运行 hello](#运行-hello)
    - [运行 GRPC 服务器](#运行-grpc-服务器)
    - [运行 GRPC 客户端](#运行-grpc-客户端)
    - [停止项目（服务器 & 客户端）](#停止项目服务器--客户端)

## 运行 hello

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
