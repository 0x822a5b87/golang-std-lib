package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

func example01() {
	length := 10
	arr := make([]int, length)
	for i := range arr {
		arr[i] = i + 1
	}
	wg := &sync.WaitGroup{}
	(*wg).Add(1)
	task := Task{
		index: 0,
		nums:  arr,
		sum:   0,
		wg:    wg,
	}
	example01Pool(&task)
	wg.Wait()
}

func example01Func(data interface{}) {
	task := data.(*Task)
	task.DoTask()
	fmt.Printf("task:%d sum:%d\n", task.index, task.sum)
}

func example01Pool(task *Task) {
	// 创建 goroutine 池
	p, _ := ants.NewPoolWithFunc(10, example01Func)
	defer p.Release()

	// p.Invoke(task) 会从 ants 空闲的 goroutine 找一个来执行
	// ants.NewPoolWithFunc 中传递的参数，并使用 task 作为参数来执行。
	err := p.Invoke(task)
	if err != nil {
		logger.Info(fmt.Sprint("p.Invoke error: ", err))
		return
	}
}
