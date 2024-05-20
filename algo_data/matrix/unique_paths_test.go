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
			name: "should return 2 when there is 2 paths",
			want: 2,
			args: args{
				rows:    2,
				columns: 4,
			},
		},
		{
			name: "should return 3 when there is 3 paths",
			want: 3,
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
		})
	}
}
