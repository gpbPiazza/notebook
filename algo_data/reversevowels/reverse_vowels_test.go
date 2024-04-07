package reversevowels

import (
	"testing"
)

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "given 'hello' should return 'holle'",
			input: "hello",
			want:  "holle",
		},
		{
			name:  "given 'leetcode' should return 'leotcede'",
			input: "leetcode",
			want:  "leotcede",
		},
		{
			name:  "given 'piazza' should return 'paazzi'",
			input: "piazza",
			want:  "paazzi",
		},
		{
			name:  "given 'gabriel' should return 'gabriel'",
			input: "gabriel",
			want:  "gebrial",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.input); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "given 'hello' should return 'holle'",
			input: "hello",
			want:  "olleh",
		},
		{
			name:  "given 'leetcode' should return 'edocteel'",
			input: "leetcode",
			want:  "edocteel",
		},
		{
			name:  "given 'piazza' should return 'azzaip'",
			input: "piazza",
			want:  "azzaip",
		},
		{
			name:  "given 'gabriel' should return 'leirbag'",
			input: "gabriel",
			want:  "leirbag",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.input); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
