package main

var NeverStop <-chan struct{} = make(chan struct{})

func main() {
	Example02()
}
