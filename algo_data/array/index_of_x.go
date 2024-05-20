package array

func indexOfX(text string, index int) int {
	currentChar := rune(text[index])
	if currentChar == 'x' {
		return index
	}

	return indexOfX(text, index+1)
}

func indexOfX_SecondImplementation(text string) int {
	currentChar := rune(text[0])
	if currentChar == 'x' {
		return 0
	}
	return indexOfX_SecondImplementation(text[1:]) + 1
}
