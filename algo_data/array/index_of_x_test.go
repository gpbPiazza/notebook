package array

import "testing"

func TestIndexOfX(t *testing.T) {
	type args struct {
		text  string
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should find x in the index 23",
			args: args{
				text:  "abcdefghijklmnopqrstuvwxyz",
				index: 0,
			},
			want: 23,
		},
		{
			name: "should find x in the index 4",
			args: args{
				text:  "abcdxyz",
				index: 0,
			},
			want: 4,
		},
		{
			name: "should find the first x in the index 2",
			args: args{
				text:  "abxdxyz",
				index: 0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexOfX(tt.args.text, tt.args.index); got != tt.want {
				t.Errorf("indexOfX() = %v, want %v", got, tt.want)
			}
			if got := indexOfX_SecondImplementation(tt.args.text); got != tt.want {
				t.Errorf("indexOfX_SecondImplementation() = %v, want %v", got, tt.want)
			}
		})
	}
}
