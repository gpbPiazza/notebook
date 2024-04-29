package array

import "testing"

func TestFindFirstDuplicated(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return c",
			args: args{
				slice: []string{"a", "b", "c", "d", "c", "e", "f"},
			},
			want: "c",
		},
		{
			name: "should return a",
			args: args{
				slice: []string{"a", "a", "c", "d", "c", "e", "f"},
			},
			want: "a",
		},
		{
			name: "should return d since d is the frist duplicated",
			args: args{
				slice: []string{"d", "a", "c", "d", "c", "e", "f"},
			},
			want: "d",
		},
		{
			name: "should return empty string since there is no duplicated",
			args: args{
				slice: []string{"a", "b", "c", "d", "e"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFirstDuplicated(tt.args.slice); got != tt.want {
				t.Errorf("FindFirstDuplicated() = %v, want %v", got, tt.want)
			}
		})
	}
}
