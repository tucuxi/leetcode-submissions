func getAverages(nums []int, k int) []int {
    res := make([]int, len(nums))
    for i := range res {
        res[i] = -1
    }
    if len(nums) <= 2 * k {
        return res
    }
    sum := nums[k]
    for i := 1; i <= k; i++ {
        sum += nums[k - i] + nums[k + i]
    }
    i := k
    for {
        res[i] = sum / (k + k + 1)
        if i + 1 == len(nums) - k {
            return res
        } 
        sum += -nums[i - k] + nums[i + 1 + k]
        i++
    }
}