package sorting

import "testing"

func TestFindGreatest_On(t *testing.T) {
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
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{9, 3, 2, 5, 6, 7, 1, 0, 4},
			},
			want: 9,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{99, 3, 2, 5, 6, 7, 1, 0, 4},
			},
			want: 99,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{-3, 2, -5, 6, 799, 1, 0, 999},
			},
			want: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindGreatest_O_of_n(tt.args.nums); got != tt.want {
				t.Errorf("FindGreatest_O_of_n() = %v, want %v", got, tt.want)
			}
			if got := FindGreatest_O_of_n_2(tt.args.nums); got != tt.want {
				t.Errorf("FindGreatest_O_of_n_2() = %v, want %v", got, tt.want)
			}
			if got := FindGreatest_O_of_n_log_n(tt.args.nums); got != tt.want {
				t.Errorf("FindGreatest_O_of_n_log_n() = %v, want %v", got, tt.want)
			}
		})
	}
}
