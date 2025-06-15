func sumCounts(nums []int) int {
    res := 0
    for i := range nums {
        var s [101]bool
        k := 0
        for _, n := range nums[i:] {
            if !s[n] {
                s[n] = true
                k++
            }
            res += k * k
        }
    }
    return res
}