package main

import (
	"fmt"
	/*"io/ioutil"
	"log"
	"net/http"*/
	"net/url"
)

func main() {
	//获取底库照片
	//u, _ := url.Parse("http://192.168.30.123:8088/FetchImage?logicid=1&templateid=1003")
	//获取采集照片
	//u, _ := url.Parse("http://192.168.30.123:8088/FetchCaptureImage?logid=3&havetime=1494932150")
	//获取轨迹照片
	//u, _ := url.Parse("http://192.168.30.123:8088/FetchTrialImage?jobid=145")
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(u.User.String())

	/*
		res, err := http.Get(u.String())
		if err != nil {
			log.Fatal(err)
			return
		}

		result, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%s", result)*/
}
