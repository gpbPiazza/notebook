package array

import "testing"

func TestFindHighestProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "return 60",
		// 	args: args{
		// 		nums: []int{5, -10, -6, 9, 4},
		// 	},
		// 	want: 60,
		// },
		{
			name: "return 20",
			args: args{
				nums: []int{-5, -4, -3, 0, 3, 4},
			},
			want: 20,
		},
		{
			name: "return 21",
			args: args{
				nums: []int{-9, -2, -1, 2, 3, 7},
			},
			want: 21,
		},
		{
			name: "return 28",
			args: args{
				nums: []int{-7, -4, -3, 0, 4, 6},
			},
			want: 28,
		},
		{
			name: "return 30",
			args: args{
				nums: []int{-6, -5, -1, 2, 3, 9},
			},
			want: 30,
		},
		{
			name: "return 42",
			args: args{
				nums: []int{-9, -4, -3, 0, 6, 7},
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindHighestProduct(tt.args.nums); got != tt.want {
				t.Errorf("FindHighestProduct() = %v, want %v", got, tt.want)
			}
			if got := FindHighestProduct_BookImplementation(tt.args.nums); got != tt.want {
				t.Errorf("FindHighestProduct_BookImplementation() = %v, want %v", got, tt.want)
			}
		})
	}
}
