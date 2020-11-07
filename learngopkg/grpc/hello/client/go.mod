module github.com/1000Delta/gopractice/learngopkg/grpc/hello/helloclient

go 1.13

replace github.com/1000Delta/gopractice/learngopkg/grpc/hello/helloserver => ../server

require (
	github.com/1000Delta/gopractice/learngopkg/grpc/hello/helloserver v0.0.0
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.0 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/grpc/examples v0.0.0-20200626195603-c95dc4da23cb // indirect
)
