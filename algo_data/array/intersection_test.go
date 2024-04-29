package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return [2,4]",
			args: args{
				slice1: []int{1, 2, 3, 4, 5},
				slice2: []int{0, 2, 4, 6, 8},
			},
			want: []int{2, 4},
		},
		{
			name: "changing the slices position should still return [2,4]",
			args: args{
				slice2: []int{1, 2, 3, 4, 5},
				slice1: []int{0, 2, 4, 6, 8},
			},
			want: []int{2, 4},
		},
		{
			name: "slice order should not matter still return [4,2]",
			args: args{
				slice2: []int{4, 5, 2, 1, 3},
				slice1: []int{0, 6, 2, 4, 8},
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.slice1, tt.args.slice2); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
