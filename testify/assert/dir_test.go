package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDirExists DirExists断言路径path是一个目录，如果path不存在或者是一个文件，断言失败。
func TestDirExists(t *testing.T) {
	assert.DirExists(t, "/dev")
}
