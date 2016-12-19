// websocket示例
package main

import (
	"fmt"
	"net/http"
	"src/go/server_push/common"
	"time"

	"github.com/gorilla/websocket"
)

type WsHandler struct {
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var messageType = 1

func (a WsHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// 开
	conn, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		panic(err)
	}

	// 关
	defer func() {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	// 写
	go func() {
		for {
			conn.WriteMessage(messageType, []byte("hello i am ajax poll."))
			time.Sleep(5 * time.Second)
		}
	}()

	// 读
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(p))
	}
}

func main() {
	port := ":5920"
	home := common.HomeHandler{
		FileName: "tmpl.html",
	}
	http.Handle("/ws", WsHandler{})
	http.Handle("/", home)
	err := http.ListenAndServe(port, nil)
	fmt.Println(err)
}
