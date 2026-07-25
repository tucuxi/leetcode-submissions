func maxValidPairSum(nums []int, k int) int {
    n := len(nums)
    suffix := make([]int, n-k)
    m := 0
    
    for i := n-k-1; i >= 0; i-- {
        m = max(nums[i+k], m)
        suffix[i] = m
    }

    res := 0

    for i := range n-k {
        sum := nums[i] + suffix[i]
        res = max(sum, res)
    }

    return res
}