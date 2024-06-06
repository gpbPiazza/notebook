package sorting

import "sort"

func GreatestProductTree(numbers []uint) uint {
	if len(numbers) < 3 {
		return 0
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[j] > numbers[i]
	})

	last := len(numbers) - 1

	return numbers[last] * numbers[last-1] * numbers[last-2]
}
