package matrix

// I always will start from the position matrix[0,0] to the matrix[3,7]
// 3 Unique paths
// Cases:
// rows = 3
// columns = 8
// [S],[],[],[],[],[],[],[]
// [],[],[],[],[],[],[],[]
// [],[],[],[],[],[],[],[f]

// rows = 1
// columns = 7
// [S],[],[],[],[],[],[F]

// rows = 2
// columns = 4
// [S],[],[],[]
// [],[],[],[]

func uniquePaths(rows, columns int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	return uniquePaths(rows-1, columns) + uniquePaths(rows, columns-1)
}

func uniquePaths_with_memo(rows, columns int, memo map[[2]int]int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	_, ok := memo[[2]int{rows, columns}]
	if !ok {
		memo[[2]int{rows, columns}] = uniquePaths_with_memo(rows-1, columns, memo) + uniquePaths_with_memo(rows, columns-1, memo)
	}

	return memo[[2]int{rows, columns}]
}
