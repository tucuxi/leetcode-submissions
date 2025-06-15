func longestArithSeqLength(nums []int) int {
    dp := make([]map[int]int, len(nums))
    maxLength := 0
    for r := 1; r < len(nums); r++ {
        dp[r] = make(map[int]int)
        for l := 0; l < r; l++ {
            d := nums[r] - nums[l]
            dp[r][d] = dp[l][d] + 1
            if dp[r][d] > maxLength {
                maxLength = dp[r][d]
            }
        }
    }
    return maxLength + 1
}