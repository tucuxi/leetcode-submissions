func dominantIndices(nums []int) int {
    res := 0
    sum := 0

    for i := len(nums) - 1; i > 0; i-- {
        sum += nums[i]
        if nums[i - 1] * (len(nums) - i) > sum {
            res++
        }
    }

    return res
}