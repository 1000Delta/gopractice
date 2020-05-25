package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

// ANSI Escape Sequence
const (
	AESSaveCursor    = "\033[s"
	AESRecoverCursor = "\033[u"
	AESCleanLine     = "\033[K"
)

// AESUpNLine 返回向上移动 n 行的代码
func AESUpNLine(n int) string { return fmt.Sprintf("\033[%dA", n) }

// AESDownNLine 返回向下移动 n 行的代码
func AESDownNLine(n int) string { return fmt.Sprintf("\033[%dB", n) }

const (
	pongDeadline = 60 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var write = make(chan []byte)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("./index.html")
		if err != nil {
			log.Printf("error: %v", err)
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			log.Printf("error: %v", err)
		}
		w.Write(data)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != "GET" {
		// 	http.Error(w, "method error", http.StatusMethodNotAllowed)
		// }
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("error : %v", err)
		}

		ws.SetReadDeadline(time.Now().Add(pongDeadline))
		ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongDeadline)); return nil })
		// read
		go func() {
			defer ws.Close()
			for {
				_, msg, err := ws.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Printf("error: %v", err)
					}
					break
				}
				// fmt.Print(AESSaveCursor)
				fmt.Printf(AESUpNLine(1)+"\r"+AESCleanLine+"[remote] %s\n[%s]\n\n", msg, time.Now().String())
				// fmt.Print(AESRecoverCursor)
				fmt.Print("input: ")
			}
		}()

		// write
		go func() {
			ticker := time.NewTicker(pongDeadline * 9 / 10)
			for {
				select {
				case data := <-write:
					ws.SetWriteDeadline(time.Now().Add(pongDeadline))
					err := ws.WriteMessage(websocket.TextMessage, data)
					if err != nil {
						log.Printf("error: %v", err)
					}
					// re disp
					// fmt.Print(AESSaveCursor)
					fmt.Printf(AESUpNLine(1)+"\r"+AESCleanLine+"[me] %s\n[%s]\n\n", data, time.Now().String())
					// fmt.Print(AESRecoverCursor)
					fmt.Print("input: ")
				case <-ticker.C:
					ws.SetWriteDeadline(time.Now().Add(pongDeadline))
					if err := ws.WriteMessage(websocket.PingMessage, nil); err != nil {
						return
					}
				default:
				}
			}
		}()

		go func() {
			fmt.Println("")
			var input []byte
			fmt.Print("input: ")
			for {
				fmt.Scan(&input)
				fmt.Print(AESUpNLine(1) + "\r" + AESCleanLine)
				write <- input
			}
		}()
	})

	_ = http.ListenAndServe(":8080", nil)
}
