package bench

import (
	_func "main/func"
	"testing"
)

// 需要特别注意的是N，go test会一直调整这个数值，直到测试时间能得出可靠的性能数据为止。

func BenchmarkFib1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_func.Fib1(20)
	}
}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_func.Fib2(20)
	}
}

func BenchmarkFib3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_func.Fib3(20)
	}
}