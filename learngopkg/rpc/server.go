package main

import (
	"log"
	"net"
	"net/rpc"
)

// HelloService provide hello rpc service
type HelloService struct{}

// Hello response hello and request data
func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// BaseRPC provide base RPC service server
func BaseRPC() {
	// 注册rpc服务
	err := rpc.Register(&HelloService{})
	if err != nil {
		log.Fatal("Register rpc service failed:" + err.Error())
	}

	host := ":8080"
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal("Listen tcp port " + host + " error:" + err.Error())
	}

	log.Print("Listening " + host)
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept port listener error:" + err.Error())
	}
	
	rpc.ServeConn(conn)
}

func main() {
	// 配置logger
	log.SetPrefix("[HelloRPCServer]")
	log.Print("Starting...")

	// run RPC service server
	go BaseRPC()

}


