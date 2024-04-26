package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var cond = sync.NewCond(&mutex)

var queue []int

func producer() {
	i := 0
	for {
		cond.L.Lock()
		queue = append(queue, i)
		i++
		cond.L.Unlock()

		cond.Signal()
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(consumerName string) {
	for {
		cond.L.Lock()
		for len(queue) == 0 {
			cond.Wait()
		}

		fmt.Println(consumerName, queue[0])
		queue = queue[1:]
		cond.L.Unlock()
	}
}

func main() {
	// 开启一个 producer
	go producer()

	// 开启两个 consumer
	go consumer("consumer-0")
	go consumer("consumer-1")
	go consumer("consumer-2")
	go consumer("consumer-3")
	go consumer("consumer-4")
	go consumer("consumer-5")
	go consumer("consumer-6")
	go consumer("consumer-7")
	go consumer("consumer-8")
	go consumer("consumer-9")

	time.Sleep(1 * time.Minute)
}
