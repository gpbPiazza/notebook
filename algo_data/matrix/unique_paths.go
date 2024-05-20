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
