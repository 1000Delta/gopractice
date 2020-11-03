package main

import (
	"log"
	"net"
	"net/rpc"

	pb "github.com/1000Delta/gopractice/learngopkg/protobuf/hello"
	pbcodec "github.com/mars9/codec"
)

// SayHello 在客户端调用，模拟本地调用 SayHello 的方法
func SayHello(name string) string {
	// c, err := pb.DialHelloService("tcp", ":8080")
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println("Dial: ", err)
	}
	c := rpc.NewClientWithCodec(pbcodec.NewClientCodec(conn))
	// if err != nil {
	// 	log.Println("DialHelloService: ", err)
	// }
	req := &pb.HelloRequest{Name: name}
	// reqData, err := proto.Marshal(req)
	// if err != nil {
	// 	log.Println("Marshal", err)
	// }

	var rsp pb.HelloResponse
	// var rsp []byte
	// err = c.Call("HelloService.SayHello", reqData, &rsp)
	err = c.Call("HelloService.SayHello", req, &rsp)
	if err != nil {
		log.Println("Call: ", err)
	}

	// return string(rsp)
	return rsp.Msg
}
