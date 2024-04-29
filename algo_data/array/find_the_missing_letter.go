package array

import (
	"strings"
)

var alphabetLetters = map[rune]bool{
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'g': true,
	'h': true,
	'i': true,
	'j': true,
	'k': true,
	'l': true,
	'm': true,
	'n': true,
	'o': true,
	'p': true,
	'q': true,
	'r': true,
	's': true,
	't': true,
	'u': true,
	'v': true,
	'w': true,
	'x': true,
	'y': true,
	'z': true,
}

// this function assume thata always will have a string with all letters of
// the alphabet and will miss only one letter.
func FindMissingAlphabetLetter(text string) string {
	withoutSpaces := strings.Replace(text, " ", "", -1)
	allAlphabetLetter := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	existingLetterInTheTextMap := make(map[rune]bool)
	for _, char := range withoutSpaces {
		existingLetterInTheTextMap[char] = true
	}

	for _, char := range allAlphabetLetter {
		if !existingLetterInTheTextMap[char] {
			return string(char)
		}
	}

	return "input text is not in the expected format"
}
