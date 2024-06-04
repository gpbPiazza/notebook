package sorting

func quickSort(nums []int, left, right int) {
	if right-left <= 0 {
		return
	}

	pivot := partition(nums, left, right)

	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}
