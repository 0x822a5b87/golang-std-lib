package assert

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// EqualError断言theError.Error()的返回值与errString相等。
func TestEqualError(t *testing.T) {
	assert.EqualError(t, errors.New("test equal error"), "test equal error")
}
