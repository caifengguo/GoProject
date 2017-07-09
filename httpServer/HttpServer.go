package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func loop(j int) {
	for i := 1; i < j; i++ {
		fmt.Println(i)
	}
}

var count int = 0

func Hello(w http.ResponseWriter, req *http.Request) {
	log.Println("guo  ")
	w.Write([]byte("Hello World"))

	if count == 0 {
		loop(10)
		count++
	} else {
		loop(15)
		count = 0
	}

	boby, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		fmt.Println(string(boby))
	}
}

func main() {
	http.HandleFunc("/hello", Hello)
	err := http.ListenAndServe("192.168.30.123:8088", nil)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("http server success!")
	}

}
