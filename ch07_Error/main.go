package main

import "fmt"
import "errors"

func main(){
	sum, err := add(-1, 2)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}

	sum2, err2 := add2(-1, 2)
	if cm, ok:= err2.(*commonError); ok{
		fmt.Println("錯誤代碼:" , cm.errorCode , "; 錯誤訊息:" , cm , ok)
	} else {
		fmt.Println(sum2)
	}

	// newErr := MyError{err, "第二層錯誤"}
	// w := fmt.Errorf("wrap一個錯誤:", newErr)
	// fmt.Println(w)
	// fmt.Println(errors.Unwrap(w))
} 

// 基本
func add(a, b int)(int, error){
	if a < 0 || b < 0 {
		return 0, errors.New("ab不能為負數")
	} else {
		return a+b, nil
	}
}

// 定義自己的錯誤格式
type commonError struct{
	errorCode int
	errorMsg string
}

func (ce *commonError) Error() string{
	return ce.errorCode
}

func add2(a, b int)(int, error){
	if a < 0 || b < 0 {
		return 0, &commonError{
			errorCode: 1,
			errorMsg: "ab不能為負數"}
	} else {
		return a+b, nil
	}
}

// // Error Wrapping
// type MyError struct{
// 	err error
// 	msg string
// }

// func (e *MyError) Error() string{
// 	return e.err.Error() + e.msg
// }

// DEFER
func ReadFile(fileName string)([]byte, error){
	f, err := os.Open(fileName)
	if err != nil{
		return nil, err
	}

	defer f.Close() // 會比「return readAll(f, n)」晚執行

	return readAll(f, n)
}

// PAINC
func connectMySQL(ip){
	if ip = "" {
		painc("ip不能為空")
	}
}