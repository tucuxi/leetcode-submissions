func predictTheWinner(nums []int) bool {
    n := len(nums)
    memo := make([][]int, n)

    for i := range memo {
        memo[i] = make([]int, n)
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
        memo[i][j] = max(nums[i] - dp(i + 1, j), nums[j] - dp(i, j-1))
        return memo[i][j]
    }
    
    return dp(0, n-1) >= 0
}