func minOperations(nums []int, x int) int {
    target := -x
    for _, num := range nums {
        target += num
    }
    maxLen := -1 
    curSum := 0
    p := 0
    for q := range nums {
        curSum += nums[q]
        for ; curSum > target && p <= q; p++ {
            curSum -= nums[p]
        }
        if curSum == target {
            maxLen = max(maxLen, q-p+1)
        }
    }
    if maxLen == -1 {
        return -1
    }
    return len(nums) - maxLen
}
