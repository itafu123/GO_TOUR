package main

import(
	"fmt"
	"sync"
	"time"
)

var(
	sum int
	mutex sync.RWMutex
)

func main(){
	run()
	doOnce()
	race()
}

func run(){
	var wg sync.WaitGroup

	wg.Add(110)

	for i:=0; i<100; i++{
		go func() {
			defer wg.Done()
			go add(10)
		}()
	}

	for j:=0; j<10; j++{
		go func() {
			defer wg.Done()
			go fmt.Println("sum ", j, "= " , readSUM() )
		}()
	}

 	wg.Wait()
}

func add(val int){
	mutex.Lock()
	defer mutex.Unlock()
	sum += val
}

func readSUM() int{
	mutex.RLock()
	defer mutex.RUnlock()
	b:=sum
	return b
}

func doOnce(){
	var once sync.Once
	onceBody := func(){
		fmt.Println("Only once")
	}

	done := make(chan bool)
	for i := 0; i < 10; i++{
		go func(){
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++{
		<-done
	} 
}

func race(){
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)

	for i:=0; i< 10; i++{
		go func(num int){
			defer wg.Done()

			fmt.Println(num, "號已就位")

			cond.L.Lock()
			cond.Wait()

			fmt.Println(num, "號開始跑")
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(2 * time.Second)

	go func(){
		defer wg.Done()
		fmt.Println("比賽開始")
		cond.Broadcast()
	}()

	wg.Wait()
}