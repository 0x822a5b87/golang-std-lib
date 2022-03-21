package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertion(t *testing.T) {
	assertions := assert.New(t)
	assertions.Empty("")
	assertions.Contains("Hello world!!", "Hello")
}
