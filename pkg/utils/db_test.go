package utils

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSQLSlice(t *testing.T) {
	res := SQLSlice([]int{1, 2})
	assert.Equal(t, "1,2", res)
}
