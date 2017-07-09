package main

import (
	"fmt"
	"time"
	"utils"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("helloworld")
}

func main() {
	/*
		//日志处理
		utils.FileLogs.Info("this is a file log with info.")
		utils.FileLogs.Debug("this is a file log with debug.")
		utils.FileLogs.Alert("this is a file log with alert.")
		utils.FileLogs.Error("this is a file log with error.")
		utils.FileLogs.Trace("this is a file log with trace.")

		fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
			" |", "server", ": init file logger")
	*/

	beego.Router("/", &MainController{})
	beego.Run()

	var any string
	fmt.Scanln(&any)

}
