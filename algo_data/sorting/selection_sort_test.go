package sorting

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "should sort ASC 1",
			input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "should sort ASC 2",
			input: []int{9, 5, 4, 3, 2, 1, 8, 7, 6},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "should sort ASC 3",
			input: []int{-9, 0, 5, -24, 3, 2, 1, 8, 7, 6},
			want:  []int{-24, -9, 0, 1, 2, 3, 5, 6, 7, 8},
		},
		{
			name:  "given a alredy sorted array should return sorted ASC 4",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectionSort(tt.input)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("%v - want: %d got %d", tt.name, tt.want, got)
			}
		})
	}
}
