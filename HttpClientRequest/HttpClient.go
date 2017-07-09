package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 图像
type IImage struct {
	Channel int32  `json:"channel"`
	Width   int32  `json:"width"`
	Height  int32  `json:"height"`
	Raw     []byte `json:"raw"`
}

//存储底库照片结构体
type IStoreImageInfo struct {
	// 模板ID
	TemplateID int64 `json:"templateId"`
	// 分库ID
	LogicdbID int32 `json:"logicdbId"`
	// 照片
	Image IImage `json:"Image"`
}

//模板删除
type IDeleteTemplateID struct {
	// 分库ID
	LogicdbID int32 `json:"logicdbId"`
	// 模板ID
	TemplateID int64 `json:"templateId"`
}

//逻辑分库删除
type IDeleteLogicdbID struct {
	// 分库ID
	LogicdbID int32 `json:"logicdbId"`
}

//采集存储结构体
type ICaptureImageInfo struct {
	// 分库ID
	LogID int64 `json:"logId"`
	//时间
	HaveTime int64 `json:"HaveTime"`
	//照片
	Image IImage `json:"Image"`
}

//删除过期图片时间
type IDeleteDate struct {
	Days int64 `json:"days"`
}

//存储轨迹信息结构体
type IStoreTrialImage struct {
	JodID string `json:"jobID"`
	Image IImage `json:"image"`
}

func main() {
	//建立连接
	req, err := http.Get("http://192.168.30.123:8088/StoreCaptureImage")
	defer req.Body.Close()
	if err != nil {
		log.Println("http connect fail!")
	} else {
		log.Println("http connect success!")
	}

	/********底库图片处理**********/
	/*
		//存储底库照片
		//读取图片
		imagebuffer, errcode := ioutil.ReadFile("d:/1.jpg")
		if errcode != nil {
			fmt.Println("read image file fail")
			return
		}
		//将对象打包成Json格式
		imageinfo := &IStoreImageInfo{1003, 1, IImage{1, 0, 0, imagebuffer}}
		Json, err := json.Marshal(imageinfo)
		if err != nil {
			log.Println("json.Marshal fail!")
			return
		}
	*/
	/*
		//删除模板
		imageinfo := IDeleteTemplateID{1, 1003}
		Json, err := json.Marshal(imageinfo)
		if err != nil {
			log.Println("json.Marshal fail!")
			return
		}
	*/
	/*
		//删除分库
		imageinfo := IDeleteLogicdbID{1}
		Json, err := json.Marshal(imageinfo)
		if err != nil {
			log.Println("json.Marshal fail!")
			return
		}
	*/
	/*******采集端***********/

	//读取图片
	imagebuffer, errcode := ioutil.ReadFile("d:/1.jpg")
	if errcode != nil {
		fmt.Println("read image file fail")
		return
	}
	//将对象打包成Json格式
	imageinfo := &ICaptureImageInfo{3, 1494932150, IImage{1, 0, 0, imagebuffer}}
	Json, err := json.Marshal(imageinfo)
	if err != nil {
		log.Println("json.Marshal fail!")
		return
	}

	/*
		//删除采集图片
		Datainfo := IDeleteDate{1494932150}
		Json, err := json.Marshal(Datainfo)
		if err != nil {
			log.Println("json.Marshal fail!")
			return
		}
	*/
	/***********轨迹任务**************/
	/*
		//存储轨迹任务
		imagebuffer, errcode := ioutil.ReadFile("d:/1.jpg")
		if errcode != nil {
			fmt.Println("read image file fail")
			return
		}
		//将对象打包成Json格式
		imageinfo := &IStoreTrialImage{"145", IImage{1, 0, 0, imagebuffer}}
		Json, err := json.Marshal(imageinfo)
		if err != nil {
			log.Println("json.Marshal fail!")
			return
		}
	*/
	body := bytes.NewReader(Json)
	client := &http.Client{}
	resp, _ := http.NewRequest("POST", "http://192.168.30.123:8088/StoreCaptureImage", body)
	resp.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//发送
	respon, err := client.Do(resp)
	if err != nil {
		return
	}

	defer respon.Body.Close()
	data, err := ioutil.ReadAll(respon.Body)
	log.Println(string(data))

	var anykey string
	log.Println("please input any key for exit ...")
	fmt.Scanln(&anykey)

}
