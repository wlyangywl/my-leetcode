package unique_path_62

func uniquePaths(m int, n int) int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		cur := make([]int, n)
		arr[i] = cur
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				arr[i][j] = 1
			} else if i == 0 {
				arr[i][j] = 1
			} else if j == 0 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j] + arr[i][j-1]
			}
		}
	}
	return arr[m-1][n-1]
}
