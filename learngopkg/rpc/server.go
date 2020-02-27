package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
)

// HelloServiceName is the name of HelloService
const (
	HelloServiceName = "github.com/1000delta/gopractice/learngopkg/rpc/HelloService"
	BackServiceName = "github.com/1000delta/gopractice/learngopkg/rpc/BackService"
)

var wg *sync.WaitGroup

// HelloServiceInterface  bind the service implement Hello method
type HelloServiceInterface interface {
	Hello(request string, response *string) error
}

// RegisterHelloService register the service implement HelloService interface
func RegisterHelloService(srv HelloServiceInterface) error {
	err := rpc.RegisterName(HelloServiceName, srv)
	if err != nil {
		return err
	}
	return nil
}

// FirstService provide a rpc service
type FirstService struct{}

// Hello response hello and request data
func (f *FirstService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// BaseRPC provide base RPC service server
func BaseRPC() {
	defer wg.Done()

	BaseRPCHost := ":8080"
	// 注册rpc服务
	err := RegisterHelloService(&FirstService{})
	if err != nil {
		log.Fatal("Register rpc service failed: " + err.Error())
	}

	listener, err := net.Listen("tcp", BaseRPCHost)
	if err != nil {
		log.Fatal("Listen tcp port " + BaseRPCHost + " error: " + err.Error())
	}

	log.Print("Listening " + BaseRPCHost)
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept port listener error: " + err.Error())
	}

	rpc.ServeConn(conn)
	log.Print("BaseRPC service over")
}

// BackServiceInterface bind the service implement Back method
type BackServiceInterface interface{
	Back(request string, response *string) error
}

// RegisterBackService register BackService
func RegisterBackService(srv BackServiceInterface) error {
	err := rpc.RegisterName(BackServiceName, srv)
	if err != nil {
		return err
	}
	return nil
}

// SecondService provide rpc service 
type SecondService struct{}

// Back return the info of request
func (s *SecondService) Back(request string, response *string) error {
	*response = "Request: " + request
	return nil
}

// JSONRPC 提供了json rpc服务
func JSONRPC() {
	defer wg.Done()

	JSONRPCHost := ":8081"
	// 注册rpc服务
	err := RegisterBackService(&SecondService{})
	if err != nil {
		log.Fatal("Register rpc service failed: " + err.Error())
	}

	listener, err := net.Listen("tcp", JSONRPCHost)
	if err != nil {
		log.Fatal("Listen tcp port " + JSONRPCHost + " error: " + err.Error())
	}

	log.Print("Listening " + JSONRPCHost)
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept port listener error: " + err.Error())
	}

	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	log.Print("JSONRPC service over")
}

func main() {
	// 配置logger
	log.SetPrefix("[HelloRPCServer]")
	log.Print("Starting...")

	// 阻塞
	wg = &sync.WaitGroup{}

	// run RPC service server
	wg.Add(1)
	go BaseRPC()

	wg.Add(1)
	go JSONRPC()

	wg.Wait()
}
