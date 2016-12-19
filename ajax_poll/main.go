// ajax轮询示例
package main

import (
	"fmt"
	"net/http"
	"src/go/server_push/common"
)

type AjaxPollHandler struct {
}

func (a AjaxPollHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello i am ajax poll."))

}

func main() {
	port := ":5920"
	home := common.HomeHandler{
		FileName: "tmpl.html",
	}
	http.Handle("/ajaxpoll", AjaxPollHandler{})
	http.Handle("/", home)
	err := http.ListenAndServe(port, nil)
	fmt.Println(err)
}
