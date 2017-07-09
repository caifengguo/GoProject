package main

import (
	"os"
	"strconv"
)

func main() {
	var Imagepath string = "d:/temp"
	Imagepath = Imagepath + "/" + strconv.Itoa(int(1234))
	//不存在，则创建
	err := os.MkdirAll(Imagepath, 0666)
	if err != nil {
		return
	}
}
