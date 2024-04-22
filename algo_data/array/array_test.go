package array

import (
	"runtime/debug"
	"testing"
)

func didPanic(f func()) (didPanic bool, message interface{}, stack string) {
	didPanic = true

	defer func() {
		message = recover()
		if didPanic {
			stack = string(debug.Stack())
		}
	}()

	f()
	didPanic = false

	return
}

func TestArray_Read(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		array       *Array
		shouldPanic bool
		want        int
	}{
		{
			name:  "given index 0 should return 1 as first element in the array",
			input: 0,
			array: &Array{data: []int{1, 2, 3, 4}},
			want:  1,
		},
		{
			name:  "given the last index should return 4 as first element in the array",
			input: 3,
			array: &Array{data: []int{1, 2, 3, 4}},
			want:  4,
		},
		{
			name:  "given the middle index should return 2 as first element in the array",
			input: 2,
			array: &Array{data: []int{0, 1, 2, 3, 4}},
			want:  2,
		},
		{
			name:        "given the input negative should panics",
			input:       -2,
			array:       &Array{data: []int{0, 1, 2, 3, 4}},
			shouldPanic: true,
		},
		{
			name:        "given the input bigger then len of array should panics",
			input:       20,
			array:       &Array{data: []int{0, 1, 2, 3, 4}},
			shouldPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				didPanic, message, stack := didPanic(func() { tt.array.Read(tt.input) })
				if !didPanic {
					t.Errorf("Array.Read() should have panic message: %v stack%v", message, stack)
				}
			} else {
				if got := tt.array.Read(tt.input); got != tt.want {
					t.Errorf("Array.Read() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestArray_LinearSearch(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		array   *Array
		wantErr error
		want    int
	}{
		{
			name:    "given value 1 should return 1",
			input:   1,
			array:   &Array{data: []int{1, 2, 3, 4}},
			want:    1,
			wantErr: nil,
		},
		{
			name:    "given value 4 should return 4",
			input:   4,
			array:   &Array{data: []int{1, 2, 3, 4}},
			want:    4,
			wantErr: nil,
		},
		{
			name:    "given value not existing in the array should return err",
			input:   20,
			array:   &Array{data: []int{1, 2, 3, 4}},
			want:    0,
			wantErr: ErrValueNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.array.LinearSearch(tt.input); got != tt.want && err != tt.wantErr {
				t.Errorf("Array.Search() = %v, %v, want %v %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestArray_BinarySearch(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		array   *Array
		wantErr error
		want    int
	}{
		{
			name:    "given value 1 should return 1",
			input:   1,
			array:   &Array{data: []int{1, 2, 3, 4}},
			want:    1,
			wantErr: nil,
		},
		{
			name:    "given value 4 should return 4",
			input:   4,
			array:   &Array{data: []int{1, 2, 3, 4}},
			want:    4,
			wantErr: nil,
		},
		{
			name:    "given value not existing in the array should return err",
			input:   20,
			array:   &Array{data: []int{1, 2, 3, 4}},
			want:    0,
			wantErr: ErrValueNotFound,
		},
		{
			name:    "given value 19 should return 19",
			input:   19,
			array:   &Array{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}},
			want:    19,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.array.BinarySearch(tt.input); got != tt.want && err != tt.wantErr {
				t.Errorf("Array.Search() = %v, %v, want %v %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestArray_Append(t *testing.T) {
	t.Run("given a array with capacity 5 insert new element should not panic", func(t *testing.T) {
		array := newArray(5)

		array.Append(0)

		val := array.Read(0)

		if val != 0 {
			t.Errorf("Inserted val 0 and received %d", val)
		}
	})

	t.Run("given a array with capacity 2 insert two elements and val -2 should be at the index 0 and -1 index 1", func(t *testing.T) {
		array := newArray(2)

		array.Append(-2)
		if array.Len() != 1 {
			t.Errorf("Expected len be 0")
		}
		array.Append(-1)
		if array.Len() != 2 {
			t.Errorf("Expected len be 1")
		}

		val := array.Read(0)

		if val != -2 {
			t.Errorf("Expected val %d be at the index %d and received %d", 2, 0, val)
		}

		val = array.Read(1)

		if val != -1 {
			t.Errorf("Expected val %d be at the index %d and received %d", -1, 1, val)
		}
	})

	t.Run("given a array with capacity 2 insert tree elements and should panics", func(t *testing.T) {
		array := newArray(2)

		array.Append(-2)
		if array.Len() != 1 {
			t.Errorf("Expected len be 1")
		}

		array.Append(-1)
		if array.Len() != 2 {
			t.Errorf("Expected len be 2")
		}

		val := array.Read(0)
		if val != -2 {
			t.Errorf("Expected val %d be at the index %d and received %d", 2, 0, val)
		}

		val = array.Read(1)
		if val != -1 {
			t.Errorf("Expected val %d be at the index %d and received %d", -1, 1, val)
		}

		didPanic, message, stack := didPanic(func() { array.Append(10) })
		if !didPanic {
			t.Errorf("Should have panic message %v, stack %v", message, stack)
		}
	})
}
