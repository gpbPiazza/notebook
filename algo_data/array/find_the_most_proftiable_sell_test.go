package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMostProftiableSell(t *testing.T) {
	stocks := []int{10, 7, 5, 8, 11, 2, 6}

	got := FindMostProftiableSell(stocks)

	assert.Equal(t, 6, got)
}
