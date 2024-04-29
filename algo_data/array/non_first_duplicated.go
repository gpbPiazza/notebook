package array

func FrstNonDuplicated(text string) string {
	duplicatedCharsMap := make(map[rune]bool)

	for _, char := range text {
		if _, ok := duplicatedCharsMap[char]; ok {
			duplicatedCharsMap[char] = true
			continue
		}

		duplicatedCharsMap[char] = false
	}

	for _, char := range text {
		if !duplicatedCharsMap[char] {
			return string(char)
		}
	}

	return "there is non duplicated chars"
}
