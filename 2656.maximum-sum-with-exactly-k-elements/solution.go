func maximizeSum(nums []int, k int) int {
    m := 0
    for _, n := range nums {
        if n > m {
            m = n
        }
    }
    return k * m + k * (k - 1) / 2
}