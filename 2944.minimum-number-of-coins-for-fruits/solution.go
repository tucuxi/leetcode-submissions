func minimumCoins(prices []int) int {
    memo := make([]int, len(prices) + 1)
    
    var dp func(int) int
    dp = func(i int) int {
        if i >= len(memo) {
            return 0
        }
        if memo[i] > 0 {
            return memo[i]
        }
        min := math.MaxInt
        for j := i + 1; j <= 2 * i + 1; j++ {
            if p := dp(j); p < min {
                min = p
            } 
        }
        memo[i] = prices[i - 1] + min
        return memo[i]
    }
    
    return dp(1)
}