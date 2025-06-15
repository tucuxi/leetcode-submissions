func longestSquareStreak(nums []int) int {
    m := make(map[int]bool, len(nums))
    for _, n := range nums {
        m[n] = true
    }
    res := 0
    for _, n := range nums {
        l := 0
        for k := n; m[k]; k *= k {
            l++
        }
        if l > res {
            res = l
        }
    }
    if res < 2 {
        return -1
    }
    return res
}