package main

func MinDiff(n int, earnings []int, k int) int {
	var alteredEarnings []int
	for i := range earnings {
		earn := earnings[i]

		earnToAlter := earn
		if earn >= k {
			earnToAlter = earnToAlter - k
		}

		if earn < k {
			earnToAlter = earnToAlter + k
		}
		alteredEarnings = append(alteredEarnings, earnToAlter)
	}

	maxEarn, minEarn := findMaxAndMinEarns(alteredEarnings)
	return maxEarn - minEarn
}

func findMaxAndMinEarns(earnings []int) (int, int) {
	minEarn := earnings[0]
	maxEarn := earnings[0]
	for i := range earnings {
		earn := earnings[i]

		if earn > maxEarn {
			maxEarn = earn
		}
		if earn < minEarn {
			minEarn = earn
		}
	}
	return maxEarn, minEarn
}
