package array

import "testing"

func TestCountAllCharsInTheStrings(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "shoul return 10 when has 10 chars",
			args: args{
				strings: []string{"ab", "c", "def", "gihj"},
			},
			want: 10,
		},
		{
			name: "shoul return 3 when has 3 chars in the slice",
			args: args{
				strings: []string{"ab", "c"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountAllChars(tt.args.strings); got != tt.want {
				t.Errorf("CountAllCharsInTheStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
