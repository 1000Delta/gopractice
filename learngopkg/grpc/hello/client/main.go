package main

import (
	"context"
	"os"

	pb "github.com/1000Delta/gopractice/learngopkg/grpc/hello/helloserver/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	address = "127.0.0.1:52081"
)

func main() {
	// 设置 stdout 为 logger 输出
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr))

	// grpc 链接 参数忽略证书
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("Dial conn %v error: %v", address, err.Error())
	}

	client := pb.NewHelloClient(conn)

	rsp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Delta"})
	if err != nil {
		grpclog.Errorf("request error: %v", err.Error())
	}

	grpclog.Infoln("grpc response: " + rsp.Message)
}
