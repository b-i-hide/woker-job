package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var wg *sync.WaitGroup
var counter int32 = 0

const (
	maxJob       = 20
	maxGoRoutine = 5
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int, maxJob)
	defer close(ch)

	for i := 0; i < maxGoRoutine; i++ {
		go func() {
			for range ch {
				doSomething()
				wg.Done()
			}
		}()
	}

	for i := 0; i < 200; i++ {
		wg.Add(1)
		ch <- i
	}

	wg.Wait()
	log.Println("end main")
}

// 時間がかかるダミー処理
func doSomething() {
	counter := atomic.AddInt32(&counter, 1)
	log.Printf("start doSomething: %d\n", counter)
	time.Sleep(1 * time.Second)
	log.Printf("end doSomething: %d\n", counter)
}
