package sequences

func golomb(nth int) int {
	if nth == 1 {
		return 1
	}
	return 1 + golomb(nth-golomb(golomb(nth-1)))
}

func golomb_with_memo(nth int, memo map[int]int) int {
	if nth == 1 {
		return 1
	}

	_, ok := memo[nth]
	if !ok {
		memo[nth] = 1 + golomb_with_memo(nth-golomb_with_memo(golomb_with_memo(nth-1, memo), memo), memo)
	}

	return memo[nth]
}
