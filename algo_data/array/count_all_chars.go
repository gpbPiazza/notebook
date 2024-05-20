package array

func CountAllChars(strings []string) int {
	if len(strings) == 0 {
		return 0
	}

	firstElement := strings[0]

	return len(firstElement) + CountAllChars(strings[1:])
}
