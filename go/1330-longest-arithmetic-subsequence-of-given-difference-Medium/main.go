func longestSubsequence(arr []int, difference int) int {
    max := 1
    dp := make(map[int]int)
    for _, a := range arr {
        p := dp[a - difference]
        dp[a] = p + 1
        if dp[a] > max {
            max = dp[a]
        }
    }
    return max
}