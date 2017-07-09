/* 删除文件函数
   os.Remove(file string) error

   file  文件名
   error 如果失败则返回错误信息

*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	//"strconv"
	"strings"
)

func main() {
	/*
		//获取当前工作路径
		dirctory, err := filepath.Abs(filepath.Dir(os.Args[0]))
		dirctory = dirctory + "\\" + strconv.Itoa(123)
		fmt.Println(dirctory)
	*/
	/*
		//删除文件
		file := "E:\\Work_Project_监控\\3601170113001动态识别轨迹跟踪系统\\03_代码\\GO\\trunk\\src\\ivss_image\\1" //源文件路径
		err = os.Remove(file)
		if err != nil {
			//如果删除失败则输出 file remove Error!
			fmt.Println("file remove Error!")
			//输出错误详细信息
			fmt.Printf("%s", err)
		} else {
			//如果删除成功则输出 file remove OK!
			fmt.Print("file remove OK!")
		}
	*/

	//遍历某个文件夹
	fileNames := make([]string, 0)
	dirNames := make([]string, 0)
	walkDir := "E:\\Work_Project_监控\\3601170113001动态识别轨迹跟踪系统\\03_代码\\GO\\trunk\\src\\ivss_image\\20170516"
	//遍历文件夹并把文件或文件夹名称加入相应的slice
	err := filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirNames = append(dirNames, path)
		} else {
			fileNames = append(fileNames, path)
		}
		return err
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(dirNames)
	fmt.Println(fileNames)

	//把所有文件名称连接成一个字符串
	fileNamesAll := strings.Join(fileNames, "")
	for i := len(dirNames) - 1; i >= 0; i-- {
		//文件夹名称不存在文件名称字符串内说明是个空文件夹
		if !strings.Contains(fileNamesAll, dirNames[i]) {
			fmt.Printf("%s is empty\n", dirNames[i])
			err := os.Remove(dirNames[i])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
