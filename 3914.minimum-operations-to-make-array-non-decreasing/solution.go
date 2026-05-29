func minOperations(nums []int) int64 {
    var res int64

    for i := 1; i < len(nums); i++ {
        if d := nums[i-1] - nums[i]; d > 0 {
            res += int64(d)
        }
    }
    return res
}