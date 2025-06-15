func maxSumAfterPartitioning(arr []int, k int) int {
    dp := make([]int, len(arr) + 1)
    for i := len(arr) - 1; i >= 0; i-- {
        a := 0
        for j := 0; j < min(k, len(arr) - i); j++ {
            p := i + j
            if arr[p] > a {
                a = arr[p]
            }
            if s := (j + 1) * a + dp[p + 1]; s > dp[i] {
                dp[i] = s
            }
        }
    }
    return dp[0]
}