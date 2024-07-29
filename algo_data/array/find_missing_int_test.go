package array

import "testing"

func TestFindMissingInt(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "retrun missing int 4",
			args: args{
				nums: []int{2, 3, 0, 6, 1, 5},
			},
			want: 4,
		},
		{
			name: "retrun missing int 1",
			args: args{
				nums: []int{8, 2, 3, 9, 4, 7, 5, 0, 6},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMissingInt(tt.args.nums); got != tt.want {
				t.Errorf("FindMissingInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
