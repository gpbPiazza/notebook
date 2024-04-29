package array

func FindFirstDuplicated(slice []string) string {
	alreadyFoundMap := make(map[string]bool)

	for _, item := range slice {
		if alreadyFoundMap[item] {
			return item
		}
		alreadyFoundMap[item] = true
	}

	return ""
}
