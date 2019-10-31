package main

import (
	"fmt"
	"time"
)

var ch1 =make(chan int,1)

func main() {
	// 启用20个goroutine同时求1~20以内各个数的阶乘
	for i := 1; i <= 20; i++ {
		go factorial(i)

	}
	time.Sleep(1 * time.Second)
}
// 求n的阶乘，并将结果写进存入通道ch1
func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch1 <- res
	b :=<-ch1
	fmt.Println(b)
}
