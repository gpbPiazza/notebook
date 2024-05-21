package array

import "testing"

func Test_addUntil100(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should ignore sum indexes 0,1 and return 90",
			args: args{numbers: []int{20, 30, 40, 50}},
			want: 90,
		},
		{
			name: "should ignore sum indexe 0 and return 95",
			args: args{numbers: []int{99, 10, 20, 55, 10}},
			want: 95,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addUntil100(tt.args.numbers); got != tt.want {
				t.Errorf("addUntil100() = %v, want %v", got, tt.want)
			}
			if got := addUntil100_On_implementation(tt.args.numbers); got != tt.want {
				t.Errorf("addUntil100_On_implementation() = %v, want %v", got, tt.want)
			}
		})
	}
}
