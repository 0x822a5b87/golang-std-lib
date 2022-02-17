package main

import (
	"flag"
	"fmt"
)

func example01Main() {
	flag.Parse()

	fmt.Println("int flag:", intFlag)
	fmt.Println("bool flag:", boolFlag)
	fmt.Println("string flag:", stringFlag)

	// 输出在从输入参数读取完之后，剩余的参数列表
	fmt.Println("flag.Args : ", flag.Args())
	// 输出未读取的参数数量
	fmt.Println("flag.NArg: ", flag.NArg())
	// 输出未解析的参数的 [0]
	fmt.Println("flag.Arg[0]: ", flag.Arg(0))

	fmt.Println("flag.NFlag: ", flag.NFlag())
}

func example02Main() {
	flag.Parse()

	fmt.Println("int flag:", *pIntFlag)
	fmt.Println("bool flag:", *pBoolFlag)
	fmt.Println("string flag:", *pStringFlag)
}

func example03Main() {
	flag.Parse()

	fmt.Println(intervalFlag)
}

func main() {
	example04Main()
}
