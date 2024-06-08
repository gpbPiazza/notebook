package sorting

import "sort"

// findMissingNumber expects to receive a slice of numbers
// from 0 to the lenght of the slice
func findMissingNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] > nums[i]
	})

	for i := range nums {
		num := nums[i]
		if num != i {
			return i
		}
	}

	return 0
}
