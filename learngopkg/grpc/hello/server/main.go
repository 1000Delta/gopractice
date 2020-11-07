package main

import (
	"context"
	"fmt"
	"net"
	"os"

	pb "github.com/1000Delta/gopractice/learngopkg/grpc/hello/helloserver/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	rsp := &pb.HelloResponse{}
	rsp.Message = fmt.Sprintf("Hello %s", in.Name)

	// logging
	grpclog.Infoln("Call: " + "HelloService.SayHello")

	return rsp, nil
}

const (
	// Address 是 TCP 连接地址
	Address = "127.0.0.1:52081"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Listen %v error: %v", Address, err.Error())
	}

	// 设置 stdout 为 logger 输出
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr))

	// TLS 认证
	creds, err := credentials.NewServerTLSFromFile("../server.pem", "../server.key")
	if err != nil {
		grpclog.Errorln("NewServerTLSFromFile: " + err.Error())
	}

	s := grpc.NewServer(grpc.Creds(creds)) // 参数为启用 TLS 认证
	// s := grpc.NewServer()

	pb.RegisterHelloServer(s, helloService{})

	grpclog.Println("listen on " + Address)
	s.Serve(listen)
}
