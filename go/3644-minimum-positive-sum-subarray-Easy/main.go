func minimumSumSubarray(nums []int, l int, r int) int {
    minSum := math.MaxInt
    
    pre := make([]int, len(nums) + 1)
    for i := range nums {
        pre[i + 1] = pre[i] + nums[i]
    }

    for size := l; size <= r; size++ {
        for i := 0; i < len(pre) - size; i++ {
            if sum := pre[i + size] - pre[i]; sum > 0 {
                minSum = min(minSum, sum)
            }
        }
    }

    if minSum == math.MaxInt {
        return -1
    }
    return minSum
}