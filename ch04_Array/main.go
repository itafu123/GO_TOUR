package main

import "fmt"

func main(){
	// Array ==========
	array1 := [3]string{"a", "b", "c"}
	array2 := [...]string{"a", "b", "c", "d"}
	array3 := [5]string{1:"b",3:"c"}

	// 多維陣列
	array4 := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	array4 = append(array4, row1)
	array4 = append(array4, row2)
	
	array5 := [2][2]int{
		{0, 1},
		{2, 3},
	}

	fmt.Println(array4[1][1])
	fmt.Println(array5[1][1])


	// for迴圈
	// for range
	for i,v := range array3{
		fmt.Printf("索引:%d,值:%s\n", i, v)
	}

	// 下劃線捨棄索引
	for _,v2 := range array2{
		fmt.Println(v2)
	}


	// Slice ==========
	// array[start,end] : 取start(包含) 到 end(不包含)
	slice1 := array1[1:3]
	
	slice2 := array1[:3] // [0:3]
	
	slice3 := array1[1:] // [1:3]
	
	slice4 := array1[:] // [0:3]

	slice7 :=[] int {1,2,3 } 
	
	// Make
	// 製造Slice: make(array,長度,容量)
	slice5 := make([]string, 4, 8)
	slice6 := append(slice5, "Z")

	// slice後的資料還是指向原本的數組
	slice1[0] = "should_be_b"


	// Map ==========
	nameAgeMap := make(map[string] int)
	nameAgeMap["小明"] = 20

	nameAgeMap2 := map[string] int {"小花":10}

	// 獲取直
	age, ok := nameAgeMap2["小花"]
	
	// delete
	delete(nameAgeMap, "小明")
	

	// String ==========
	s := "Hello世界"
	bs := []byte(s)
	
	fmt.Println(slice1)
	fmt.Println(array1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println("make([]string,4,8):", slice5)
	fmt.Println("slice5:", slice5)
	fmt.Println("append(slice5,'Z'):", slice6)
	fmt.Println(age, ok)
	fmt.Println(bs)
	fmt.Println(s[0], s[1], s[10])
}