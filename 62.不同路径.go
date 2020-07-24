package solution

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	result := make([][]int, n)
	for index := range result {

		result[index] = make([]int, m)
		if index == 0 {
			for i := range result[index] {
				result[index][i] = 1
			}
		}
		result[index][0] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}
	return result[n-1][m-1]

}
