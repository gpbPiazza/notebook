package bubblesort

func BubbleSort(nums []int) []int {
	lastIndex := len(nums) - 1
	isSorted := false

	for !isSorted {
		isSorted = true
		for i := 0; i < lastIndex; i++ {
			nextIndex := i + 1

			if nums[i] > nums[nextIndex] {
				nums[nextIndex], nums[i] = nums[i], nums[nextIndex]
				isSorted = false
			}
		}
		lastIndex--
	}

	return nums
}
