package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	zap := []int{1, 2, 3, 4, 5}
	got := Reverse(zap)

	assert.Equal(t, []int{5, 4, 3, 2, 1}, got)
}
