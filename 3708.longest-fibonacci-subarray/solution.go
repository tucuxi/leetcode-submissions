func longestSubarray(nums []int) int {
    res := 0
    run := 2

    for i := 2; i < len(nums); i++ {
        if nums[i - 2] + nums[i - 1] == nums[i] {
            run++
        } else {
            run = 2
        }
        if run > res {
            res = run
        }
    }

    return res
}