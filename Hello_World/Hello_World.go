package main

import (
	"fmt"
	"log"
)

/*
#include "Add.h"
#include <stdlib.h>
*/
import "C"

func loop() {
	for i := 1; i < 10; i++ {
		log.Println(i)
	}
}

//线程退出
func waitForQuit(quit chan bool) {
	// NewReader将os.Stdin封装成一个拥有size大小缓存的bufio.Reader对象
	reader := bufio.NewReader(os.Stdin)
	for {
		//用于读取一行数据，不包括行尾标记（\n 或 \r\n）
		data, _, _ := reader.ReadLine()
		command := string(data)
		//判断是否为“q”
		if command == "q" {
			g_stop = true
			break //退出线程
		}
	}
	//向信道存放数据
	quit <- true
}

func main() {

	/* fmt输出日志*/
	i := 10000
	j := 86
	k := i + j
	fmt.Println(k)
	fmt.Println("Hello World,您好，世界！")

	/*log输出日志*/
	log.Println("郭彩凤")

	//调用c函数
	a := C.sub(1, 2)
	log.Println(a)

	//线程
	var complete chan int = make(chan int)
	go loop()
	complete <- 0
	<-complete

	//创建一个退出信道
	quit := make(chan bool)
	defer close(quit)
	//退出线程
	go waitForQuit(quit)
	<-quit             //取quit消息
	g_waitgroup.Wait() //判断sync.WaitGroup中计数字段，如果为0，则返回
	//如果计数大于0，Wait()可以用来阻塞直到队列中的所有任务都完成时才解除阻塞

	//释放资源（初始化sdk在建模包的init函数中实现）
	//ExtractFeature.Go_Uninit()

	var anykey string
	log.Println("please input any key for exit ...")
	/*从标准输入中读取数据，并将数据用空白分割并解析后存入anykey提供的变量中
	（换行符会被当作空白处理），变量必须以指针传入,遇到换行符就停止扫描*/
	fmt.Scanln(&anykey)
}
