package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func logErrorln(v ...interface{}) {
	log.Println("Error: ", fmt.Sprint(v...))
}

func logNoticeln(v ...interface{}) {
	log.Println("Notice: ", fmt.Sprint(v...))
}

func tcpServer() {
	lis, err := net.Listen("tcp", ":8880")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			logErrorln(err.Error())
		}
		go func(c net.Conn) {
			defer c.Close()

			data, err := ioutil.ReadAll(c)
			if err != nil {
				logErrorln(err.Error())
			}
			logNoticeln(data)
			// close cmd
			if string(data) == "close" {
				lis.Close()
				return
			}
			c.Write(data)
		}(conn)
	}
}

func tcpConn() {
	conn, err := net.Dial("tcp", ":8880")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()
	
	for {
		
	}
}