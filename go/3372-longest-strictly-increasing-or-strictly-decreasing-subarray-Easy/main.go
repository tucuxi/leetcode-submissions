func longestMonotonicSubarray(nums []int) int {
    res, incr, decr := 1, 1, 1
    for i := 1; i < len(nums); i++ {
        switch sgn(nums[i] - nums[i - 1]) {
            case -1: incr =  1; decr++
            case 0: incr = 1; decr = 1
            case 1: incr++; decr = 1
        }
        res = max(res, incr, decr)
    }
    return res
}

func sgn(a int) int {
    if a < 0 {
        return -1
    }
    if a > 0 {
        return 1
    }
    return 0
}