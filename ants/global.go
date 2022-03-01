package main

import (
	"go.uber.org/zap"
	"sync"
)

const (
	DataSize    = 10000
	DataPerTask = 100
)

var logger, _ = zap.NewProduction()

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) DoTask() {
	for _, num := range t.nums {
		t.sum += num
	}

	t.wg.Done()

}
