package array

func Intersection(slice1, slice2 []int) []int {
	sliceMap1 := make(map[int]bool)
	var result []int

	for _, item := range slice1 {
		sliceMap1[item] = true
	}

	for _, item := range slice2 {
		if sliceMap1[item] {
			result = append(result, item)
		}
	}

	return result
}
