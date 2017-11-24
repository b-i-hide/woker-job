package main

import (
	"log"
	"sync/atomic"
	"time"
)

func main() {

	for i := 0; i < 1000; i++ {
		go doSomething()
	}

	time.Sleep(5 * time.Second)
	log.Println("end main")
}

var counter int32 = 0

// 時間がかかるダミー処理
func doSomething() {
	counter := atomic.AddInt32(&counter, 1)
	log.Printf("start doSomething: %d\n", counter)
	time.Sleep(4 * time.Second)
	log.Printf("end doSomething: %d\n", counter)
}
