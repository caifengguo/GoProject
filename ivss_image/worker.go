package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)
import (
	"protocol"
)

//创建图片存储路径
func CreateStoreImagePath(LogicdbID int32) error {
	//获取当前工作路径
	dirctory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		G_FileLogs.Error("获取当前工作路径失败!")
		return err
	}
	//根据逻辑分库id创建对应的文件夹(0~99)
	for i := 0; i < 100; i++ {
		var Imagepath string
		Imagepath = dirctory + "\\" + strconv.Itoa(int(LogicdbID))
		Imagepath = Imagepath + "\\" + strconv.Itoa(i)
		//判断图片文件夹是否存在
		_, err = os.Stat(Imagepath)
		if os.IsExist(err) {
			return nil
		}
		//不存在，则创建
		err = os.MkdirAll(Imagepath, 0666)
		if err != nil {
			G_FileLogs.Error("创建图片保存路径失败!")
			return err
		}
	}
	return nil
}

//存储底库照片
func StoreImage(body []byte) []byte {
	//初始化
	r := &protocol.IAnswerResult{0, "success"}
	//1、解析json
	var ImageInfo IStoreImageInfo
	errMsg := json.Unmarshal(body, &ImageInfo)
	if errMsg != nil {
		G_FileLogs.Error("StoreImage Unmarshal json err:", errMsg)
		r.Code = 1
		r.Desc = "StoreImage Unmarshal json fail"
		b, _ := json.Marshal(r)
		return b
	}
	//2、根据逻辑分库id创建(0~99)文件夹
	dirctory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dirctory = dirctory + "\\" + strconv.Itoa(int(ImageInfo.LogicdbID))
	//判断当前逻辑分库对应的文件夹是否存在
	_, err = os.Stat(dirctory)
	if os.IsNotExist(err) {
		G_FileLogs.Info("logic file no exit!")
		//不存在，则根据逻辑分库创建文件夹
		err = CreateStoreImagePath(ImageInfo.LogicdbID)
		if err != nil {
			r.Code = 1
			r.Desc = "create file faild"
			b, _ := json.Marshal(r)
			return b
		}
	}
	//3、根据TemplateID确定图片文件的存放路径、图片名称
	ModeId := ImageInfo.TemplateID % 100 //针对模板id求余，确定存放在的文件夹
	dirctory = dirctory + "\\" + strconv.FormatInt(ModeId, 10)
	dirctory = dirctory + "\\" + strconv.FormatInt(ImageInfo.TemplateID, 10)
	dirctory += ".jpg"

	//4、保存图片信息
	err = ioutil.WriteFile(dirctory, ImageInfo.Image.Raw, 0666)
	if err != nil {
		G_FileLogs.Error("StoreImage save image file fail!")
		r.Code = 1
		r.Desc = "StoreImage save image fail"
		b, _ := json.Marshal(r)
		return b
	}
	//构造Json
	b, _ := json.Marshal(r)
	return b
}

//获取底库照片
func FetchImage(body []byte) []byte {
	//初始化返回值
	var ImageResp IMultiImageResp
	//1、解析json
	var ImageInfo IMultiImageInfo
	errMsg := json.Unmarshal(body, &ImageInfo)
	if errMsg != nil {
		G_FileLogs.Error("FetchImage Unmarshal json err:", errMsg)
		ImageResp.Answs.Code = 1
		ImageResp.Answs.Desc = "FetchImage Unmarshal json fail"
		b, _ := json.Marshal(ImageResp)
		return b
	}
	//2、根据图片路径获取图片信息
	dirctory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	for i, _ := range ImageInfo.FetchImageInfo {
		var ImagePath string
		ImagePath = dirctory + "\\" + strconv.Itoa(int(ImageInfo.FetchImageInfo[i].LogicdbID))
		//根据模板id与逻辑分库id获取图片信息
		ModeId := ImageInfo.FetchImageInfo[i].TemplateID % 100 //针对模板id求余，确定存放在的文件夹
		ImagePath = ImagePath + "\\" + strconv.FormatInt(ModeId, 10)
		ImagePath = ImagePath + "\\" + strconv.FormatInt(ImageInfo.FetchImageInfo[i].TemplateID, 10)
		ImagePath += ".jpg"
		G_FileLogs.Info("get imagefile path info:", ImagePath)
		//读取图片信息
		imagebuffer, errcode := ioutil.ReadFile(ImagePath)
		if errcode != nil {
			ImageResp.Answs = protocol.IAnswerResult{1, "read imageinfo fail"}
			continue
		}
		ImageResp.Answs = protocol.IAnswerResult{Code: 0,
			Desc: "read imageinfo success"}
		image := protocol.IImage{Channel: 1,
			Width:  0,
			Height: 0,
			Raw:    imagebuffer}
		ImageResp.Images = append(ImageResp.Images, image)
	}

	//3.构造Json对象
	JsonAnswer, _ := json.Marshal(ImageResp)
	return JsonAnswer
}

//通过模板id删除底库图片
func DeleteImageByTemplateId(body []byte) []byte {
	r := &protocol.IAnswerResult{0, "success"}

	//1、解析json
	var FetchImageInfo IFetchImageInfo
	errMsg := json.Unmarshal(body, &FetchImageInfo)
	if errMsg != nil {
		G_FileLogs.Error("DeleteImageByTemplateId Unmarshal json err:", errMsg)
		r.Code = 1
		r.Desc = "DeleteImageByTemplateId Unmarshal json fail"
		b, _ := json.Marshal(r)
		return b
	}

	//2、构造图片路径名
	dirctory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dirctory = dirctory + "\\" + strconv.Itoa(int(FetchImageInfo.LogicdbID))
	//根据模板id与逻辑分库id获取图片信息
	ModeId := FetchImageInfo.TemplateID % 100 //针对模板id求余，确定存放在的文件夹
	dirctory = dirctory + "\\" + strconv.FormatInt(ModeId, 10)
	dirctory = dirctory + "\\" + strconv.FormatInt(FetchImageInfo.TemplateID, 10)
	dirctory += ".jpg"

	//3、删除图片文件
	err := os.Remove(dirctory)
	if err != nil {
		G_FileLogs.Error("Remove file fail")
		r.Code = 1
		r.Desc = "Remove file fail"
		b, _ := json.Marshal(r)
		return b
	}

	//4、构造Json对象
	JsonAnswer, _ := json.Marshal(r)
	return JsonAnswer
}

//通过逻辑分库id删除底库图片(默认模板id为0)
func DeleteImageByLogicId(body []byte) []byte {
	r := &protocol.IAnswerResult{0, "success"}

	//1、解析json
	var FetchImageInfo IFetchImageInfo
	errMsg := json.Unmarshal(body, &FetchImageInfo)
	if errMsg != nil {
		G_FileLogs.Error("DeleteImageByLogicId Unmarshal json err:", errMsg)
		r.Code = 1
		r.Desc = "DeleteImageByLogicId Unmarshal json fail"
		b, _ := json.Marshal(r)
		return b
	}
	//2、构造删除逻辑分库
	dirctory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dirctory = dirctory + "\\" + strconv.Itoa(int(FetchImageInfo.LogicdbID))
	//删除文件夹
	err := DeleteFileDir(dirctory)
	if err != nil {
		G_FileLogs.Error("Remove file fail")
		r.Code = 1
		r.Desc = "Remove file fail"
		b, _ := json.Marshal(r)
		return b
	}
	//4、构造Json对象
	JsonAnswer, _ := json.Marshal(r)
	return JsonAnswer
}

//存储采集照片
func StoreCaptureImage(body []byte) []byte {
	r := &protocol.IAnswerResult{0, "success"}

	//1、解析json
	var CaptureImageInfo ICaptureImageInfo
	errMsg := json.Unmarshal(body, &CaptureImageInfo)
	if errMsg != nil {
		G_FileLogs.Error("StoreCaptureImage Unmarshal json err:", errMsg)
		r.Code = 1
		r.Desc = "StoreCaptureImage Unmarshal json fail"
		b, _ := json.Marshal(r)
		return b
	}
	//2、根据时间创建图片文件夹
	havatime := time.Unix(CaptureImageInfo.CaptureInfo.HaveTime, 0).Format("20060102")
	dirctory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dirctory = dirctory + "\\" + havatime
	//判断当前逻辑分库对应的文件夹是否存在
	_, err = os.Stat(dirctory)
	if os.IsNotExist(err) {
		G_FileLogs.Info("logic file no exit!")
		//不存在，则根据逻辑分库创建文件夹
		fileId, _ := strconv.Atoi(havatime)
		err = CreateStoreImagePath(int32(fileId))
		if err != nil {
			r.Code = 1
			r.Desc = "create file faild"
			b, _ := json.Marshal(r)
			return b
		}
	}
	//3、根据LogID确定图片文件的存放路径、图片名称
	ModeId := CaptureImageInfo.CaptureInfo.LogID % 100 //针对模板id求余，确定存放在的文件夹
	dirctory = dirctory + "\\" + strconv.Itoa(int(ModeId))
	dirctory = dirctory + "\\" + strconv.Itoa(int(CaptureImageInfo.CaptureInfo.LogID))
	dirctory += ".jpg"

	//4、保存图片信息
	err = ioutil.WriteFile(dirctory, CaptureImageInfo.Image.Raw, 0666)
	if err != nil {
		G_FileLogs.Error("StoreCaptureImage save image file fail!")
		r.Code = 1
		r.Desc = "StoreCaptureImage save image fail"
		b, _ := json.Marshal(r)
		return b
	}

	//4、构造Json对象
	JsonAnswer, _ := json.Marshal(r)
	return JsonAnswer
}

//获取采集照片
func FetchCaptureImage(body []byte) []byte {
	//初始化返回值
	var ImageResp IMultiImageResp
	//1、解析json
	var CaptureInfo IMultiCaptureInfo
	errMsg := json.Unmarshal(body, &CaptureInfo)
	if errMsg != nil {
		G_FileLogs.Error("FetchCaptureImage Unmarshal json err:", errMsg)
		ImageResp.Answs.Code = 1
		ImageResp.Answs.Desc = "FetchCaptureImage Unmarshal json fail"
		b, _ := json.Marshal(ImageResp)
		return b
	}
	//2、根据图片路径获取图片信息
	dirctory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	for i, _ := range CaptureInfo.MultiCaptureInfo {
		var ImagePath string
		havatime := time.Unix(CaptureInfo.MultiCaptureInfo[i].HaveTime, 0).Format("20060102")
		ImagePath = dirctory + "\\" + havatime
		//根据时间与逻辑分库id获取图片信息
		ModeId := CaptureInfo.MultiCaptureInfo[i].LogID % 100 //针对模板id求余，确定存放在的文件夹
		ImagePath = ImagePath + "\\" + strconv.Itoa(int(ModeId))
		ImagePath = ImagePath + "\\" + strconv.Itoa(int(CaptureInfo.MultiCaptureInfo[i].LogID))
		ImagePath += ".jpg"
		G_FileLogs.Info("get imagefile path info:", ImagePath)

		//读取图片信息
		imagebuffer, errcode := ioutil.ReadFile(ImagePath)
		if errcode != nil {
			ImageResp.Answs = protocol.IAnswerResult{1, "read imageinfo fail"}
			continue
		}
		ImageResp.Answs = protocol.IAnswerResult{Code: 0,
			Desc: "读取图片信息成功!"}
		image := protocol.IImage{Channel: 1,
			Width:  0,
			Height: 0,
			Raw:    imagebuffer}
		ImageResp.Images = append(ImageResp.Images, image)
	}

	//3、构造Json对象
	JsonAnswer, _ := json.Marshal(ImageResp)
	return JsonAnswer
}

//删除过期照片
func DeleteTimeoutImage(body []byte) []byte {
	//初始化返回值
	r := &protocol.IAnswerResult{0, "success"}
	//1、解析json
	var DayData IDeleteDate
	errMsg := json.Unmarshal(body, &DayData)
	if errMsg != nil {
		G_FileLogs.Error("DeleteTimeoutImage Unmarshal json err:", errMsg)
		r.Code = 1
		r.Desc = "DeleteTimeoutImage Unmarshal json fail"
		b, _ := json.Marshal(r)
		return b
	}
	//2、删除照片文件
	dirctory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	daytime := time.Unix(DayData.Days, 0).Format("20060102")
	dirctory = dirctory + "\\" + daytime
	//删除文件夹
	err := DeleteFileDir(dirctory)
	if err != nil {
		G_FileLogs.Error("Remove file fail")
		r.Code = 1
		r.Desc = "Remove file fail"
		b, _ := json.Marshal(r)
		return b
	}
	//3、构造Json对象
	JsonAnswer, _ := json.Marshal(r)
	return JsonAnswer
}

//遍历删除目录下的所有文件
func DeleteFileDir(dirctory string) error {
	fileNames := make([]string, 0) //存放文件
	dirNames := make([]string, 0)  //存放空目录
	err := filepath.Walk(dirctory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirNames = append(dirNames, path)
		} else {
			fileNames = append(fileNames, path)
		}
		return err
	})
	if err != nil {
		return err
	}

	//删除所有文件夹下的文件
	for i := len(fileNames) - 1; i >= 0; i-- {
		err = os.Remove(fileNames[i])
		if err != nil {
			return err
		}
	}
	//删除所有空文件夹
	for j := len(dirNames) - 1; j >= 0; j-- {
		err = os.Remove(dirNames[j])
		if err != nil {
			return err
		}
	}
	return nil
}

//存储轨迹任务图片
func StoreTrialImage(body []byte) []byte {
	r := &protocol.IAnswerResult{0, "success"}

	//1、解析json
	var StoreTrialImage IStoreTrialImage
	errMsg := json.Unmarshal(body, &StoreTrialImage)
	if errMsg != nil {
		G_FileLogs.Error("StoreTrialImage Unmarshal json err:", errMsg)
		r.Code = 1
		r.Desc = "StoreTrialImage Unmarshal json fail"
		b, _ := json.Marshal(r)
		return b
	}
	//2、根据任务ID确定图片名称
	dirctory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dirctory = dirctory + "\\" + strconv.FormatInt(StoreTrialImage.JodID, 10)
	dirctory += ".jpg"

	//3、保存图片信息
	err = ioutil.WriteFile(dirctory, StoreTrialImage.Image.Raw, 0666)
	if err != nil {
		G_FileLogs.Error("StoreTrialImage save image file fail!")
		r.Code = 1
		r.Desc = "StoreTrialImage save image fail"
		b, _ := json.Marshal(r)
		return b
	}

	//4、构造Json对象
	JsonAnswer, _ := json.Marshal(r)
	return JsonAnswer
}

//获取轨迹任务图片
func FetchTrialImage(body []byte) []byte {
	//初始化返回值
	var ImageResp IMultiImageResp
	//1、解析json
	var TrialImageInfo IFetchTrialImage
	errMsg := json.Unmarshal(body, &TrialImageInfo)
	if errMsg != nil {
		G_FileLogs.Error("FetchTrialImage Unmarshal json err:", errMsg)
		ImageResp.Answs.Code = 1
		ImageResp.Answs.Desc = "FetchTrialImage Unmarshal json fail"
		b, _ := json.Marshal(ImageResp)
		return b
	}
	//2、根据图片路径获取图片信息
	dirctory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dirctory = dirctory + "\\" + strconv.FormatInt(TrialImageInfo.JobID, 10)
	dirctory += ".jpg"
	G_FileLogs.Info("get imagefile path info:", dirctory)

	//读取图片信息
	imagebuffer, errcode := ioutil.ReadFile(dirctory)
	if errcode != nil {
		ImageResp.Answs = protocol.IAnswerResult{1, "read imageinfo fail"}
		continue
	}
	ImageResp.Answs = protocol.IAnswerResult{Code: 0,
		Desc: "读取图片信息成功!"}
	ImageResp.Images = protocol.IImage{Channel: 1,
		Width:  0,
		Height: 0,
		Raw:    imagebuffer}

	//3、构造Json对象
	JsonAnswer, _ := json.Marshal(ImageResp)
	return JsonAnswer
}
