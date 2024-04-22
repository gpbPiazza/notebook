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

func (a *Array) LinearSearch(val int) (int, error) {
	for i := 0; i < len(a.data); i++ {
		if a.data[i] == val {
			return a.data[i], nil
		}
	}

	return 0, ErrValueNotFound
}

func (a *Array) BinarySearch(val int) (int, error) {
	isValFound := false
	arrayToLookUp := a.data
	for !isValFound {
		middleIndex := len(arrayToLookUp) / 2
		guessedVal := arrayToLookUp[middleIndex]

		if guessedVal == val {
			isValFound = true
			return val, nil
		}

		if middleIndex == 0 {
			return 0, ErrValueNotFound
		}

		if guessedVal > val {
			arrayToLookUp = arrayToLookUp[:middleIndex]
		} else {
			arrayToLookUp = arrayToLookUp[middleIndex:]
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
