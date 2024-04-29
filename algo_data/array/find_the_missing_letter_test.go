package array

import "testing"

func TestFindMissingAlphabetLetter(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return f",
			args: args{
				text: "the quick brown box jumps over a lazy dog",
			},
			want: "f",
		},
		{
			name: "should return f",
			args: args{
				text: "the quick brown box jumps over a lazy dof",
			},
			want: "g",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMissingAlphabetLetter(tt.args.text); got != tt.want {
				t.Errorf("FindMissingAlphabetLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
