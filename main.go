package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var wg *sync.WaitGroup
var counter int32 = 0

func main() {
	wg = new(sync.WaitGroup)

	ch := make(chan int32, 10)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		counter = atomic.AddInt32(&counter, 1)
		ch <- counter
		go doSomething(wg, ch)
	}
	close(ch)
	wg.Wait()
	log.Println("end main")
}

// 時間がかかるダミー処理
func doSomething(wg *sync.WaitGroup, ch chan int32) {
	counter := <-ch
	log.Printf("start doSomething: %d\n", counter)
	time.Sleep(5 * time.Second)
	log.Printf("end doSomething: %d\n", counter)
	wg.Done()
}
