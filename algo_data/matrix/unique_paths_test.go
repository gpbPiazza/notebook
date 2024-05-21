package matrix

import "testing"

func TestUniquePaths(t *testing.T) {
	type args struct {
		rows    int
		columns int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 1 when there is only one path",
			want: 1,
			args: args{
				rows:    1,
				columns: 7,
			},
		},
		{
			name: "should return 4",
			want: 4,
			args: args{
				rows:    2,
				columns: 4,
			},
		},
		{
			name: "should return 28",
			want: 28,
			args: args{
				rows:    3,
				columns: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.rows, tt.args.columns); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}

			if got := uniquePaths_with_memo(tt.args.rows, tt.args.columns, make(map[[2]int]int)); got != tt.want {
				t.Errorf("uniquePaths_with_memo() = %v, want %v", got, tt.want)
			}
		})
	}
}
