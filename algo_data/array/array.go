package array

import (
	"errors"
	"fmt"
)

var (
	ErrValueNotFound = errors.New("the given element not found in the array")
)

type Array struct {
	len      int
	capacity int
	data     []int
}

func newArray(capacity int) *Array {
	return &Array{
		capacity: capacity,
		len:      0,
		data:     []int{},
	}
}

func (a *Array) Read(i int) int {
	if i > len(a.data) || i < 0 {
		panic(fmt.Sprintf("index %d is oute of range %d", i, len(a.data)))
	}

	return a.data[i]
}

func (a *Array) Search(val int) (int, error) {
	for i := 0; i < len(a.data); i++ {
		if a.data[i] == val {
			return a.data[i], nil
		}
	}

	return 0, ErrValueNotFound
}

// Append function will insert a new val in the array always at the of it
// will panics if the capacity over flow
func (a *Array) Append(val int) {
	newLen := a.len + 1

	if a.capacity < newLen {
		panic("capacity over flow")
	}

	a.data = append(a.data, val)
	a.len = newLen
}

func (a *Array) Len() int {
	return a.len
}

func (a *Array) InsertAt(val int, i int) int {
	return 0
}
