func zeroFilledSubarray(nums []int) int64 {
    var res, run int64
    for _, n := range nums {
        if n == 0 {
            run++
            res += run
        } else {
            run = 0
        }
    }
    return res
}