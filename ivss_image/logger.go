package main

import "github.com/astaxie/beego/logs"

var G_FileLogs *logs.BeeLogger

func G_Logger_Init(logfile string) {
	//初始化log的缓存大小
	G_FileLogs = logs.NewLogger(1000)
	//EnableFuncCallDepth和SetLogFuncCallDepth, 用来设置函数的调用层级
	G_FileLogs.EnableFuncCallDepth(true)
	G_FileLogs.SetLogFuncCallDepth(2)
	G_FileLogs.Async()
	//设置日志文件的文件名、文件大小、文件的最大行数
	file := "{\"filename\":\"" + logfile + "\",\"maxsize\":10485760}"
	//FileLogs.SetLogger("file", "{\"filename\":\"log.log\"}")
	G_FileLogs.SetLogger(logs.AdapterFile, file)
}
