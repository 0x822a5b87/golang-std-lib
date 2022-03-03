package unit

import (
	"testing"
	"main/func"
)

type parameter struct {
	left          int64
	right         int64
	expected      int64
	operator      rune
	expectedError error
}

// 在 Go 中编写测试很简单，只需要在待测试功能所在文件的同级目录中创建一个以_test.go结尾的文件。
// 在该文件中，我们可以编写一个个测试函数。测试函数名必须是TestXxx这个形式，
// 而且Xxx必须以大写字母开头，另外函数带有一个*testing.T类型的参数：

// 表格驱动测试
func TestTable(t *testing.T) {
	testCases := []parameter{
		{1, 1, 2, '+', nil},
		{-1, 1, 0, '+', nil},
		{-1, 1, -1, '*', nil},
		{1, -1, -1, '*', nil},
		{-1, -1, 1, '*', nil},
	}
	for _, testCase := range testCases {
		testMath(testCase, t)
	}
}

func TestAdd(t *testing.T) {
	testMath(parameter{
		left:          0,
		right:         1,
		expected:      1,
		operator:      '+',
		expectedError: nil,
	}, t)
}

func TestSub(t *testing.T) {
	testMath(parameter{
		left:          1,
		right:         0,
		expected:      1,
		operator:      '-',
		expectedError: nil,
	}, t)
}

func TestMul(t *testing.T) {
	testMath(parameter{
		left:          1,
		right:         0,
		expected:      0,
		operator:      '*',
		expectedError: nil,
	}, t)
}

func TestDiv(t *testing.T) {
	testMath(parameter{
		left:          1,
		right:         0,
		expected:      1,
		operator:      '/',
		expectedError: nil,
	}, t)
}

func testMath(p parameter, t *testing.T) {
	ret, err := _func.Math(p.left, p.right, p.operator)

	if p.expectedError != nil {
		if err != p.expectedError {
			t.Errorf("expected error: [%s], got: [%s]", p.expectedError, err)
		}
		return
	}
	if ret != p.expected {
		t.Errorf("%d %c %d, expected : %d, got : %d",
			p.left, p.operator, p.right, p.expected, ret)
	}
}
