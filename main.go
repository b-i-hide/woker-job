package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	defer wg.Wait()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go doSomething(wg)
	}
	log.Println("end main")
}

var counter int32 = 0

// 時間がかかるダミー処理
func doSomething(wg *sync.WaitGroup) {
	counter := atomic.AddInt32(&counter, 1)
	log.Printf("start doSomething: %d\n", counter)
	time.Sleep(5 * time.Second)
	log.Printf("end doSomething: %d\n", counter)
	wg.Done()
}
