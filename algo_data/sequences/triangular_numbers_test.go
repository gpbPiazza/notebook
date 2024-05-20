package sequences

import "testing"

func TestGetTriangularNumber(t *testing.T) {
	type args struct {
		nthNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 28 when nthNumber is 7",
			args: args{
				nthNumber: 7,
			},
			want: 28,
		},
		{
			name: "should return 3 when nthNumber is 2",
			args: args{
				nthNumber: 2,
			},
			want: 3,
		},
		{
			name: "should return 15 when nthNumber is 5",
			args: args{
				nthNumber: 5,
			},
			want: 15,
		},
		{
			name: "should return 66 when nthNumber is 11",
			args: args{
				nthNumber: 11,
			},
			want: 66,
		},
		{
			name: "should return 1 when nthNumber is 0",
			args: args{
				nthNumber: 0,
			},
			want: 1,
		},
		{
			name: "should return 1 when nthNumber is -2",
			args: args{
				nthNumber: -2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTriangularNumber(tt.args.nthNumber); got != tt.want {
				t.Errorf("GetTriangularNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
