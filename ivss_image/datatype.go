package main

import (
	"protocol"
)

//存储底库照片结构体
type IStoreImageInfo struct {
	// 模板ID
	TemplateID int64 `json:"templateId"`
	// 分库ID
	LogicdbID int32 `json:"logicdbId"`
	// 照片
	Image protocol.IImage `json:"Image"`
}

//获取图片信息
type IFetchImageInfo struct {
	// 模板ID
	TemplateID int64 `json:"templateId"`
	// 分库ID
	LogicdbID int32 `json:"logicdbId"`
}

//批量输入获取图片信息
type IMultiImageInfo struct {
	FetchImageInfo []IFetchImageInfo `json:"FetchImageInfo"`
}

//批量获取图片返回信息
type IMultiImageResp struct {
	Answs  protocol.IAnswerResult `json:"answs"`
	Images []protocol.IImage      `json:"images"`
}

//采集存储结构体
type ICaptureInfo struct {
	// 分库ID
	LogID int32 `json:"logicdbId"`
	//时间
	HaveTime int64 `json:"HaveTime"`
}

//批量输入采集信息
type IMultiCaptureInfo struct {
	MultiCaptureInfo []ICaptureInfo `json:"CaptureInfo"`
}

type ICaptureImageInfo struct {
	//采集信息
	CaptureInfo ICaptureInfo `json:"CaptureInfo"`
	//照片
	Image protocol.IImage `json:"Image"`
}

type IDeleteDate struct {
	Days int64 `json:"days"`
}

//存储轨迹信息结构体
type IStoreTrialImage struct {
	JodID int64           `json:"jobID"`
	Image protocol.IImage `json:"image"`
}

//获取轨迹信息
type IFetchTrialImage struct {
	JobID int64 `json:"jobID"`
}
