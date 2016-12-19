package common

import (
	"html/template"
	"net/http"
)

type HomeHandler struct {
	FileName string
}

func (a HomeHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(a.FileName)
	if err != nil {
		panic(err)
	}
	data := []byte{}
	err = tmpl.Execute(resp, data)
	if err != nil {
		panic(err)
	}
}
