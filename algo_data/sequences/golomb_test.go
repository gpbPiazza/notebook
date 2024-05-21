package sequences

import "testing"

// 1, 2, 2, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 12, 12, 12, 12, 12, 12

func Test_golomb(t *testing.T) {
	type args struct {
		nth int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nth: 6},
			want: 4,
		},
		{
			args: args{nth: 5},
			want: 3,
		},
		{
			args: args{nth: 2},
			want: 2,
		},
		{
			args: args{nth: 20},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := golomb(tt.args.nth); got != tt.want {
				t.Errorf("golomb() = %v, want %v", got, tt.want)
			}
			if got := golomb_with_memo(tt.args.nth, make(map[int]int)); got != tt.want {
				t.Errorf("golomb_with_memo() = %v, want %v", got, tt.want)
			}
		})
	}
}
