package frequencysortstr

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				s: "tree",
			},
			want: "eert",
		},
		{
			name: "02",
			args: args{
				s: "cccaaa",
			},
			want: "aaaccc",
		},
		{
			name: "03",
			args: args{
				s: "Aabb",
			},
			want: "bbAa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
