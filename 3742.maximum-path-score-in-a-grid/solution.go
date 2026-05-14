func maxPathScore(grid [][]int, k int) int {

    costScore := func(i, j int) (int, int) {
        g := grid[i][j]
        return min(1, g), g
    }

	m, n := len(grid), len(grid[0])

	const INF = math.MinInt32

	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for c := range dp[i][j] {
				dp[i][j][c] = INF
			}
		}
	}

	dp[0][0][0] = 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for c := 0; c <= k; c++ {
				if dp[i][j][c] == INF {
					continue
				}

				if i+1 < m {
					cost, score := costScore(i+1, j)
                    newCost, newScore := c+cost, dp[i][j][c]+score
					if newCost <= k {
						if dp[i+1][j][newCost] < newScore {
							dp[i+1][j][newCost] = newScore
						}
					}
				}

				if j+1 < n {
   					cost, score := costScore(i, j+1)
                    newCost, newScore := c+cost, dp[i][j][c]+score
					if newCost <= k {
						if dp[i][j+1][newCost] < newScore {
							dp[i][j+1][newCost] = newScore
						}
					}
				}
			}
		}
	}

	ans := INF
	for c := 0; c <= k; c++ {
		if dp[m-1][n-1][c] > ans {
			ans = dp[m-1][n-1][c]
		}
	}

	if ans < 0 {
		return -1
	}
	return ans
}