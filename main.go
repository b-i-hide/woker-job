package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Dispatcher struct {
	queue chan int32
	wg    sync.WaitGroup
}

var wg *sync.WaitGroup
var counter int32 = 0

const (
	maxWorkers = 3
	maxQueues  = 100
)

func main() {
	d := NewDispatcher()

	for v := 0; v < maxWorkers; v++ {
		go doSomething(wg, d)
	}

	for i := 0; i < 1000; i++ {
		d.wg.Add(maxWorkers)
		counter = atomic.AddInt32(&counter, 1)
		d.queue <- counter

	}
	close(d.queue)
	d.wg.Wait()
	log.Println("end main")
}

// 時間がかかるダミー処理
func doSomething(wg *sync.WaitGroup, d *Dispatcher) {
	counter := <-d.queue
	log.Printf("start doSomething: %d\n", counter)
	time.Sleep(5 * time.Second)
	log.Printf("end doSomething: %d\n", counter)
	d.wg.Done()
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		queue: make(chan int32, maxQueues),
	}
	return d
}
