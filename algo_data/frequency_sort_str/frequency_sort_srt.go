package frequencysortstr

import "sort"

func frequencySort(s string) string {
	frequencyChars := make(map[rune]int)
	var sliceChars []rune

	for _, char := range s {
		frequency, exists := frequencyChars[char]
		sliceChars = append(sliceChars, char)
		if exists {
			frequency++
			frequencyChars[char] = frequency
			continue
		}
		frequencyChars[char] = 1
	}

	sort.Slice(sliceChars, func(i int, j int) bool {
		return sliceChars[i] < sliceChars[j]
	})

	sort.Slice(sliceChars, func(i int, j int) bool {
		return frequencyChars[sliceChars[i]] > frequencyChars[sliceChars[j]]
	})

	return string(sliceChars)
}
