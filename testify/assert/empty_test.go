package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Empty断言object是空，根据object中存储的实际类型，空的含义不同：
//指针：nil；
//整数：0；
//浮点数：0.0；
//字符串：空串""；
//布尔：false；
//切片或 channel：长度为 0。
func TestEmpty(t *testing.T) {
	assert.Empty(t, nil)
	assert.Empty(t, 0)
	assert.Empty(t, 0.0)
	assert.Empty(t, "")
	assert.Empty(t, false)
	assert.Empty(t, []int{})
}
