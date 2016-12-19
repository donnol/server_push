// long轮询示例
package main

import (
	"fmt"
	"net/http"
	"src/go/server_push/common"
	"time"
)

type LongPollHandler struct {
}

var c chan int

func (a LongPollHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(time.Now(), ": begin...")
	go func() {
		time.Sleep(5 * time.Second)
		c <- 1
	}()
	fmt.Println(time.Now(), " wait...")
	<-c

	fmt.Println(time.Now(), " end...")
	resp.Write([]byte("hello i am long poll."))
}

func main() {
	c = make(chan int)
	port := ":5920"
	home := common.HomeHandler{
		FileName: "tmpl.html",
	}
	http.Handle("/longpoll", LongPollHandler{})
	http.Handle("/", home)
	err := http.ListenAndServe(port, nil)
	fmt.Println(err)
}
