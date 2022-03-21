package assert

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test01(t *testing.T) {
	var expected = 10
	var actual = 10
	assert.Equal(t, expected, actual, fmt.Sprintf("expected [%d], actual [%d]", expected, actual))
}
