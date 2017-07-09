package main

import (
	"fmt"
	//导入beego文件包（导入包时,beego包中会初始化一个BeeAPP的应用，初始化一些参数）
	"github.com/astaxie/beego/logs" //beego日志文件包的路径
)

var FileLogs *logs.BeeLogger

func test(file string) {
	//初始化log的缓存大小
	FileLogs := logs.NewLogger(10000)
	//EnableFuncCallDepth和SetLogFuncCallDepth, 用来设置函数的调用层级
	FileLogs.EnableFuncCallDepth(true)
	FileLogs.SetLogFuncCallDepth(2)
	FileLogs.Async()

	logfile := "{\"filename\":\"" + file + "\",\"maxsize\":10240000}"
	fmt.Println(logfile)

	//设置日志文件的文件名、文件大小、文件的最大行数
	//FileLogs.SetLogger(logs.AdapterFile, "{\"filename\":\"test.log\",\"maxsize\":100000,\"maxlines\":100}")
	FileLogs.SetLogger(logs.AdapterFile, "{\"filename\":\"test_log.log\",\"maxsize\":100000}")
	for i := 0; i < 100; i++ {
		FileLogs.Info("xxxxx")
		FileLogs.Error("%v: %v", 2, "xxxxx")
		FileLogs.Trace("%v: %v", 3, "xxxxx")
		FileLogs.Warn("%v: %v", 4, "xxxxx")
	}
}

func main() {
	test("test.log")

}
