# grpc/hello

- [grpc/hello](#grpchello)
  - [初始化](#初始化)
    - [证书生成](#证书生成)
  - [运行 GRPC 项目](#运行-grpc-项目)
    - [运行 GRPC 服务器](#运行-grpc-服务器)
    - [运行 GRPC 客户端](#运行-grpc-客户端)
    - [停止项目（服务器 & 客户端）](#停止项目服务器--客户端)

## 初始化

### 证书生成

使用 `make`：

```bash
make init
```

证书自定义信息：

```
-----
Country Name (2 letter code) [AU]:
State or Province Name (full name) [Some-State]:
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:helloserver
Email Address []:
```

重点在于 `Common Name`，在 [`google.golang.org/grpc/credentials.NewClientTLSFromFile`](https://pkg.go.dev/google.golang.org/grpc@v1.30.0/credentials#NewClientTLSFromFile) 中需要设置 `serverNameOverride` 参数值

## 运行 GRPC 项目

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
