package sequences

func GetTriangularNumber(nthNumber int) int {
	if nthNumber == 1 || nthNumber <= 0 {
		return 1
	}

	return nthNumber + GetTriangularNumber(nthNumber-1)
}
