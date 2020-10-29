package main

import (
	"fmt"
	"net/rpc"

	pb "github.com/1000Delta/gopractice/learngopkg/protobuf/hello"
)

// SayHello 在客户端调用，模拟本地调用 SayHello 的方法
func SayHello(name string) string {
	c, err := rpc.DialHTTPPath("tcp", ":8080", "/hello")
	if err != nil {
		panic(err)
	}

	var rsp pb.HelloResponse
	err = c.Call("HelloService.SayHello", &pb.HelloRequest{Name: name}, &rsp)
	if err != nil {
		fmt.Println(err)
	}

	return rsp.Msg
}
