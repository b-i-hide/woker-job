package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var counter int32 = 0

type Dispatcher struct {
	queue chan int
	wg    sync.WaitGroup
}

const (
	maxWorkers = 20
	maxQueues  = 100
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	d := NewDispatcher()
	d.Start(ctx)
	d.SendJob()

	close(d.queue)
	d.wg.Wait()
	log.Println("end main")
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		queue: make(chan int, maxQueues),
	}
	return d
}

func (d *Dispatcher) Start(ctx context.Context) {
	d.wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go func() {
			defer d.wg.Done()
			for range d.queue {
				doSomething(ctx)
			}
		}()
	}
}

func (d *Dispatcher) SendJob() {
	for i := 0; i < 1000; i++ {
		d.queue <- i
	}
}

// 時間がかかるダミー処理
func doSomething(ctx context.Context) {
	counter := atomic.AddInt32(&counter, 1)
	log.Printf("start doSomething: %d\n", counter)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	log.Printf("end doSomething: %d\n", counter)
	select {
	case <-ctx.Done():
		log.Printf("cancel work func: %d", counter)
		return
	default:
		return
	}
}
