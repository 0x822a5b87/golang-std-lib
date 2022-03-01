package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
)

type taskFunc func()

// taskFuncWrapper ants支持使用无参函数构造协程池，所以需要通过闭包来传递
func taskFuncWrapper(nums []int, i int, sum *int, wg *sync.WaitGroup) taskFunc {
	return func() {
		for _, num := range nums[i*DataPerTask : (i+1)*DataPerTask] {
			*sum += num
		}

		fmt.Printf("task:%d sum:%d\n", i+1, *sum)
		wg.Done()
	}
}

func example03() {
	// 在构造协程池的时候不指定执行的函数
	p, _ := ants.NewPool(10)
	defer p.Release()

	nums := make([]int, DataSize, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	partSums := make([]int, DataSize/DataPerTask, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		// 在协程池中执行无参函数
		err := p.Submit(taskFuncWrapper(nums, i, &partSums[i], &wg))
		if err != nil {
			return
		}
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	var sum int
	for _, num := range partSums {
		sum += num
	}

	var expect int
	for _, num := range nums {
		expect += num
	}

	fmt.Printf("finish all tasks, result is %d expect:%d\n", sum, expect)
}
