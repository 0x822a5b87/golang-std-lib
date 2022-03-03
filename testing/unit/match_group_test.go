package unit

import "testing"

// 分组测试
// 分别测试正数和负数

var (
	positiveGroup []parameter
	negativeGroup []parameter
)

// TestMathGroup 使用分组测试
func TestMathGroup(t *testing.T) {
	t.Run("positive", testPositiveGroup)
	t.Run("negative", testNegativeGroup)
}

func testPositiveGroup(t *testing.T) {
	for _, testCase := range positiveGroup {
		testMath(testCase, t)
	}
}

func testNegativeGroup(t *testing.T) {
	for _, testCase := range negativeGroup {
		testMath(testCase, t)
	}
}

func init() {
	positiveGroup = []parameter{
		{1, 1, 2, '+', nil},
	}
	negativeGroup = []parameter{
		{-1, -1, -2, '+', nil},
	}
}