package main

import (
	//"fmt"
	"os"

	//"hisign.com.cn/uuid"
)

func main() {
	/*	var UUID *uuid.ISnowflake
		UUID, _ = uuid.NewSnowflake(123)
		id, _ := UUID.GenerateID()
		fmt.Println(id)*/

	err := os.MkdirAll("D:/guocaifeng/123/12", 0777)
	if err != nil {
		return
	}
	return
}
