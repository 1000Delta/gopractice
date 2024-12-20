package main

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
	"time"
)

func TestBaseRPC(t *testing.T) {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("Dial rpc error: " + err.Error())
	}

	var reply string
	err = client.Call(HelloServiceName+".Hello", "Li Hua", &reply)
	if err != nil {
		log.Fatal("Call rpc service error: " + err.Error())
	}

	log.Print("RPC return: " + reply)
}

func TestJSONRPC(t *testing.T) {
	client, err := jsonrpc.Dial("tcp", ":8081")
	if err != nil {
		log.Fatal("Dial rpc error: " + err.Error())
	}
	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatalf("TCP client close error: %s", err.Error())
		}
	}()

	var reply string
	err = client.Call(BackServiceName+".Back", time.Now().String(), &reply)
	if err != nil {
		log.Fatal("Call rpc service error: " + err.Error())
	}

	log.Print("RPC return: " + reply)
}
