class Solution {

    fun uniquePathsWithObstacles(obstacleGrid: Array<IntArray>): Int {
  		val M = obstacleGrid.size
		val N = obstacleGrid[0].size

		val dp = Array(M + 1){IntArray(N + 1)}

		for (i in 1..M) {
			for (j in 1..N) {
				if (obstacleGrid[i - 1][j - 1] == 1) {
					continue
				}
				dp[i][j] = if (i == 1 && j == 1) {
					1
				} else {
					dp[i - 1][j] + dp[i][j - 1]
				}            
			}
		}
		return dp[M][N]
    }
}