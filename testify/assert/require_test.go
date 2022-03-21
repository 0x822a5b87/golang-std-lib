package assert

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestRequire require提供了和assert同样的接口，但是遇到错误时，require直接终止测试，而assert返回false。
func TestRequire(t *testing.T) {
	require.Equal(t, "hello", "hello")
}
