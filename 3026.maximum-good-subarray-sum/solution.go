func maximumSubarraySum(nums []int, k int) int64 {
    var sum, maxSum int64 = 0, math.MinInt64
    minPrefix := make(map[int]int64)
    
    for _, n := range nums {
        if s, exists := minPrefix[n]; !exists || s > sum {
            minPrefix[n] = sum
        }
        sum += int64(n)
        if s, exists := minPrefix[n-k]; exists && sum - s > maxSum {
            maxSum = sum - s
        }
        if s, exists := minPrefix[n+k]; exists && sum - s > maxSum {
            maxSum = sum - s
        }
    }
    
    if maxSum == math.MinInt64 {
        return 0
    }
    return maxSum 
}