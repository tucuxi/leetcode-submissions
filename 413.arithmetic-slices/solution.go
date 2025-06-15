func numberOfArithmeticSlices(nums []int) int {
    res := 0
    run := 0
    for i := 2; i < len(nums); i++ {
        if nums[i - 2] - nums[i - 1] == nums[i - 1] - nums[i] {
            run++
            res += run
        } else {
            run = 0
        }
    }
    return res
}
