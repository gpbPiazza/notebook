package array

func OnlyEven(numbers []int) []int {
	if len(numbers) == 0 {
		return nil
	}

	firstElement := numbers[0]
	if firstElement%2 == 0 {
		return append([]int{firstElement}, OnlyEven(numbers[1:])...)
	} else {
		return OnlyEven(numbers[1:])
	}
}
