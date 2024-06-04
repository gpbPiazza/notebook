package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "book case",
			args: args{
				nums:  []int{0, 1, 2, 3, 6, 5},
				left:  0,
				right: 5,
			},
			want: []int{0, 1, 2, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums, tt.args.left, tt.args.right)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
