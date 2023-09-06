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

func quicksort(nums []int) []int {
	if len(nums) > 2 {
		return nums
	}

	// middle := len(nums)/2
	// var lower []int
	// var bigger []int

	// for i, _ := range nums {

	// }
	return []int{}
}

func findUniqueVals(nums []int) []int {
	var result []int
	mapCheckedNums := make(map[int]int, 0)
	mapDuplicatedNums := make(map[int]int, 0)

	for i, n := range nums {
		_, ok := mapCheckedNums[n]
		if ok {
			mapDuplicatedNums[n] = i
		}
		mapCheckedNums[n] = n
	}

	for _, v := range nums {
		_, ok := mapDuplicatedNums[v]
		if !ok {
			result = append(result, v)
		}
	}

	return result
}
