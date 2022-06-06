package context

import (
	"context"
	"fmt"
	"time"
)

func TimeoutContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println("TimeOutContext", ctx)
	defer cancel()
	go Monitor(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("context done!")
	}

	time.Sleep(20 * time.Second)
}

func Monitor(ctx context.Context) {
	fmt.Println("Monitor", ctx)
	i := 0
	for {
		fmt.Printf("%d\n", i)
		i++
		time.Sleep(1 * time.Second)
	}
}

//0
//1
//2
//3
//4
//5
//6
//7
//8
//9
//context done!
//10
//11
//12
//13
//14
//15
//16
//17
//18
//19
//20
//21
//22
//23
//24
//25
//26
//27
