func longestSubarray(nums []int) int {
    res := 0
    s1, s2 := -1, 0
    for i := range nums {
        if nums[i] == 0 {
            if s1 >= 0 && s1 + s2 > res {
                res = s1 + s2
            }
            s1 = s2
            s2 = 0
        } else {
            s2++
        }
    }
    if s1 < 0 {
        return s2 - 1
    }
    if s1 + s2 > res {
        res = s1 + s2
    }
    return res
}