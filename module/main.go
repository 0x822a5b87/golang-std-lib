package main

import (
	"fmt"
	"main/sub"
)

func main() {
	fmt.Println("hello module")
	fmt.Println(sub.Add(10, 20))
}