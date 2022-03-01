package main

// 为什么会需要 goroutine 池？
// 1. goroutine 每个需要8KB，也是有有内存占用额
// 2. goroutine 只能等执行完毕后自动退出，如果由于某种原因 goroutine 无法正常退出会引起 goroutine 泄漏
// 3. 频繁的创建 goroutine 也有开销
func main() {
	example03()
}
