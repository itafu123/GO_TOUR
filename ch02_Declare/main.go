package main

import "fmt"

func main(){
	var i int = 10

	// 可不用定義型態，go自行判斷
	var l = 10
	
	// 可批量定義
	var (
		j int
		k int
	)	

	// 字串
	var (
		s1 string = "HELLO"
		s2 string = "世界"
	)
	
	// 簡短申明
	v := 1
	
	// 指針
	pi := &i
	
	// 常量
	const name = "常量"
	
	// iota
	const (
		one = iota+1
		two
		three
		four
	)


	fmt.Println(i)
	fmt.Println(l)
	fmt.Println(j+k)
	fmt.Println(s1+s2)
	fmt.Println(v)
	fmt.Println(*pi)
	fmt.Println(two)
	fmt.Println(three)
}