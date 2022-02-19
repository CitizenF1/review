package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	res := taskOne()
	expe := []int{0, 1, 2, 3, 4}

	assert.Equal(t, expe, res)
}
