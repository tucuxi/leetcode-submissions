func check(nums []int) bool {
    found := false
    p := nums[len(nums) - 1]
    for _, n := range nums {
        if p > n {
            if found {
                return false
            }
            found = true
        }
        p = n
    }
    return true
}