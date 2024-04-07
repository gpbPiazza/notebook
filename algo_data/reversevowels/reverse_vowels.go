package reversevowels

func reverseVowels(s string) string {
	left, right := 0, len(s)-1

	result := []rune(s)
	for left < right {
		if !isVowel(rune(s[right])) {
			right--
			continue
		}

		if !isVowel(rune(s[left])) {
			left++
			continue
		}

		// if left and rifght are vowels
		result[left] = rune(s[right])
		result[right] = rune(s[left])
		right--
		left++
	}

	return string(result)
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
func reverseString(s string) string {
	left, right := 0, len(s)-1

	result := []byte(s)
	for left < right {
		result[right] = s[left]
		result[left] = s[right]
		left++
		right--
	}

	return string(result)
}
