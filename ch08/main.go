package main

import(
	"fmt"
	"time"
)

func main(){
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)

	go func(){
		firstCh <- download("firstCh")
	}()
	go func(){
		secondCh <- download("secondCh")
	}()
	go func(){
		threeCh <- download("threeCh")
	}()

	select{
	case filePath := <- firstCh:
		fmt.Println(filePath)
	case filePath := <- secondCh:
		fmt.Println(filePath)
	case filePath := <- threeCh:
		fmt.Println(filePath)

	}
}

func download(name string) string{
	time.Sleep(time.Second)
	return name + ":filepath"
}