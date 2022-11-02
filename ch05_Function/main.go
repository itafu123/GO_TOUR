package main

import "fmt"
import "errors"

type Age uint

// 等同construct
func init() {
	fmt.Println("will be the first")
}

func sum(a, b int) int {
	return a + b
}

// 多值返回
func sumM(a, b int)(int, error){
	if a < 0{
		return 0, errors.New("錯誤!")
	}
	return a+b, nil
}

// 命名返回參數
func sumN(a, b int)(sum int, err error){
	if a < 0{
		return 0, errors.New("錯誤!")
	}

	sum = a+b
	err = nil
	return
}

// 可變參數
func sumC(params...int) int {
	sum:=0
	for _,v:=range params{
		sum+=v
	}
	return sum
}

// 包級函數
func(age Age) String(){
	fmt.Println("age is ", age)
}

func(age *Age) Modify(){
	*age = Age(30)
}

func main(){
	result, error := sumM(-1, 1)
	
	// 下劃線丟棄不要得返回值
	result2, _ := sumM(-1, 1)
	
	// 匿名函數
	sum2 := func(a, b int) int {
		return a+b
	}
	
	age := Age(25)
	age.String()
	age.Modify()
	age.String()
	
	fmt.Println(result)
	fmt.Println(error)
	fmt.Println(result2)
	fmt.Println(sum2(1,3))
}
