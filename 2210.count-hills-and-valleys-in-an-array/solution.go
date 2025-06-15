func countHillValley(nums []int) int {
    asc := false
    dsc := false
    res := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i - 1] {
            asc = true
            if dsc {
                res++
                dsc = false
            }
        } else if nums[i] < nums[i - 1] {
            dsc = true
            if asc {
                res++
                asc = false
            }
        }
    }
    return res
}