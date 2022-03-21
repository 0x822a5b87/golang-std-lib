package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestElementMatch 断言listA和listB包含相同的元素，忽略元素出现的顺序。
// listA/listB必须是数组或切片。如果有重复元素，重复元素出现的次数也必须相等。
func TestElementMatch(t *testing.T) {
	var listA = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	var listB = [][]int{
		{4, 5, 6},
		{1, 2, 3},
		{7, 8, 9},
	}

	assert.ElementsMatch(t, listA, listB)
}
