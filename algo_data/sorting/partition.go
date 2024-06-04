package sorting

func partition(nums []int, left, right int) int {
	pivot := right
	right--

	for {
		for !(nums[left] >= nums[pivot]) {
			left++
		}

		for !(nums[right] <= nums[pivot]) {
			right--
		}

		if left >= right {
			break
		} else {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
	}

	nums[pivot], nums[left] = nums[left], nums[pivot]

	return left
}
