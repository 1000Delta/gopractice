package main

import (
	"log"
	"net"
	"net/rpc"

	pb "github.com/1000Delta/gopractice/learngopkg/protobuf/hello"
	pbcodec "github.com/mars9/codec"
)

// HelloService 定义一个类型，用来实现 rpc HelloServic
// 遵循 Go rpc Method 的定义，采用 protobuf 生成的 Request/Response 类型作为参数
type HelloService struct{}

// SayHello 方法供客户端调用
//
func (srv *HelloService) SayHello(req pb.HelloRequest, rsp *pb.HelloResponse) (err error) {
	msg := "Hello"
	if req.Name != "" {
		msg += " " + req.Name
	}
	rsp.Msg = msg + "!"
	return nil
}

// Serve 注册并处理 HTTP 请求，
// 通过 rpc.Server.HandleHTTP 指定处理的 HTTP 请求路径，在客户端需要通过 rpc.DialHTTPPath 来指定访问的路径
// TODO 使用 protobuf codec 编解码
func Serve() {
	srv := rpc.NewServer()
	err := pb.RegisterHelloService(srv, &HelloService{})
	if err != nil {
		log.Fatal("RegisterHelloService: ", err)
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Listen: ", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print("Accept: ", err)
			continue
		}
		// 使用编解码器处理
		codec := pbcodec.NewServerCodec(conn)
		go srv.ServeCodec(codec)
	}

	// go http.Serve(l, nil)
}
