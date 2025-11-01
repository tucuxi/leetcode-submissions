func magicalSum(m int, k int, nums []int) int {
	MOD := 1000000007
	pw := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		pw[i] = make([]int, m+1)
		pw[i][0] = 1
		for j := 1; j <= m; j++ {
			pw[i][j] = pw[i][j-1] * nums[i]
			pw[i][j] %= MOD
		}
	}
	comb := make([][]int, m+1)
	for i := range comb {
		comb[i] = make([]int, m+1)
		comb[i][i], comb[i][0] = 1, 1
		for j := 1; j < i; j++ {
			comb[i][j] = (comb[i-1][j] + comb[i-1][j-1]) % MOD
		}
	}
	dp := make([][][][]int, len(nums))
	for i := range dp {
		dp[i] = make([][][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([][]int, m+1)
			for l := range dp[i][j] {
				dp[i][j][l] = make([]int, k+1)
				for c := range dp[i][j][l] {
					dp[i][j][l][c] = -1
				}
			}
		}
	}
	return dfs(dp, nums, 0, m, k, 0, comb, pw)
}

func dfs(dp [][][][]int, nums []int, index int, m int, k int, carry int, comb [][]int, pw [][]int) int {
	if m == 0 && bits.OnesCount(uint(carry)) == k {
			return 1
	}
    if index == len(nums) {
		return 0
	}
	if dp[index][m][carry][k] != -1 {
		return dp[index][m][carry][k]
	}
	dp[index][m][carry][k] = 0
	for c := 0; c <= m; c++ {
		total := carry + c
		bit := total & 1
		carryOut := total >> 1
		m2 := m - c
		k2 := k - bit
		if k2 < 0 {
			continue
		}
		ways := comb[m][c] * pw[index][c] % 1000000007
		dp[index][m][carry][k] = (dp[index][m][carry][k] + dfs(dp, nums, index+1, m2, k2, carryOut, comb, pw)*ways) % 1000000007
	}
	return dp[index][m][carry][k]
}