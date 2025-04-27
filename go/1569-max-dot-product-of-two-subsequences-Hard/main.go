func maxDotProduct(nums1 []int, nums2 []int) int {
    firstMax := math.MinInt
    secondMax := math.MinInt
    firstMin := math.MaxInt
    secondMin := math.MaxInt
        
    for _, num := range nums1 {
        firstMax = max(firstMax, num)
        firstMin = min(firstMin, num)
    }
        
    for _, num := range nums2 {
        secondMax = max(secondMax, num)
        secondMin = min(secondMin, num)
    }
        
    if firstMax < 0 && secondMin > 0 {
        return firstMax * secondMin
    }
        
    if firstMin > 0 && secondMax < 0 {
        return firstMin * secondMax
    }
        
    dp := make([][]int, len(nums1) + 1)
    for i := range dp {
        dp[i] = make([]int, len(nums2) + 1)
    }
    
    for i := len(nums1) - 1; i >= 0; i-- {
        for j := len(nums2) - 1; j >= 0; j-- {
            use := nums1[i] * nums2[j] + dp[i + 1][j + 1]
            dp[i][j] = max(use, max(dp[i + 1][j], dp[i][j + 1]))
        }
    }
        
    return dp[0][0];
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}