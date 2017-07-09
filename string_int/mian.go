package main

import (
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	//获取当前工作路径
	dirctory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}

	//根据逻辑分库id创建对应的文件夹(0~99)
	for i := 0; i < 100; i++ {
		var Imagepath string
		dirctory += "\\"
		Imagepath = dirctory + strconv.Itoa(10)
		Imagepath += "\\"
		Imagepath = Imagepath + strconv.Itoa(i)
		//判断图片文件夹是否存在
		_, err := os.Stat(Imagepath)
		if os.IsNotExist(err) {
			Direrr := os.MkdirAll(Imagepath, 0666)
			if Direrr != nil {
				return
			}

		}
	}
}
