package sorting

import "testing"

func TestGreatestProductTree(t *testing.T) {
	type args struct {
		numbers []uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "case 1",
			args: args{
				numbers: []uint{1, 2, 3, 4, 6},
			},
			want: 72,
		},
		{
			name: "case 2",
			args: args{
				numbers: []uint{4, 6, 0, 1, 2, 3},
			},
			want: 72,
		},
		{
			name: "case 3",
			args: args{
				numbers: []uint{2, 3, 0, 1, 2, 10},
			},
			want: 60,
		},
		{
			name: "case 3",
			args: args{
				numbers: []uint{2, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreatestProductTree(tt.args.numbers); got != tt.want {
				t.Errorf("GreatestProductTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
