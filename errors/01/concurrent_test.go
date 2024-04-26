package t

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"strconv"
	"sync"
	"testing"
	"time"
)

const LoopSize = 10

func TestSyncMap(t *testing.T) {
	syncMap := sync.Map{}
	valueStore := func(name string) {
		for i := 0; i < LoopSize; i++ {
			syncMap.Store(name+strconv.Itoa(i), "")
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(LoopSize)
	for i := 0; i < LoopSize; i++ {
		go func(val int) {
			defer wg.Done()
			valueStore("setter-" + strconv.Itoa(val) + "-")
		}(i)
	}
	wg.Wait()

	var l int
	syncMap.Range(func(k, v any) bool {
		l++
		return true
	})
	assert.Equal(t, LoopSize*LoopSize, l)
}

func TestOnceAndContext(t *testing.T) {
	once := sync.Once{}
	var val int
	cleanup := func() {
		val = LoopSize
	}

	doTask := func(ctx context.Context) {
		select {
		case <-ctx.Done():
			once.Do(cleanup)
		}
	}

	background := context.Background()
	timeout, cancelFunc := context.WithTimeout(background, 1*time.Second)
	defer cancelFunc()
	doTask(timeout)
	assert.Equal(t, val, LoopSize)
}

func TestOnce1(t *testing.T) {
	type Resource struct {
		Id int
	}

	var id int
	once := sync.Once{}
	var resource *Resource
	initResource := func() {
		resource = &Resource{Id: id}
		id++
	}
	getResource := func() *Resource {
		once.Do(func() {
			initResource()
		})
		return resource
	}

	wg := sync.WaitGroup{}
	wg.Add(LoopSize)
	for i := 0; i < LoopSize; i++ {
		go func() {
			defer wg.Done()
			r := getResource()
			assert.Equal(t, r.Id, 0)
		}()
	}
	wg.Wait()
}

func TestOnce(t *testing.T) {
	once0 := sync.Once{}
	var sum = 0
	for i := 0; i < LoopSize; i++ {
		once0.Do(func() {
			sum++
		})
	}
	assert.Equal(t, sum, 1)

	once1 := sync.Once{}
	var sum2 = 0
	once1.Do(func() {
		sum2++
	})
	once1.Do(func() {
		sum2 += LoopSize
	})
	assert.Equal(t, 1, sum2)
}

func TestPool(t *testing.T) {
	type User struct{ Id int }
	var autoIncrementId = 0
	pool := sync.Pool{New: func() any {
		autoIncrementId++
		return User{autoIncrementId}
	}}

	for i := 0; i < LoopSize; i++ {
		u := pool.Get().(User)
		pool.Put(u)
		// user 应该一直被复用
		assert.Equal(t, autoIncrementId, u.Id)
	}
}

func TestCond1(t *testing.T) {
	var mutex = sync.Mutex{}
	var queue []int

	producer := func() {
		for i := 0; ; i++ {
			mutex.Lock()
			queue = append(queue, i)
			mutex.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}

	consumer := func(name string) {
		for {
			mutex.Lock()

			if len(queue) != 0 {
				val := queue[0]
				fmt.Printf("%s receive : %d\n", name, val)
				queue = queue[1:]
			} else {
				fmt.Printf("%s : queue is empty\n", name)
			}

			mutex.Unlock()
		}
	}

	go producer()

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

func TestCond0(t *testing.T) {
	var mutex = sync.Mutex{}
	var cond = sync.NewCond(&mutex)
	var queue []int

	producer := func() {
		for i := 0; ; i++ {
			cond.L.Lock()
			queue = append(queue, i)
			cond.L.Unlock()
			cond.Signal()
			time.Sleep(100 * time.Millisecond)
		}
	}

	consumer := func(name string) {
		for {
			cond.L.Lock()
			// 第一次进入时，consumer 的 lock 和 producer 的 lock不能确定具体发生顺序
			// 但是我们可以确保通过 cond.Wait() 进入等待
			// 注意，这里必须是for循环去判断
			// 因为 cond.Wait() 中会先调用 cond.L.Unlock()
			// 也就是说，此时这段代码可能存在另外一个进程可以开始访问临界区
			// 举个例子
			// 1. consumer-0 Wait()，此时 consumer-0 释放锁，并且通过runtime_notifyListWait挂起自身
			// 2. producer 写入一条数据，但是尚未调用 Signal()
			// 3. consumer-1 判断 len(queue) != 0 消费掉 queue 中的数据并释放锁
			// 4. producer 调用 Signal() 唤醒 consumer-0
			// 5. consumer-0 此时操作 queue 就会导致空指针
			// 而如果当所有的进程都进入 Wait() 状态，那么此时反而不会有这个问题，
			// 因为每次 Signal() 只会启动一个协程
			for len(queue) == 0 {
				fmt.Printf("%s wait\n", name)
				cond.Wait()
			}
			val := queue[0]
			fmt.Printf("%s receive : %d\n", name, val)
			queue = queue[1:]
			cond.L.Unlock()
		}
	}

	go producer()

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

func TestMutex(t *testing.T) {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	mutex.Lock()

	countFunc := func(ret *int) {
		for i := 0; i < LoopSize; i++ {
			*ret = *ret + i
		}
	}

	var sum = 0
	countFunc(&sum)

	var mutexSum = 0
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer mutex.Unlock()
		// 休眠1秒等其他的协程执行
		time.Sleep(1 * time.Second)
		countFunc(&mutexSum)
		// 确认数据是否正常
		assert.Equal(t, sum, mutexSum)
		fmt.Printf("sum = %d, mutexSum = %d\n", sum, mutexSum)
	}()

	go func() {
		mutex.Lock()
		defer wg.Done()
		defer mutex.Unlock()
		countFunc(&mutexSum)
		assert.Equal(t, 2*sum, mutexSum)
		fmt.Printf("sum = %d, mutexSum = %d\n", sum, mutexSum)
	}()

	wg.Wait()
}

func TestTick(t *testing.T) {
	tick := time.Tick(1 * time.Second)
	after := time.After(5 * time.Second)

	var sum = 0
	var running = true
	for running {
		select {
		case <-tick:
			sum++
		case <-after:
			//case <-time.After(5 * time.Second):
			// 这里我们的代码是在一个for循环中，这意味着我们每次进入循环都新建了一个Timer对象
			running = false
		}
	}
	fmt.Println(sum)
}

func TestAfter(t *testing.T) {
	// 这个里面的assert大部分情况下是对的，然而实际上是未定义行为
	// 因为这里存在两个协程的竞争关系
	ch := make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		ch <- 0
	}()

	var selectedVal int
	select {
	case selectedVal = <-ch:
		break
	case _ = <-time.After(1 * time.Second):
		// 这个case有一个很特殊的地方，就是我们在select里初始化了Timer，
		// 在这个位置不会有问题，因为我们这里是没有循环的
		selectedVal = 10
		break
	}

	assert.Equal(t, 0, selectedVal)
}

func TestSelect(t *testing.T) {
	ch0 := make(chan int)
	ch1 := make(chan int)
	stop := make(chan *int)

	go func() {
		for i := 0; i < LoopSize; i++ {
			if i%2 == 0 {
				ch0 <- i
			} else {
				ch1 <- i
			}
		}

		stop <- nil
	}()

	running := true
	var sum int
	for running {
		select {
		case val := <-ch0:
			sum += val
		case val := <-ch1:
			sum += val
		case <-stop:
			running = false
		}
	}

	var sum2 int
	for i := 0; i < LoopSize; i++ {
		sum2 += i
	}

	assert.Equal(t, sum, sum2)
}

func TestRangeAndClose(t *testing.T) {
	ch := make(chan int)
	// 注意，这里虽然在闭包里引用了sum变量，但是并不会导致数据错误
	// 因为这里只启动了一个协程，闭包引用了这个变量
	// 闭包引起的变量引用异常，本质是协程竞争访问引起的问题
	var sum int
	go func() {
		for val := 0; val < LoopSize; val++ {
			sum += val
			ch <- val
		}
		close(ch)
	}()

	var sum2 int
	for val := range ch {
		sum2 += val
	}

	assert.Equal(t, sum, sum2)
}

func TestErrorGroup(t *testing.T) {
	var eg errgroup.Group
	var errorIndex = 8
	for i := 0; i < LoopSize; i++ {
		// 这里特别容易出错，因为 eg.Go() 是在内部调用 go func() 产生了一个闭包
		localVal := i
		eg.Go(func() error {
			// 在特定的位置抛出异常
			if localVal == errorIndex {
				return errors.New(strconv.Itoa(localVal))
			}
			return nil
		})
	}

	err := eg.Wait()

	assert.Error(t, err, strconv.Itoa(errorIndex))
}

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(LoopSize)
	for i := 0; i < LoopSize; i++ {
		go func() {
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestChan(t *testing.T) {
	input := 10

	// 不带缓冲区
	ch := make(chan int)
	//ch := make(chan int, 1)
	go func() {
		ch <- input
	}()

	val := <-ch

	assert.Equal(t, input, val)

}
