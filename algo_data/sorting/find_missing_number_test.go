package sorting

import "testing"

func Test_findMissingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{5, 2, 4, 1, 0},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{9, 3, 2, 5, 6, 7, 1, 0, 4},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingNumber(tt.args.nums); got != tt.want {
				t.Errorf("finMissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
