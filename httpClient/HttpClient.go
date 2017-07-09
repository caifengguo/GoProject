package main

import (
	"bytes"
	//"encoding/base64"
	"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
	//"unsafe"
)

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lTHIDFaceSDK
#include "THIDFaceSDK.h"
#include "THIDFaceSDK.hpp"
#include <stdlib.h>
*/
import "C"

// 图像
type IImage struct {
	Raw     []byte `json:"raw"`
	Channel uint32 `json:"channel"`
	Width   uint32 `json:"width"`
	Height  uint32 `json:"height"`
}

func main() {

	//初始化sdk
	/*	var device C.int = -1
		retval, _ := C.nn_THID_Init(device)
		if retval != C.THID_OK {
			log.Println("THID_Init SDK err:", retval)
			return
		} else {
			log.Println("THID_Init SDK :", retval)
		}*/

	req, err := http.Get("http://192.168.30.123:8088/Serverhost/SingleExtractFeat")
	defer req.Body.Close()
	if err != nil {
		log.Println("http connect fail!")
	} else {
		log.Println("http connect success!")
	}

	/*
		//转换成灰度图
		var image C.THID_Image
		file := C.CString("1.jpg")
		ret := C.nn_THID_DecodeFile(&image, file)
		if ret != C.THID_OK {
			log.Println("THID_DecodeJPG call fail!")
			return
		} else {
			log.Println("THID_DecodeJPG call success!")
		}

		str := C.GoString((*C.char)(unsafe.Pointer(image.raw)))
		log.Println(str)
		//编码转换成Base64
		Base64Str := base64.StdEncoding.EncodeToString([]byte(str))
		log.Println(Base64Str)
		ff, _ := os.Create("22.dat") //创建文件
		io.WriteString(ff, Base64Str)
		ff.Close()
		//string转换成[]byte
		/*btye := C.GoBytes(unsafe.Pointer(&Base64Str), C.int(len(Base64Str)))
		log.Println(btye)

		//解码Base64
		btye, err := base64.StdEncoding.DecodeString(Base64Str)
		log.Println(btye, err)*/

	for i := 0; i < 10000; i++ {
		//构造发送对象
		var img IImage
		img.Width = uint32(100)
		img.Height = uint32(200)
		img.Channel = 1 //灰度图

		//将对象打包成Json格式
		Json, err := json.Marshal(img)
		if err != nil {
			log.Println("json.Marshal fail!")
			return
		}
		body := bytes.NewReader(Json)
		log.Println(body)

		client := &http.Client{}
		resp, _ := http.NewRequest("POST", "http://192.168.30.123:8088/Serverhost/SingleExtractFeat", body)
		resp.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		//发送
		respon, err := client.Do(resp)
		if err != nil {
			return
		}
		defer respon.Body.Close()
		data, err := ioutil.ReadAll(respon.Body)
		log.Println(string(data))
	}

	var anykey string
	log.Println("please input any key for exit ...")
	fmt.Scanln(&anykey)

	/*
		//释放sdk资源
		retval = C.THID_Uninit()
		if retval != C.THID_OK {
			log.Println("THID_Uninit err:", retval)
		}
	*/
}
