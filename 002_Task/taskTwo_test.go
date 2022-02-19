package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskTwo(t *testing.T) {
	str := taskTwo()
	expec := "any string"

	assert.Equal(t, str, expec)
}
