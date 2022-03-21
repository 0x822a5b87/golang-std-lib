package testify

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

// 再次总结 mock 的调用
// 1. 定义需要 mock 的接口以便于后续调用
// 2. 通过 mock.On(methodName, arguments...).Return(returnArguments...) 定义在
//		<1> 中mock的接口在不同参数下的返回值
// 3. 实际调用 mock 接口
// 4. 判断是否符合我们的期望。

type ExampleMock struct {
	mock.Mock
}

func (e *ExampleMock) Hello(n int) int {
	args := e.Called(n)
	return args.Int(0)
}

func TestExample(t *testing.T) {
	e := new(ExampleMock)

	// 定义函数 Hello 在不同的参数下的返回值
	e.On("Hello", 1).Return(1).Times(1)
	e.On("Hello", 2).Return(2).Times(2)
	e.On("Hello", 3).Return(3).Times(3)

	// 调用 Hello 方法
	ExampleFunc(e)

	e.AssertExpectations(t)
	e.AssertCalled(t, "Hello", 1)
	e.AssertCalled(t, "Hello", 2)
	e.AssertCalled(t, "Hello", 3)

	e.AssertNumberOfCalls(t, "Hello", 6)
}
