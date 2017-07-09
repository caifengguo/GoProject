package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// 根据时间设置随机数种子
	rand.Seed(int64(time.Now().Nanosecond()))
	// 获取指定范围内的随机数
	for i := 1; i < 2; i++ {
		fmt.Println(rand.Int63())
		fmt.Println(strconv.FormatInt(rand.Int63(), 10))
	}
}
