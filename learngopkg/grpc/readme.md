# grpc 练习

## Arch

- [`hello`](./hello/): grpc 基础练习，包含服务器和客户端

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
