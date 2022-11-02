package main

import "fmt"

func main(){
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}