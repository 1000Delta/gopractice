package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type testHandler struct{}

func (h *testHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}
	// m.Body is []byte type
	fmt.Println(string(m.Body))
	return nil
}

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "test_consume", config)
	if err != nil {
		panic(err)
	}
	defer consumer.Stop()

	consumer.AddHandler(&testHandler{})
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		panic(err)
	}
	// monitoring messages consume
	for true {
		// stats := consumer.Stats()
		// fmt.Printf("MR %d | MF %d\n", stats.MessagesReceived, stats.MessagesFinished)
		// if stats.MessagesReceived == stats.MessagesFinished {
		// 	break
		// }
	}
}
