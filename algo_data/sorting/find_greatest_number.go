package sorting

import "sort"

func FindGreatest_O_of_n(numbers []int) int {
	var greatest int
	for i := range numbers {
		if greatest < numbers[i] {
			greatest = numbers[i]
		}
	}
	return greatest
}

func FindGreatest_O_of_n_2(numbers []int) int {
	var greatest int

	for i := range numbers {
		for j := range numbers {
			if greatest < numbers[j] {
				greatest = numbers[i]
			}
		}

	}
	return greatest
}

func FindGreatest_O_of_n_log_n(numbers []int) int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[j] > numbers[i]
	})

	return numbers[len(numbers)-1]
}
