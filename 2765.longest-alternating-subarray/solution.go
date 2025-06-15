func alternatingSubarray(nums []int) int {
    max := -1
    run := 1
    exp := 1
    for i := 1; i < len(nums); i++ {
        if d := nums[i] - nums[i - 1]; d == exp {
            run++
            if run > max {
                max = run
            }
            exp = -exp
        } else if d == 1 {
            run = 2
            exp = -1
        } else {
            run = 1
            exp = 1
        }
    }
    return max
}