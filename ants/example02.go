package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
)

func example02Func(data interface{}) {
	task := data.(*Task)
	task.DoTask()
	fmt.Printf("task:%d sum:%d\n", task.index, task.sum)
}

func example02() {
	p, err := ants.NewPoolWithFunc(10, example02Func)
	if err != nil {
		return
	}
	defer p.Release()

	nums := make([]int, DataSize, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(100)
	}

	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	tasks := make([]*Task, 0, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		task := &Task{
			index: i + 1,
			nums:  nums[i*DataPerTask : (i+1)*DataPerTask],
			wg:    &wg,
		}

		tasks = append(tasks, task)
		err := p.Invoke(task)
		if err != nil {
			return
		}
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	var sum int
	for _, task := range tasks {
		sum += task.sum
	}

	var expect int
	for _, num := range nums {
		expect += num
	}

	fmt.Printf("finish all tasks, result is %d expect:%d\n", sum, expect)
}
