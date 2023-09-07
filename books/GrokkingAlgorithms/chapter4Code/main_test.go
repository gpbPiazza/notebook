package main

import (
	"reflect"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	orderedInts := []int{1, 2, 3, 4, 5, 6, 12, 55, 66, 69, 77, 88}
	oneElementSlice := []int{1}
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "give slice with any elements and has the target should return target",
			args: args{
				nums:   orderedInts,
				target: 77,
			},
			want: 77,
		},
		{
			name: "give slice with one elements and this elements is not target should return -1",
			args: args{
				nums:   oneElementSlice,
				target: 44,
			},
			want: -1,
		},
		{
			name: "give empty slice should return -1",
			args: args{
				nums:   []int{},
				target: 9009,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findUniqueVals(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "should return [6, 1, 2, 3]",
			nums: []int{6, 5, 1, 2, 3, 4, 4, 5, 7, 7},
			want: []int{6, 1, 2, 3},
		},
		{
			name: "should return [1,2,3]",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findUniqueVals(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findUniqueVals() = %v, want %v", got, tt.want)
			}
		})
	}
}
