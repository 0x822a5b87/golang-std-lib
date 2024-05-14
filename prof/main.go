package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var dataList []string

func Add(str string) int {
	// This code is to ensure that it takes up more CPU to ensure better debugging
	for i := 0; i < 10000000000; i++ {
		if i%100000000 == 0 {
			log.Printf("add\n")
		}
	}
	data := []byte(str)
	dataList = append(dataList, string(data))
	return len(dataList)
}

func main() {
	go func() {
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}
