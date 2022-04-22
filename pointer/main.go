package main

import "fmt"

func main(){
	// 指針
	var a int = 10
	var ip *int

	// & -> 位址
	// * -> 值
	ip = &a
	
	fmt.Println("a:", a)
	fmt.Println("&a", &a)
	fmt.Println("ip", ip)
	fmt.Println("*ip", *ip)
}