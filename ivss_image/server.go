package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	//"strings"
	"time"
)

func ReadRequestBody(req *http.Request) (int, []byte) {
	//判断请求方式
	if req.Method != "POST" {
		return http.StatusMethodNotAllowed, nil
	}
	//获取消息包内容
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusBadRequest, nil
	}
	return http.StatusOK, body
}

//存储底库图片
func StoreImageHandler(w http.ResponseWriter, req *http.Request) {
	status, body := ReadRequestBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	AnswerResp := StoreImage(body)
	w.Write(AnswerResp)
}

//获取底库图片
func FetchImageHandler(w http.ResponseWriter, req *http.Request) {
	status, body := ReadRequestBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	ImageResp := FetchImage(body)
	w.Write(ImageResp)
}

//通过模板id删除底库图片
func DeleteImageByTemplateIdHandler(w http.ResponseWriter, req *http.Request) {
	status, body := ReadRequestBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	AnswerResp := DeleteImageByTemplateId(body)
	w.Write(AnswerResp)
}

//通过逻辑分库id删除底库图片
func DeleteImageByLogicIdHandler(w http.ResponseWriter, req *http.Request) {
	status, body := ReadRequestBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	AnswerResp := DeleteImageByLogicId(body)
	w.Write(AnswerResp)
}

//存储采集图片
func StoreCaptureImageHandler(w http.ResponseWriter, req *http.Request) {
	status, body := ReadRequestBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	AnswerResp := StoreCaptureImage(body)
	w.Write(AnswerResp)
}

//获取采集图片
func FetchCaptureImageHandler(w http.ResponseWriter, req *http.Request) {
	status, body := ReadRequestBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	ImageResp := FetchCaptureImage(body)
	w.Write(ImageResp)
}

//删除过期图片
func DeleteTimeoutImageHandler(w http.ResponseWriter, req *http.Request) {
	status, body := ReadRequestBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	AnswerResp := DeleteTimeoutImage(body)
	w.Write(AnswerResp)
}

//存储轨迹任务图片
func StoreTrialImageHandler(w http.ResponseWriter, req *http.Request) {
	status, body := ReadRequestBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	AnswerResp := StoreTrialImage(body)
	w.Write(AnswerResp)
}

//获取轨迹任务图片
func FetchTrialImageHandler(w http.ResponseWriter, req *http.Request) {
	status, body := ReadRequestBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	AnswerResp := FetchTrialImage(body)
	w.Write(AnswerResp)
}

func main() {
	// Parse input param
	// Get server name
	//serverName := strings.TrimSpace(os.Args[1])
	serverName := "ivss_image_server"

	//Init Log
	G_Logger_Init("log/" + serverName + ".log")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), //2006-01-02 15:04:05 固定时间
		" |", serverName, ": init file logger")
	G_FileLogs.Info("G_Logger_Init success")

	// Init config_xml
	if G_Option_Init("conf/"+serverName+".xml") == false {
		G_FileLogs.Error("G_Option_Init failed")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
			" |", serverName, ": read config xml failed")
		return
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": read config xml")
	G_FileLogs.Info("G_Option_Init success")

	//底库图片处理接口
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": Serverhost/StoreImage")
	http.HandleFunc("/StoreImage", StoreImageHandler)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": Serverhost/FetchImage")
	http.HandleFunc("/FetchImage", FetchImageHandler)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": Serverhost/DeleteImageByTemplateId")
	http.HandleFunc("/DeleteImageByTemplateId", DeleteImageByTemplateIdHandler)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": Serverhost/DeleteImageByLogicId")
	http.HandleFunc("/DeleteImageByLogicId", DeleteImageByLogicIdHandler)

	//采集图片处理接口
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": Serverhost/StoreCaptureImage")
	http.HandleFunc("/StoreCaptureImage", StoreCaptureImageHandler)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": Serverhost/FetchCaptureImage")
	http.HandleFunc("/FetchCaptureImage", FetchCaptureImageHandler)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": Serverhost/DeleteTimeoutImage")
	http.HandleFunc("/DeleteTimeoutImage", DeleteTimeoutImageHandler)

	//轨迹任务处理接口
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": Serverhost/FetchCaptureImage")
	http.HandleFunc("/StoreTrialImage", StoreTrialImageHandler)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": Serverhost/DeleteTimeoutImage")
	http.HandleFunc("/FetchTrialImage", FetchTrialImageHandler)

	//服务器要监听的主机地址和端口号
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		" |", serverName, ": ListenHttp start :", G_Serverconfig.ListenUrl)
	err := http.ListenAndServe(G_Serverconfig.ListenUrl, nil)
	if err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
			" |", serverName, ": Listen Http error")
		G_FileLogs.Error("Listen Http error:", err.Error())
	}
}
