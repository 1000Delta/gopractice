package main

import (
	"context"
	"os"
	"time"

	"google.golang.org/grpc/credentials"

	pb "github.com/1000Delta/gopractice/learngopkg/grpc/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	address = "127.0.0.1:52081"
)

func main() {
	// 设置 stdout 为 logger 输出
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr))

	// 配置*客户端* TLS 证书
	creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "helloserver")
	if err != nil {
		grpclog.Fatalln(err)
	}

	// grpc 链接 使用 TLS 证书
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("Dial conn %v error: %v", address, err.Error())
	}

	client := pb.NewHelloClient(conn)

	ticker := time.NewTicker(time.Second)

	for {
		rsp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Delta"})
		if err != nil {
			grpclog.Errorf("request error: %v", err.Error())
		}
		grpclog.Infoln("grpc response: " + rsp.Message)

		select {
		case <-ticker.C:
		}
	}

}
