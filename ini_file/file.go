package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

import (
	"github.com/go-ini/ini" // 处理ini配置文件
)

var (
	g_strPathName  string
	g_iniFile      *ini.File   = new(ini.File)
	G_p_Ini_Option *ini_config = new(ini_config)
)

// config.ini 配置文件结构体
type ini_config struct {
}

const (
	g_str_ini_name string = "config.ini"
)

func main() {
	/*
		var dirctory string = ""
		var err error = nil
		//获取当前工作路径
		dirctory, err = filepath.Abs(filepath.Dir(os.Args[0]))
		fmt.Println(dirctory)
		fmt.Println(err)

		//拼接配置文件所在路径
		g_strPathName = dirctory + string(os.PathSeparator) + g_str_ini_name
		//加载配置文件
		g_iniFile, err = ini.Load(g_strPathName)
		fmt.Println(g_strPathName)
		fmt.Println(err)
		//将配置文件映射到G_p_Ini_Option结构体指针中
		err = g_iniFile.MapTo(G_p_Ini_Option)
	*/

	/*
		//解析命令行参数
		args := os.Args //获取当前exe程序所在路径
		fmt.Println(args)
		for _, v := range args {
			if v == "-reg" {
				fmt.Println("Hello!")
			}
		}
	*/

	/*  常量iota
	const (
		a = iota
		b
		c
	)
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
	*/

	//获取当前更新文件夹路径
	dirctory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	update_path := dirctory + string(os.PathSeparator) + "update"
	fmt.Println("文件全路径如下：")
	fmt.Println(update_path)

	var DirList []string
	var FileList []string
	//获取文件信息，返回一个FileInfo的接口
	_, err := os.Stat(update_path)
	if err == nil || os.IsExist(err) {
		// ReadDir 读取目录 update_path 中的所有目录和文件（不包括子目录，即下一层目录）
		// 返回读取到的文件的信息列表和读取过程中遇到的任何错误
		// 返回的文件列表是经过排序的
		files, err := ioutil.ReadDir(update_path)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("文件列表如下：")
		fmt.Println(files)

		for _, fi := range files {
			//获取更新文件的全路径
			fullpath := update_path + string(os.PathSeparator) + fi.Name()
			file, e := os.Stat(fullpath)
			if e == nil {
				//判断当前是文件还是目录
				if file.IsDir() {
					DirList = append(DirList, fi.Name())
				} else {
					FileList = append(FileList, fi.Name())
				}
			}
		}
		fmt.Println("文件夹名：")
		fmt.Println(DirList)
		fmt.Println("文件名：")
		fmt.Println(FileList)
	} else {
		log.Println("Update 目录不存在")
	}

	//moveDir(DirList, update_path, g_str_currentpath)
	//moveFile(FileList, update_path, g_str_currentpath)

}
