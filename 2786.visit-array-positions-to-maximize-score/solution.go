func maxScore(nums []int, x int) int64 {
    dp := make([]int64, len(nums))
    dp[0] = int64(nums[0])
    for i := 1; i < len(nums); i++ {
        dp[i] = math.MinInt
    }
    var max int64 = dp[0]
    for i := range nums {
        j := i + 1
        for j < len(nums) && nums[i] & 1 == nums[j] & 1 {
            j++
        }
        if j < len(nums) {
            d := dp[i] + int64(nums[j]) - int64(x)
            if d > dp[j] {
                dp[j] = d
            }
            if d > max {
                max = d
            }
        }
        k := i + 1
        for k < len(nums) && nums[i] & 1 != nums[k] & 1 {
            k++
        }
        if k < len(nums) {
            d := dp[i] + int64(nums[k])
            if d > dp[k] {
                dp[k] = d
            }
            if d > max {
                max = d
            }
        }
    }
    return max   
}