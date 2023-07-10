package main

import "fmt"

func main() {
	zap := []int{1, 2, 3, 7, 0, 5, 1, 4, 66, 33}
	countElements(zap)
	fmt.Println(sum(zap))
	fmt.Println(findBiggest(zap))
}

func countElements(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	printCount := 1 + countElements(nums[1:])
	fmt.Println(printCount)
	return printCount
}

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sum(nums[1:])
}

func findBiggest(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	biggest := nums[0]
	newBiggest := findBiggest(nums[1:])

	if newBiggest > biggest {
		biggest = newBiggest
	}

	return biggest
}

func binarySearch(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return nums[0]
		}
		return -1
	}

	if len(nums) == 0 {
		return -1
	}

	quantityOfElements := len(nums)
	middle := quantityOfElements / 2
	middleElement := nums[middle]

	if middleElement == target {
		return middleElement
	}

	if middleElement > target {
		return binarySearch(nums[:middle], target)
	}
	if middleElement < target {
		return binarySearch(nums[middle:], target)
	}
	return -1
}
