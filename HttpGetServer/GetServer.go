package main

import (
	"fmt"
	//"html"
	"net/http"
	//"strings"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers   []Server
	ServersID string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8089", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.Form.Get("logicID"))
	fmt.Println(r.Form.Get("templateID"))
}
