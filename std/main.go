package main

import (
	_select "hangyudu/std/select"
)

func main() {
	_select.NewTimer(_select.NewKafkaReader([]string{"127.0.0.1:9092"}, "test", "test"))
}
