package example

import (
	"fmt"
	_func "main/func"
)

func Example() {
	v, err := _func.Math(1, 2, '+')
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(v)
}