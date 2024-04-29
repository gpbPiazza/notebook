package array

import "testing"

func TestFrstNonDuplicated(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return n since n is the first non duplicated",
			args: args{
				text: "minimum",
			},
			want: "n",
		},
		{
			name: "should return u since u is the first non duplicated",
			args: args{
				text: "munimam",
			},
			want: "u",
		},
		{
			name: "should return u since u is the first non duplicated at the last index",
			args: args{
				text: "ababwiiwccu",
			},
			want: "u",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FrstNonDuplicated(tt.args.text); got != tt.want {
				t.Errorf("FrstNonDuplicated() = %v, want %v", got, tt.want)
			}
		})
	}
}
