func rob(nums []int) int {
    res := nums[0]
    if a := rob1(nums[0:len(nums) - 1]); a > res {
        res = a
    }
    if b := rob1(nums[1:]); b > res {
        res = b
    }
    return res
}

func rob1(nums []int) int {
    p, q := 0, 0
    for _, n := range nums {
        if p + n > q {
            p, q = q, p + n
        } else {
            p = q
        }
    }
    return q
}