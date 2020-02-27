package main

import (
	"log"
	"net/rpc"
	"testing"
)

func TestBaseRPC(t *testing.T) {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("Dial rpc error: " + err.Error())
	}

	var reply string
	err = client.Call("HelloService.Hello", "Li Hua", &reply)
	if err != nil {
		log.Fatal("Call rpc service error: " + err.Error())
	}

	log.Print("RPC return: " + reply)
}