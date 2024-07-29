package array

import "slices"

// [2,3,0,6,1,5]
// [0,1,2,3,5,4]

func FindMissingInt(nums []int) int {
	slices.Sort(nums)

	var before int
	for i, num := range nums {
		if i == 0 {
			before = num
			continue
		}
		if num != before+1 {
			return before + 1
		}

		before = num
	}

	return -1
}
