func minCost(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]bool, n)
	newDP := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, 1024)
		newDP[i] = make([]bool, 1024)
	}

	dp[0][grid[0][0]] = true

	for j := 1; j < n; j++ {
		for x := range 1024 {
			if dp[j-1][x] {
				dp[j][x^grid[0][j]] = true
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := range n {
			for x := range 1024 {
				newDP[j][x] = false
			}
		}
		for x := range 1024 {
			if dp[0][x] {
				newDP[0][x^grid[i][0]] = true
			}
		}
		for j := 1; j < n; j++ {
			for x := range 1024 {
				if dp[j][x] || newDP[j-1][x] {
					newDP[j][x^grid[i][j]] = true
				}
			}
		}
		dp, newDP = newDP, dp
	}

	for x := range 1024 {
		if dp[n-1][x] {
			return x
		}
	}

	return 0
}