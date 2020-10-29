package main

import "fmt"

//go:generate protoc --go_out=. hello/hello.proto

func main() {
	go Serve()

	var name string
	for {
		fmt.Scan(&name)
		fmt.Println("Server: " + SayHello(name))
	}
}
