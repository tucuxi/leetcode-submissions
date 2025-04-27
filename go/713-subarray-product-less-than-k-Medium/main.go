func numSubarrayProductLessThanK(nums []int, k int) int {
    res := 0
    for i := range nums {
        p := 1
        for j := i; j < len(nums); j++ {
            p *= nums[j]
            if p >= k {
                break
            }
            res++
        }
    }
    return res
}