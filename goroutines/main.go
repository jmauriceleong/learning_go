package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main(){
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go work()
	}
	wg.Wait()
}

func work(){
	time.Sleep(time.Second)

	var counter int

	for i := 0; i < 1e10; i++ {
		counter++
	}

	wg.Done()
}