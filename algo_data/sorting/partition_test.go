package sorting

import (
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "book case",
			args: args{
				nums:  []int{0, 1, 2, 5, 6, 3},
				left:  0,
				right: 5,
			},
		},
		{
			name: "1",
			args: args{
				nums:  []int{0, 7, 1, 8, 2, 9, 3, 5, 6},
				left:  0,
				right: 8,
			},
		},

		{
			name: "take middle of the array",
			args: args{
				nums:  []int{0, 7, 1, 8, 2, 9, 3, 5, 6},
				left:  0,
				right: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pivotVal := tt.args.nums[tt.args.right]

			_ = partition(tt.args.nums, tt.args.left, tt.args.right)

			var leftVals []int
			var rightVals []int
			for index, zap := range tt.args.nums {
				if zap == pivotVal {
					leftVals = tt.args.nums[:index]
					rightVals = tt.args.nums[index+1:]
					break
				}
			}

			for i := range leftVals {
				if !(leftVals[i] < pivotVal) {
					t.Errorf("error on leftVals - found %d pivot is %d", leftVals[i], pivotVal)
				}
			}

			for i := range rightVals {
				if !(rightVals[i] > pivotVal) {
					t.Errorf("error on rightVals - found %d pivot is %d", rightVals[i], pivotVal)
				}
			}
		})
	}
}
