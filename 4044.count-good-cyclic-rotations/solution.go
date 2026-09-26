func countGoodRotations(nums []int) int {
    n := len(nums)
    a, b := 0, 0
    for i := range n/2 {
        a += nums[i]
        b += nums[n-1-i]
    }
    res := 0
    for i := range n {
        if a > b {
            res++
        }
        x := nums[(i + n/2) % n] - nums[i]
        a += x
        b -= x
    }
    return res
}