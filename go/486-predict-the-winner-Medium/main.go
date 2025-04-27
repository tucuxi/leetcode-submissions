func PredictTheWinner(nums []int) bool {
    memo := make([][]int, len(nums))
    for i := range memo {
        memo[i] = make([]int, len(nums))
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    
    var dp func(int, int) int
    dp = func(i, j int) int {
        if i == j {
            return nums[i]
        }
        if memo[i][j] != -1 {
            return memo[i][j]
        }
        memo[i][j] = max(nums[i] - dp(i + 1, j), nums[j] - dp(i, j - 1))
        return memo[i][j]
    }
    
    return dp(0, len(nums) - 1) >= 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}