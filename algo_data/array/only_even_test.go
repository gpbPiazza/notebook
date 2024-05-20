package array

import (
	"reflect"
	"testing"
)

func TestOnlyEven(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return 0, 2 and 4",
			args: args{
				numbers: []int{1, 5, 67, 0, 2, 1, 4},
			},
			want: []int{0, 2, 4},
		},
		{
			name: "should return all elements",
			args: args{
				numbers: []int{12, 44, 62, 0, 2, 18, 4},
			},
			want: []int{12, 44, 62, 0, 2, 18, 4},
		},
		{
			name: "should return empty list when dont have even numbers in the list",
			args: args{
				numbers: []int{1, 5, 67, 3, 77, 85, 95},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OnlyEven(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnlyEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
