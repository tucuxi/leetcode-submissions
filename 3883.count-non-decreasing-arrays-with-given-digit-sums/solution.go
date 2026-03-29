const MOD = 1_000_000_007
const MAX_VALUE = 5000

func countArrays(digitSum []int) int {
	if len(digitSum) == 0 {
		return 0
	}

	var dsum [MAX_VALUE + 1]int
	for i := 1; i <= MAX_VALUE; i++ {
		dsum[i] = dsum[i/10] + i%10
	}

	var dp [MAX_VALUE + 1]int
	for v := 0; v <= MAX_VALUE; v++ {
		if dsum[v] == digitSum[0] {
			dp[v] = 1
		}
	}

	for i := 1; i < len(digitSum); i++ {
		target := digitSum[i]
		var newDp [MAX_VALUE + 1]int
		runningSum := 0
		for v := 0; v <= MAX_VALUE; v++ {
			runningSum = (runningSum + dp[v]) % MOD
			if dsum[v] == target {
				newDp[v] = runningSum
			}
		}
		dp = newDp
	}

	ans := 0
	for v := 0; v <= MAX_VALUE; v++ {
		ans = (ans + dp[v]) % MOD
	}
	return ans
}