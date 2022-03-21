package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	testStr(t)
	testSlice(t)
}

func testStr(t *testing.T) {
	var expected = "Hello world!"
	var actual = "Hello"
	assert.Contains(t, expected, actual)
}

// testSlice 注意，s 可以是字符串，数组/切片，map
// 相应的，contains为子串，数组/切片元素，map 的键
func testSlice(t *testing.T) {
	var s = []int{1, 2, 3, 4, 5}
	var contains = 1
	assert.Contains(t, s, contains)

	var notContains = []int{1, 2, 3, 4, 5}
	assert.NotContains(t, s, notContains)

	var s2 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var contains2 = []int{1, 2, 3}
	assert.Contains(t, s2, contains2)
}
