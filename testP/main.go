package main

import (
	"fmt"
	"math/rand"
	"mywork/learning/testP/random"
	"time"
)

// 测试函数
func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("测试数据", random.RandStr(10))
}
