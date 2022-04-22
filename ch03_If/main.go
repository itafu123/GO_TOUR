package main

import "fmt"

func main(){
	// 不用括號
	// 一定要大括號
	// 可以在if裡宣告
	if i := 6; i>10{
		fmt.Println("i>10")
	} else {
		fmt.Println("i<10")
	}

	switch ii := 6;{
	case ii > 10:
		fmt.Println("ii>10")
	case ii>5&&ii<=10:
		fmt.Println("ii>5&&ii<=10")
	default:
		fmt.Println("ii>10")
	}

	switch j := 1; j{
	case 1:
		// 跑下個case
		fallthrough
	case 2:
		fmt.Println("2")
	}

	// for
	sum:=0
	for i:=1; i<=100; i++{
		sum+=i
	}

	sum2 := 0
	i := 1
	for i <= 100{
		sum2 += 1
		i++
	}
}