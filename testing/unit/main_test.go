package unit

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

//有一种特殊的测试函数，函数名为TestMain()，接受一个*testing.M类型的参数。
//这个函数一般用于在运行所有测试前执行一些初始化逻辑（如创建数据库链接），或所有测试都运行结束之后执行一些清理逻辑（释放数据库链接）。
//如果测试文件中定义了这个函数，则go test命令会直接运行这个函数，否者go test会创建一个默认的TestMain()函数。
//这个函数的默认行为就是运行文件中定义的测试。我们自定义TestMain()函数时，也需要手动调用m.Run()方法运行测试函数，否则测试函数不会运行。
func TestMain(m *testing.M) {
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("name:%s usage:%s value:%v\n", f.Name, f.Usage, f.Value)
	})
	os.Exit(m.Run())
}