package sorting

func SelectionSort(nums []int) []int {

	lastIndex := len(nums) - 1
	for i := 0; i < lastIndex; i++ {
		lowestValIndex := i

		for j := i + 1; j < len(nums); j++ {
			jVal := nums[j]
			if jVal <= nums[lowestValIndex] {
				lowestValIndex = j
			}
		}
		if lowestValIndex != i {
			nums[lowestValIndex], nums[i] = nums[i], nums[lowestValIndex]
		}
	}

	return nums
}
