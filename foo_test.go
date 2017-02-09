package foo

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlatform(t *testing.T) {
	assert.Equal(t, "windows", runtime.GOOS)
}

func TestDoSomething(t *testing.T) {
	assert.Equal(t, 123, doSomething())
}
